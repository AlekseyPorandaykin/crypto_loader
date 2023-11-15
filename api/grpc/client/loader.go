package client

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/api/grpc/client/specification"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
)

type Loader struct {
	batchCh chan Batch
	address string
}

func DefaultLoader() *Loader {
	return NewLoader(":50052")
}

func NewLoader(address string) *Loader {
	return &Loader{
		batchCh: make(chan Batch, 1),
		address: address,
	}
}

func (l *Loader) Start(ctx context.Context, durationSec int64) error {
	conn, err := grpc.Dial(l.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()
	c := specification.NewEventServiceClient(conn)
	r, err := c.TickerPrices(
		ctx,
		&specification.DurationSeconds{Second: durationSec},
		grpc.EmptyCallOption{},
		grpc.MaxRecvMsgSizeCallOption{1024 * 1024 * 10},
	)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			symbols, err := r.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			batch := NewBatch(len(symbols.Prices))

			zap.L().Debug("load prices from crypto_loader", zap.Int("count", len(symbols.Prices)))
			for _, item := range symbols.GetPrices() {
				batch.Append(&SymbolPrice{
					Exchange: item.GetExchange(),
					Symbol:   item.GetSymbol(),
					Price:    item.GetPrice(),
					Date:     item.GetDate().AsTime(),
				})
			}
			l.batchCh <- batch
		}
	}
}

func (l *Loader) Batch() <-chan Batch {
	return l.batchCh
}

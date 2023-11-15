package grpc

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/server/grpc/specification"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

type handler struct {
	priceStorage domain.PriceStorage
	specification.UnimplementedEventServiceServer
}

func NewHandler(priceStorage domain.PriceStorage) specification.EventServiceServer {
	return &handler{
		priceStorage: priceStorage,
	}
}

func (h *handler) Prices(ctx context.Context, req *specification.EmptyRequest) (*specification.SymbolPrices, error) {
	sp, err := h.priceStorage.LastPrices(ctx)
	if err != nil {
		return nil, err
	}
	prices := make([]*specification.SymbolPrice, 0, len(sp))
	for _, item := range sp {
		price, err := strconv.ParseFloat(item.Price, 32)
		if err != nil {
			zap.L().Error(
				"parse price",
				zap.Error(err),
				zap.String("price", item.Price),
				zap.String("symbol", item.Symbol),
				zap.String("exchange", item.Exchange),
			)
			continue
		}
		prices = append(prices, &specification.SymbolPrice{
			Exchange: item.Exchange,
			Symbol:   item.Symbol,
			Price:    float32(price),
			Date:     timestamppb.New(item.Date),
		})
	}
	return &specification.SymbolPrices{Prices: prices}, nil
}

func (h *handler) TickerPrices(req *specification.DurationSeconds, stream specification.EventService_TickerPricesServer) error {
	ticker := time.NewTicker(time.Duration(req.GetSecond()) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case <-ticker.C:
			sp, err := h.priceStorage.LastPrices(stream.Context())
			if err != nil {
				return err
			}
			prices := make([]*specification.SymbolPrice, 0, len(sp))
			for _, item := range sp {
				price, err := strconv.ParseFloat(item.Price, 32)
				if err != nil {
					zap.L().Error("parse price",
						zap.Error(err),
						zap.String("price", item.Price),
						zap.String("symbol", item.Symbol),
						zap.String("exchange", item.Exchange))
					continue
				}
				prices = append(prices, &specification.SymbolPrice{
					Exchange: item.Exchange,
					Symbol:   item.Symbol,
					Price:    float32(price),
					Date:     timestamppb.New(item.Date),
				})
			}
			if err := stream.Send(&specification.SymbolPrices{Prices: prices}); err != nil {
				return err
			}
		}
	}
}

func (h *handler) SymbolPrice(ctx context.Context, req *specification.SymbolPriceRequest) (*specification.SymbolPrices, error) {
	sp, err := h.priceStorage.SymbolPrice(ctx, req.GetSymbol())
	if err != nil {
		return nil, err
	}
	prices := make([]*specification.SymbolPrice, 0, len(sp))
	for _, item := range sp {
		price, err := strconv.ParseFloat(item.Price, 32)
		if err != nil {
			zap.L().Error("parse price",
				zap.Error(err),
				zap.String("price", item.Price),
				zap.String("symbol", item.Symbol),
				zap.String("exchange", item.Exchange))
			continue
		}
		prices = append(prices, &specification.SymbolPrice{
			Exchange: item.Exchange,
			Symbol:   item.Symbol,
			Price:    float32(price),
			Date:     timestamppb.New(item.Date),
		})
	}
	return &specification.SymbolPrices{Prices: prices}, nil
}

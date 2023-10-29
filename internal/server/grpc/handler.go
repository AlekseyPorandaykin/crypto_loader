package grpc

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/server/grpc/specification"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
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
			zap.L().Error("parse parse", zap.Error(err))
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

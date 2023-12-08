package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance/requests"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance/sender"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type ResponseHandlerFunc func(r io.Reader) error

type Manager struct {
	personalSenders map[string]*sender.Personal
	spotRequest     *requests.SpotRequest
	futureRequest   *requests.FutureRequest
	sender          sender.Sender
}

func NewManager(spotHost string, futureHost string) (*Manager, error) {
	sr, err := requests.NewSpotRequest(spotHost)
	if err != nil {
		return nil, err
	}
	fr, err := requests.NewFutureRequest(futureHost)
	if err != nil {
		return nil, err
	}
	return &Manager{
		sender:          sender.NewBasic(),
		personalSenders: make(map[string]*sender.Personal),
		spotRequest:     sr,
		futureRequest:   fr,
	}, nil
}

func (m *Manager) WithSender(sender sender.Sender) {
	m.sender = sender
}

func (m *Manager) Close() {
	m.sender.Close()
}

func (m *Manager) GetPrice(ctx context.Context) ([]domain.PriceSymbolDTO, error) {
	body, err := m.executeRequest(m.spotRequest.SymbolPriceTicker(ctx))
	if err != nil {
		return nil, err
	}
	prices := make([]domain.PriceSymbolDTO, 0, 3500)
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, err
	}

	return prices, nil
}

func (m *Manager) PriceChangeStatistics(
	ctx context.Context, symbols []string, widowSize string,
) ([]domain.PriceChangeStatisticDTO, error) {
	body, err := m.executeRequest(m.spotRequest.RollingWindowPriceChangeStatistics(ctx, symbols, widowSize))
	if err != nil {
		return nil, err
	}
	dto := make([]domain.PriceChangeStatisticDTO, 0, len(symbols))
	if err := json.Unmarshal(body, &dto); err != nil {
		return nil, err
	}
	return dto, nil
}

func (m *Manager) PriceChangeStatisticsLastHour(ctx context.Context, symbols []string) ([]domain.PriceChangeStatisticDTO, error) {
	return m.PriceChangeStatistics(ctx, symbols, "1h")
}

func (m *Manager) FuturesSymbolPriceTicker(ctx context.Context) error {
	_, err := m.executeRequest(m.futureRequest.SymbolPriceTicker(ctx))
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) FuturesNewOrder(cred domain.CredentialDTO, order domain.FutureOrder) (domain.ResponseOrderDTO, error) {
	var orderRep domain.ResponseOrderDTO
	req, weight, err := m.futureRequest.NewOrder(cred.APIKey, cred.ApiSecret, order)
	if err != nil {
		return domain.ResponseOrderDTO{}, err
	}
	data, err := m.sendPersonalRequest(cred, req, weight)
	if err != nil {
		return domain.ResponseOrderDTO{}, err
	}
	if err := json.Unmarshal(data, &orderRep); err != nil {
		return domain.ResponseOrderDTO{}, err
	}

	return orderRep, nil
}

func (m *Manager) FuturesExchangeInformation(ctx context.Context) ([]byte, error) {
	req, weight, err := m.futureRequest.ExchangeInformation(ctx)
	if err != nil {
		return nil, err
	}

	return m.sendRequest(req, weight)
}

func (m *Manager) FuturesQueryIndexPriceConstituents(ctx context.Context, symbol string) ([]byte, error) {
	req, weight, err := m.futureRequest.QueryIndexPriceConstituents(ctx, symbol)
	if err != nil {
		return nil, err
	}

	return m.sendRequest(req, weight)
}

func (m *Manager) FuturesPlaceMultipleOrders(
	cred domain.CredentialDTO, orders []domain.FutureOrder,
) ([]domain.ResponseOrderDTO, error) {
	req, weight, err := m.futureRequest.PlaceMultipleOrders(cred.APIKey, cred.ApiSecret, orders)
	if err != nil {
		return nil, err
	}
	data, err := m.sendPersonalRequest(cred, req, weight)
	if err != nil {
		return nil, err
	}
	ordersRep := make([]domain.ResponseOrderDTO, 0, len(orders))
	if err := json.Unmarshal(data, &ordersRep); err != nil {
		return nil, err
	}

	return ordersRep, nil
}

func (m *Manager) FutureChangeInitialLeverage(
	cred domain.CredentialDTO, symbol string, leverage int,
) (domain.LeverageDTO, error) {
	req, weight, err := m.futureRequest.ChangeInitialLeverage(cred.APIKey, cred.ApiSecret, symbol, leverage)
	if err != nil {
		return domain.LeverageDTO{}, err
	}
	data, err := m.sendPersonalRequest(cred, req, weight)
	if err != nil {
		return domain.LeverageDTO{}, err
	}
	var resp domain.LeverageDTO
	if err := json.Unmarshal(data, &resp); err != nil {
		return domain.LeverageDTO{}, err
	}
	return resp, nil
}

func (m *Manager) FutureQueryOrder(
	ctx context.Context, cred domain.CredentialDTO, symbol string, orderID int,
) (domain.ResponseOrderDTO, error) {
	req, weight, err := m.futureRequest.QueryOrder(ctx, cred.APIKey, cred.ApiSecret, symbol, orderID)
	if err != nil {
		return domain.ResponseOrderDTO{}, err
	}
	data, err := m.sendPersonalRequest(cred, req, weight)
	if err != nil {
		return domain.ResponseOrderDTO{}, err
	}
	var orderRep domain.ResponseOrderDTO
	if err := json.Unmarshal(data, &orderRep); err != nil {
		return domain.ResponseOrderDTO{}, err
	}
	return orderRep, nil
}

func (m *Manager) FutureCancelMultipleOrders(cred domain.CredentialDTO, symbol string, orderIDs []int) (interface{}, error) {
	req, weight, err := m.futureRequest.CancelMultipleOrders(cred.APIKey, cred.ApiSecret, symbol, orderIDs)
	if err != nil {
		return nil, err
	}
	data, err := m.sendPersonalRequest(cred, req, weight)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *Manager) FutureCandlestickDataOneHour(ctx context.Context, symbol string) ([]domain.CandlestickBarDTO, error) {
	return m.FutureCandlestickData(ctx, symbol, domain.OneHourCandlestickInterval)
}

func (m *Manager) FutureCandlestickDataFourHour(ctx context.Context, symbol string) ([]domain.CandlestickBarDTO, error) {
	return m.FutureCandlestickData(ctx, symbol, domain.FourHourCandlestickInterval)
}

func (m *Manager) FutureCandlestickData(ctx context.Context, symbol string, interval domain.CandlestickInterval) ([]domain.CandlestickBarDTO, error) {
	req, weight, err := m.futureRequest.CandlestickData(ctx, symbol, interval)
	if err != nil {
		return nil, err
	}
	data, err := m.sendRequest(req, weight)
	if err != nil {
		return nil, err
	}
	var (
		candlesticks             [][]interface{}
		candlestickBarCollection []domain.CandlestickBarDTO
	)
	if err := json.Unmarshal(data, &candlesticks); err != nil {
		return nil, err
	}
	for _, candlestick := range candlesticks {
		openTime, _ := candlestick[0].(float64)
		openPrice, _ := candlestick[1].(string)
		highPrice, _ := candlestick[2].(string)
		lowPrice, _ := candlestick[3].(string)
		closePrice, _ := candlestick[4].(string)
		volume, _ := candlestick[5].(string)
		closeTime, _ := candlestick[6].(float64)
		quoteAssetVolume, _ := candlestick[7].(string)
		numberOfTrades, _ := candlestick[8].(float64)
		takerBuyBaseAssetVolume, _ := candlestick[9].(string)
		takerBuyQuoteAssetVolume, _ := candlestick[10].(string)
		ignore, _ := candlestick[11].(string)

		candlestickBarCollection = append(candlestickBarCollection, domain.CandlestickBarDTO{
			OpenTime:                 openTime,
			OpenPrice:                openPrice,
			HighPrice:                highPrice,
			LowPrice:                 lowPrice,
			ClosePrice:               closePrice,
			Volume:                   volume,
			CloseTime:                closeTime,
			QuoteAssetVolume:         quoteAssetVolume,
			NumberOfTrades:           numberOfTrades,
			TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
			Ignore:                   ignore,
		})
	}
	return candlestickBarCollection, nil
}

func (m *Manager) executeRequest(req *http.Request, weight int, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	resp, err := m.sendRequest(req, weight)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *Manager) sendPersonalRequest(cred domain.CredentialDTO, req *http.Request, weight int) ([]byte, error) {
	key := fmt.Sprintf("%s-%s", cred.APIKey, cred.ApiSecret)
	if m.personalSenders[key] == nil {
		m.personalSenders[key] = sender.NewPersonal(m.sender)
	}
	resp, err := m.personalSenders[key].Send(req, weight)
	if err != nil {
		return nil, err
	}
	body, errRep := m.handlerResponse(resp)
	if errRep != nil {
		return nil, errRep
	}
	return body, nil
}

func (m *Manager) sendRequest(req *http.Request, weight int) ([]byte, error) {
	resp, err := m.sender.Send(req, weight)
	if err != nil {
		return nil, err
	}
	body, errRep := m.handlerResponse(resp)
	if errRep != nil {
		return nil, errRep
	}
	return body, nil
}

func (m *Manager) handlerResponse(resp *http.Response) ([]byte, error) {
	if resp.Body == nil {
		return nil, errors.New("empty body response")
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error handle response")
	}
	return body, nil
}

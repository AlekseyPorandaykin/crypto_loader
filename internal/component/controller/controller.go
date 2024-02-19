package controller

import (
	"errors"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/candlestick"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/order"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	host         string
	priceStorage domain.PriceStorage
	order        *order.Order
	candlestick  *candlestick.Candlestick
}

func NewController(host string, priceStorage domain.PriceStorage, order *order.Order, candlestick *candlestick.Candlestick) *Controller {
	return &Controller{
		host:         host,
		priceStorage: priceStorage,
		order:        order,
		candlestick:  candlestick,
	}
}

func (app *Controller) RegistrationPageRoute(g *echo.Group) {
	g.GET("/prices", app.prices)
	g.GET("/price/:symbol", app.symbolPrice)
	g.GET("/price/exchange/:exchange", app.exchangePrice)
	g.POST("/order", app.createOrder)
	g.GET("/snapshot/:exchange/:symbol", app.snapshot)
	g.GET("/candlesticks/:interval/:exchange/:symbol", app.candlesticks)
}

func (app *Controller) prices(c echo.Context) error {
	prices, err := app.priceStorage.LastPrices(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, prices)
}

func (app *Controller) symbolPrice(c echo.Context) error {
	symbol := c.Param("symbol")
	if symbol == "" {
		return errors.New("empty symbol")
	}
	prices, err := app.priceStorage.SymbolPrice(c.Request().Context(), symbol)
	if err != nil {
		return err
	}
	if len(prices) == 0 {
		return c.JSON(http.StatusNotFound, prices)
	}
	return c.JSON(http.StatusOK, prices)
}

func (app *Controller) exchangePrice(c echo.Context) error {
	exchange := c.Param("exchange")
	if exchange == "" {
		return errors.New("empty exchange")
	}
	prices, err := app.priceStorage.ExchangePrice(c.Request().Context(), exchange)
	if err != nil {
		return err
	}
	if len(prices) == 0 {
		return c.JSON(http.StatusNotFound, prices)
	}
	return c.JSON(http.StatusOK, prices)
}

func (app *Controller) createOrder(c echo.Context) error {
	req := dto.FutureOrderRequest{}
	if err := (&echo.DefaultBinder{}).BindBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	orders, err := app.order.CreateFutureOrder(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, orders)
}

func (app *Controller) snapshot(c echo.Context) error {
	symbol := c.Param("symbol")
	if symbol == "" {
		return errors.New("empty symbol")
	}
	exchange := c.Param("exchange")
	if symbol == "" {
		return errors.New("empty exchange")
	}
	snapshot, err := app.candlestick.SymbolSnapshot(c.Request().Context(), exchange, symbol)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, snapshot)
}

func (app *Controller) candlesticks(c echo.Context) error {
	symbol := c.Param("symbol")
	if symbol == "" {
		return errors.New("empty symbol")
	}
	exchange := c.Param("exchange")
	if symbol == "" {
		return errors.New("empty exchange")
	}
	interval := c.Param("interval")
	if symbol == "" {
		return errors.New("empty interval")
	}
	snapshot, err := app.candlestick.Candlesticks(
		c.Request().Context(), exchange, symbol, domain.CandlestickInterval(interval),
	)
	if err != nil {
		return err
	}
	if len(snapshot) == 0 {
		return c.JSON(http.StatusNotFound, snapshot)
	}
	return c.JSON(http.StatusOK, snapshot)
}

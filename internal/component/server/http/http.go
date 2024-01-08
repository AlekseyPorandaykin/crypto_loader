package http

import (
	"errors"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/candlestick"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/order"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	host         string
	priceStorage domain.PriceStorage
	e            *echo.Echo
	order        *order.Order
	candlestick  *candlestick.Candlestick
}

func NewServer(host string, priceStorage domain.PriceStorage, order *order.Order, candlestick *candlestick.Candlestick) *Server {
	return &Server{
		host:         host,
		priceStorage: priceStorage,
		order:        order,
		candlestick:  candlestick,

		e: echo.New(),
	}
}

func (s *Server) Run() error {
	s.e.HideBanner = true
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORS())
	s.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				zap.L().Error("error http execute", zap.Error(err), zap.String("url", c.Request().URL.String()))
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return nil
		}
	})
	s.e.GET("/prices", s.prices)
	s.e.GET("/price/:symbol", s.symbolPrice)
	s.e.POST("/order", s.createOrder)
	s.e.GET("/snapshot/:exchange/:symbol", s.snapshot)
	s.e.GET("/candlesticks/:interval/:exchange/:symbol", s.candlesticks)
	return s.e.Start(s.host)
}

func (s *Server) prices(c echo.Context) error {
	prices, err := s.priceStorage.LastPrices(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, prices)
}

func (s *Server) symbolPrice(c echo.Context) error {
	symbol := c.Param("symbol")
	if symbol == "" {
		return errors.New("empty symbol")
	}
	prices, err := s.priceStorage.SymbolPrice(c.Request().Context(), symbol)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, prices)
}

func (s *Server) createOrder(c echo.Context) error {
	req := dto.FutureOrderRequest{}
	if err := (&echo.DefaultBinder{}).BindBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	orders, err := s.order.CreateFutureOrder(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, orders)
}

func (s *Server) snapshot(c echo.Context) error {
	symbol := c.Param("symbol")
	if symbol == "" {
		return errors.New("empty symbol")
	}
	exchange := c.Param("exchange")
	if symbol == "" {
		return errors.New("empty exchange")
	}
	snapshot, err := s.candlestick.SymbolSnapshot(c.Request().Context(), exchange, symbol)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, snapshot)
}

func (s *Server) candlesticks(c echo.Context) error {
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
	snapshot, err := s.candlestick.Candlesticks(
		c.Request().Context(), exchange, symbol, domain.CandlestickInterval(interval),
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, snapshot)
}
func (s *Server) Close() {
	_ = s.e.Close()
}

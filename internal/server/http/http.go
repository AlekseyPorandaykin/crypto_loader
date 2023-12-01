package http

import (
	"errors"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/order"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Server struct {
	host         string
	priceStorage domain.PriceStorage
	e            *echo.Echo
	order        *order.Order
}

func NewServer(host string, priceStorage domain.PriceStorage, order *order.Order) *Server {
	return &Server{
		host:         host,
		priceStorage: priceStorage,
		order:        order,

		e: echo.New(),
	}
}

func (s *Server) Run() error {
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORS())
	s.e.GET("/prices", func(c echo.Context) error {
		prices, err := s.priceStorage.LastPrices(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, prices)
	})
	s.e.GET("/price/:symbol", func(c echo.Context) error {
		symbol := c.Param("symbol")
		if symbol == "" {
			return errors.New("empty symbol")
		}
		prices, err := s.priceStorage.SymbolPrice(c.Request().Context(), symbol)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, prices)
	})
	s.e.POST("/order", func(c echo.Context) error {
		req := dto.FutureOrderRequest{}
		if err := (&echo.DefaultBinder{}).BindBody(c, &req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		orders, err := s.order.CreateFutureOrder(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, orders)
	})
	return s.e.Start(s.host)
}

func (s *Server) Close() {
	_ = s.e.Close()
}

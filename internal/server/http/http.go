package http

import (
	"errors"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	host         string
	priceStorage domain.PriceStorage
	e            *echo.Echo
}

func NewServer(host string, priceStorage domain.PriceStorage) *Server {
	return &Server{
		host:         host,
		priceStorage: priceStorage,

		e: echo.New(),
	}
}

func (s *Server) Run() error {
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
	return s.e.Start(s.host)
}

func (s *Server) Close() {
	_ = s.e.Close()
}

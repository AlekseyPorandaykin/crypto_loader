package grpc

import (
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/server/grpc/specification"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"net"
)

type Server struct {
	addr    string
	handler specification.EventServiceServer
	serv    *grpc.Server
}

func NewServer(storage domain.PriceStorage, addr string) *Server {
	return &Server{
		handler: NewHandler(storage),
		serv:    grpc.NewServer(),
		addr:    addr,
	}
}

func (s *Server) Start() error {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	specification.RegisterEventServiceServer(s.serv, s.handler)
	return s.serv.Serve(l)
}

func (s *Server) Close() {
	s.serv.Stop()
}

package app

import (
	"fmt"
	aiservice "github.com/shamank/ai-marketplace-stats-service/internal/delivery/grpc/ai-service"
	"github.com/shamank/ai-marketplace-stats-service/internal/delivery/grpc/statistic"
	"github.com/shamank/ai-marketplace-stats-service/internal/service"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type GRPCServer struct {
	log  *slog.Logger
	gRPC *grpc.Server
	port int
}

func NewGRPCServer(log *slog.Logger, services *service.Service, port int) *GRPCServer {

	gRPCServer := grpc.NewServer()

	statistic.RegisterServerAPI(gRPCServer, services.StatisticService)
	aiservice.RegisterServerAPI(gRPCServer, services.AIService)

	return &GRPCServer{
		log:  log,
		gRPC: gRPCServer,
		port: port,
	}
}

func (s *GRPCServer) Run() error {

	s.log.Info("Starting gRPC server on port " + fmt.Sprintf("%d", s.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	if err := s.gRPC.Serve(l); err != nil {
		s.log.Error("Failed to start gRPC server: " + err.Error())
		return err
	}

	return nil
}

func (s *GRPCServer) Stop() {
	s.log.Info("Stopping gRPC server")
	s.gRPC.GracefulStop()
}

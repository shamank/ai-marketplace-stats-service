package ai_service

import (
	"context"
	statsv1 "github.com/shamank/ai-marketplace-protos/gen/go/stats-service"
	"github.com/shamank/ai-marketplace-stats-service/internal/domain/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AIServiceSVC interface {
	Create(ctx context.Context, aiService models.AIServiceCreate) (string, error)
}

type serverAPI struct {
	statsv1.UnimplementedAIServiceServer
	svc AIServiceSVC
}

func RegisterServerAPI(gRPC *grpc.Server, svc AIServiceSVC) {
	statsv1.RegisterAIServiceServer(gRPC, &serverAPI{svc: svc})
}

func (s *serverAPI) CreateAIService(ctx context.Context, req *statsv1.CreateServiceRequest) (*statsv1.CreateServiceResponse, error) {
	if req.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}

	var description *string

	if req.GetDescription() != "" {
		desc := req.GetDescription()
		description = &desc
	}

	serviceUID, err := s.svc.Create(ctx, models.AIServiceCreate{
		Title:       req.GetTitle(),
		Description: description,
		Price:       int(req.GetPrice()),
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &statsv1.CreateServiceResponse{
		ServiceUid: serviceUID,
	}, nil
}

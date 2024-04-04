package statistic

import (
	"context"
	statsv1 "github.com/shamank/ai-marketplace-protos/gen/go/stats-service"
	"github.com/shamank/ai-marketplace-stats-service/internal/domain/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate mockery --name=StatisticService
type StatisticService interface {
	CreateService(ctx context.Context, create models.AIServiceCreate) (string, error)
	Call(ctx context.Context, AIServiceUID string, userUID string) error
	GetCalls(ctx context.Context, filter models.StatisticFilter) ([]models.StatisticRead, error)
}

type serverAPI struct {
	statsv1.UnimplementedStatisticServiceServer
	svc StatisticService
}

func RegisterServerAPI(gRPC *grpc.Server, svc StatisticService) {
	statsv1.RegisterStatisticServiceServer(gRPC, &serverAPI{svc: svc})
}

func (s *serverAPI) Create(ctx context.Context, req *statsv1.CreateAIServiceRequest) (*statsv1.CreateAIServiceResponse, error) {

	input := models.AIServiceCreate{}

	if req.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "title is required")
	}
	input.Title = req.GetTitle()

	if req.GetDescription() != "" {
		description := req.GetDescription()
		input.Description = &description
	}

	if req.GetPrice() == 0 {
		return nil, status.Error(codes.InvalidArgument, "price is required")
	}
	input.Price = req.GetPrice()

	serviceUID, err := s.svc.CreateService(ctx, input)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &statsv1.CreateAIServiceResponse{
		ServiceUid: serviceUID,
	}, nil
}

func (s *serverAPI) Call(ctx context.Context, req *statsv1.CallRequest) (*statsv1.CallResponse, error) {

	if req.GetServiceUid() == "" {
		return nil, status.Error(codes.InvalidArgument, "service_uid is required")
	}
	if req.GetUserUid() == "" {
		return nil, status.Error(codes.InvalidArgument, "user_uid is required")
	}

	err := s.svc.Call(ctx, req.GetServiceUid(), req.GetUserUid())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &statsv1.CallResponse{
		Message: "ok",
	}, nil
}

func (s *serverAPI) GetCalls(ctx context.Context, req *statsv1.GetCallsRequest) (*statsv1.GetCallsResponse, error) {
	filter := handleStatisticFilter(req)

	calls, err := s.svc.GetCalls(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var callsList []*statsv1.Calls

	for _, call := range calls {
		callsList = append(callsList, &statsv1.Calls{
			UserUid:    call.UserUID,
			ServiceUid: call.AIServiceUID,
			Count:      call.Count,
			Amount:     call.FullAmount,
		})
	}

	return &statsv1.GetCallsResponse{Calls: callsList}, nil
}

func handleStatisticFilter(req *statsv1.GetCallsRequest) models.StatisticFilter {
	filter := models.StatisticFilter{}
	if userUID := req.GetUserUid(); userUID != "" {
		filter.UserUID = &userUID
	}

	if serviceUID := req.GetServiceUid(); serviceUID != "" {
		filter.AIServiceUID = &serviceUID
	}

	if order := req.GetOrder().String(); order != "" {
		filter.Order = &order
	}

	if pageNumber := req.GetPageNumber(); pageNumber != 0 {
		filter.PageNumber = &pageNumber
	}

	if pageSize := req.GetPageSize(); pageSize != 0 {
		filter.PageSize = &pageSize
	}
	return filter
}

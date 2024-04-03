package statistic

import (
	"context"
	statsv1 "github.com/shamank/ai-marketplace-protos/gen/go/stats-service"
	"github.com/shamank/ai-marketplace-stats-service/internal/domain/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StatsServiceSVC interface {
	GetStats(ctx context.Context, filter models.StatisticFilter) ([]models.StatisticRead, error)
	SetStat(ctx context.Context, userUID string, serviceUID string) (string, error)
}

type serverAPI struct {
	statsv1.UnimplementedStatsServiceServer
	svc StatsServiceSVC
}

func RegisterServerAPI(gRPC *grpc.Server, svc StatsServiceSVC) {
	statsv1.RegisterStatsServiceServer(gRPC, &serverAPI{svc: svc})
}

func (s *serverAPI) GetStats(ctx context.Context, req *statsv1.GetStatsRequest) (*statsv1.GetStatsResponse, error) {
	filter := handleStatisticFilter(req)

	stats, err := s.svc.GetStats(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var statsList []*statsv1.Stats

	for _, stat := range stats {
		statsList = append(statsList, &statsv1.Stats{
			UserUid:    stat.UserUID,
			ServiceUid: stat.AIServiceUID,
			Count:      stat.Count,
			Amount:     stat.FullAmount,
		})
	}

	return &statsv1.GetStatsResponse{Stats: statsList}, nil
}

func (s *serverAPI) SetStat(ctx context.Context, req *statsv1.SetStatRequest) (*statsv1.SetStatResponse, error) {
	if req.GetUserUid() == "" {
		return nil, status.Error(codes.InvalidArgument, "user_uid is required")
	}
	if req.GetServiceUid() == "" {
		return nil, status.Error(codes.InvalidArgument, "service_uid is required")
	}
	uid, err := s.svc.SetStat(ctx, req.GetUserUid(), req.GetServiceUid())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &statsv1.SetStatResponse{StatUid: uid}, nil
}

func handleStatisticFilter(req *statsv1.GetStatsRequest) models.StatisticFilter {
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
		filter.PageSize = &pageNumber
	}

	if pageSize := req.GetPageSize(); pageSize != 0 {
		filter.PageSize = &pageSize
	}
	return filter
}

package statistic

import (
	"context"
	statsv1 "github.com/shamank/ai-marketplace-protos/gen/go/stats-service"
	"github.com/shamank/ai-marketplace-stats-service/internal/delivery/grpc/statistic/mocks"
	"github.com/shamank/ai-marketplace-stats-service/internal/domain/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	type TestCase struct {
		input  *statsv1.CreateAIServiceRequest
		output *statsv1.CreateAIServiceResponse
		err    error
	}

	testCases := []TestCase{

		{
			input: &statsv1.CreateAIServiceRequest{
				Title:       "test",
				Description: "test",
				Price:       10,
			},
			output: &statsv1.CreateAIServiceResponse{
				ServiceUid: "123",
			},
			err: nil,
		},
	}

	for _, testCase := range testCases {
		model := models.AIServiceCreate{
			Title:       testCase.input.Title,
			Description: &testCase.input.Description,
			Price:       testCase.input.Price,
		}
		svc := &mocks.StatisticService{}

		ctx := context.Background()

		srv := serverAPI{svc: svc}

		svc.On("CreateService", ctx, model).Return(testCase.output.ServiceUid, testCase.err)

		res, err := srv.Create(ctx, testCase.input)

		assert.Equal(t, testCase.output, res)
		assert.Equal(t, testCase.err, err)
	}

}

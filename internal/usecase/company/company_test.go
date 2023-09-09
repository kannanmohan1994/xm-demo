package company

import (
	"errors"
	"testing"
	"xm/consts"
	"xm/internal/entity/models"
	"xm/internal/entity/request"
	"xm/logger"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var employeeCount int = 500
var isRegistered bool = false

func TestService_CreateCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		request       *request.CreateCompanyRequest
		err           map[string]error
		repoCallCount int
	}{
		{
			name: "should_return_error_for_repo_call",
			request: &request.CreateCompanyRequest{
				Name:          "xm",
				Description:   "Forex & CFD Trading on Stocks, Indices, Oil, Gold by XM",
				EmployeeCount: &employeeCount,
				IsRegistered:  &isRegistered,
				Type:          consts.CompanyTypeCorporations,
			},
			err:           map[string]error{"CreateCompany": errors.New("repo call fail")},
			repoCallCount: 1,
		},
		{
			name: "success",
			request: &request.CreateCompanyRequest{
				Name:          "xm",
				Description:   "Forex & CFD Trading on Stocks, Indices, Oil, Gold by XM",
				EmployeeCount: &employeeCount,
				IsRegistered:  &isRegistered,
				Type:          consts.CompanyTypeCorporations,
			},
			err:           nil,
			repoCallCount: 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMockCompanyRepository(ctrl)
			logger := logger.NewNoop()
			repo.EXPECT().CreateCompany(gomock.Any()).Return(&models.Company{}, tt.err["CreateCompany"]).Times(tt.repoCallCount)
			s := InitCompanyUsecase(repo, logger)

			_, err := s.CreateCompany(tt.request)
			if tt.err != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

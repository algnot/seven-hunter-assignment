package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"seven-hunter-assignment/question-3/logic"
	"seven-hunter-assignment/question-3/model"
	"testing"
)

type MockRequest struct {
	mock.Mock
}

func (m *MockRequest) GetBeefFromApi() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestGetBeefSummary(t *testing.T) {
	mockRequest := new(MockRequest)
	mockRequest.On("GetBeefFromApi").Return("beef pork beef chicken", nil)

	bLogic := logic.GetBeefLogic(mockRequest)
	result, err := bLogic.GetBeefSummary()

	assert.NoError(t, err)
	assert.ElementsMatch(t, []model.Beef{
		{Name: "beef", Count: 2},
		{Name: "pork", Count: 1},
		{Name: "chicken", Count: 1},
	}, result)
}

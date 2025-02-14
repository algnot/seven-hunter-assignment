package logic

import (
	"seven-hunter-assignment/question-3/model"
	"seven-hunter-assignment/question-3/util/config"
	"seven-hunter-assignment/question-3/util/request"
	"strings"
)

type BeefLogic interface {
	GetBeefSummary() ([]model.Beef, error)
}

type beefLogic struct {
	config config.AppConfig
}

func (b beefLogic) GetBeefSummary() ([]model.Beef, error) {
	var result []model.Beef

	beefFromApi, err := request.GetBeefFromApi()
	if err != nil {
		return nil, err
	}

	split := strings.Split(beefFromApi, " ")
	counts := make(map[string]int64)

	for _, str := range split {
		counts[str]++
	}

	for name, count := range counts {
		result = append(result, model.Beef{Name: name, Count: count})
	}

	return result, nil
}

func GetBeefLogic() BeefLogic {
	return &beefLogic{}
}

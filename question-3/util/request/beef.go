package request

import (
	"fmt"
	"io"
	"net/http"
	"seven-hunter-assignment/question-3/util/config"
)

func GetBeefFromApi() (string, error) {
	appConfig := config.GetAppConfig()
	url := fmt.Sprintf("%s/api/?type=meat-and-filler&paras=99&format=text", appConfig.CommonConfig.BeefApi)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err

	}

	return string(body), nil
}

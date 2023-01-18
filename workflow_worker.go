package pocworkflowworker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	TypeFirstEndpoint  = "first_endpoint"
	TypeSecondEndpoint = "second_endpoint"
	TypeThirdEndpoint  = "third_endpoint"

	FirstEndpointFormat  = "api/v1/ktp/first/%d"
	SecondEndpointFormat = "api/v1/ktp/second/%d"
	ThirdEndpointFormat  = "api/v1/ktp/third/%d"
)

type (
	PubsubMsg struct {
		Data []byte `json:"data"`
	}

	RawMsg struct {
		Data json.RawMessage `json:"data"`
	}

	CommonPublishReq struct {
		ID       uint64 `json:"id"`
		RuleName string `json:"rule_name"`
		Type     string `json:"type"`
	}
)

func WorkflowWorker(ctx context.Context, e PubsubMsg) error {
	// init httpClient
	httpClient := &http.Client{}

	// get worker url
	workerUrl := os.Getenv("WORKER_SERVICE_URL")
	// unmarshal si data dr pub/sub nya
	var req CommonPublishReq
	err := json.Unmarshal(e.Data, &req)
	if err != nil {
		logrus.Error("getting error when unmarshal data, err: %v", err)
		return err
	}

	logrus.Info("incoming data =>>", req)

	// decision utk hit endpoint mana
	if req.Type == TypeFirstEndpoint {
		// call ke first endpoint
		return sendRequest(httpClient, workerUrl, fmt.Sprintf(FirstEndpointFormat, req.ID))
	}
	if req.Type == TypeSecondEndpoint {
		// call ke second endpoint
		return sendRequest(httpClient, workerUrl, fmt.Sprintf(SecondEndpointFormat, req.ID))
	}
	if req.Type == TypeThirdEndpoint {
		// call ke third endpoint
		return sendRequest(httpClient, workerUrl, fmt.Sprintf(ThirdEndpointFormat, req.ID))
	}

	logrus.Info("do nothing.... bloookkk")
	return nil
}

func sendRequest(client *http.Client, url, endpoint string) error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", url, endpoint), nil)
	if err != nil {
		logrus.Errorf("getting error when create new request, err: %v \n", err)
		return err
	}
	req.Header.Add("x-workflow-name", "ktp_register_v1")

	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("getting error when do request, err: %v \n", err)
		return err
	}
	defer resp.Body.Close()

	logrus.Info("http request completed...")

	return nil
}

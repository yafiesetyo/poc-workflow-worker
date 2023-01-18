package pocworkflowworker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	TypeFirstEndpoint  = "first_endpoint"
	TypeSecondEndpoint = "second_endpoint"
	TypeThirdEndpoint  = "third_endpoint"
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
	// unmarshal si data dr pub/sub nya
	var raw RawMsg
	err := json.Unmarshal(e.Data, &raw)
	if err != nil {
		logrus.Error("getting error when unmarshal raw data, err: %v", err)
		return err
	}

	logrus.Info("check raw msg", raw.Data)

	var req CommonPublishReq
	err = json.Unmarshal(raw.Data, &req)
	if err != nil {
		logrus.Error("getting error when unmarshal common request, err: %v", err)
		return err
	}

	fmt.Println("Incoming request", req)

	// decision utk hit endpoint mana
	if req.Type == TypeFirstEndpoint {
		// call ke first endpoint
		return nil
	}
	if req.Type == TypeSecondEndpoint {
		// call ke second endpoint
		return nil
	}
	if req.Type == TypeThirdEndpoint {
		// call ke third endpoint
		return nil
	}

	return nil
}

package pocworkflowworker

import (
	"context"
	"encoding/json"

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
	var req CommonPublishReq
	err := json.Unmarshal(e.Data, &req)
	if err != nil {
		logrus.Error("getting error when unmarshal data, err: %v", err)
		return err
	}

	logrus.Info("check msg ", req, " ", string(e.Data))

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

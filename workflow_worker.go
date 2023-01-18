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
	CommonPublishReq struct {
		ID       uint64 `json:"id"`
		RuleName string `json:"rule_name"`
		Type     string `json:"type"`
	}
)

func WorkflowWorker(ctx context.Context, e []byte) error {
	// unmarshal si data dr pub/sub nya
	var req CommonPublishReq
	err := json.Unmarshal(e, &req)
	if err != nil {
		logrus.Error("getting error when unmarshal, err: %v", err)
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

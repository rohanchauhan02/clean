package schemas

import (
	"fmt"
	"strings"
)

type (
	SQSActionsRequest struct {
		Operation string      `json:"operation"`
		Data      interface{} `json:"data"`
	}

	SQSActionsData struct {
		Queue string `json:"queue"`
		Delay int    `json:"delay"`
	}

	SQSActionsEnv struct {
		Local SQSActionsData
		Dev   SQSActionsData
		Prod  SQSActionsData
		UAT   SQSActionsData
	}
)

const (
	OPERATION_QUOTATION_UPDATE_STATUS                                    = "UPDATE_STATUS"
	OPERATION_FINANCE_UPDATE_STOPLOSS                                    = "UPDATE_STOPLOSS"
	OPERATION_FINANCE_UPDATE_STOPLOSS_INCREMENT                          = "UPDATE_STOPLOSS_INCREMENT"
	OPERATION_FINANCE_UPDATE_STOPLOSS_DECREMENT                          = "UPDATE_STOPLOSS_DECREMENT"
	OPERATION_FINANCE_UPDATE_STOPLOSS_CLAIMED_PREMIUM_INCREMENT          = "UPDATE_STOPLOSS_CLAIMED_PREMIUM_INCREMENT"
	OPERATION_FINANCE_UPDATE_STOPLOSS_CLAIMED_PREMIUM_DECREMENT          = "UPDATE_STOPLOSS_CLAIMED_PREMIUM_DECREMENT"
	OPERATION_FINANCE_UPDATE_STOPLOSS_HISTORY                            = "UPDATE_STOPLOSS_CLAIMED_STOPLOSS_HISTORY"
	OPERATION_FINANCE_PENDING_CLAIMED_PREMIUM_DECREMENT                  = "UPDATE_STOPLOSS_PENDING_CLAIMED_PREMIUM_DECREMENT"
	OPERATION_FINANCE_PENDING_CLAIMED_PREMIUM_DECREMENT_FOR_CANCEL_CLAIM = "UPDATE_STOPLOSS_PENDING_CLAIMED_PREMIUM_DECREMENT_FOR_CANCEL_CLAIM"
)

var (
	SQSActionsValue map[string]SQSActionsEnv = map[string]SQSActionsEnv{
		"QUOTATION_UPDATE": {
			Local: SQSActionsData{
				Queue: "local-quotation-std-svc-update.fifo",
				Delay: 0,
			},
			Dev: SQSActionsData{
				Queue: "dev-quotation-std-svc-update.fifo",
				Delay: 0,
			},
			UAT: SQSActionsData{
				Queue: "uat-quotation-std-svc-update.fifo",
				Delay: 0,
			},
			Prod: SQSActionsData{
				Queue: "prod-quotation-std-svc-update.fifo",
				Delay: 0,
			},
		},
		"FINANCE_UPDATE": {
			Local: SQSActionsData{
				Queue: "local-finance-std-svc-update.fifo",
				Delay: 1,
			},
			Dev: SQSActionsData{
				Queue: "dev-finance-std-svc-update.fifo",
				Delay: 1,
			},
			UAT: SQSActionsData{
				Queue: "uat-finance-std-svc-update.fifo",
				Delay: 1,
			},
			Prod: SQSActionsData{
				Queue: "prod-finance-std-svc-update.fifo",
				Delay: 1,
			},
		},
	}
)

func GetSQSActions(env string, keyActions string) (result SQSActionsData, err error) {
	for k, v := range SQSActionsValue {
		if strings.EqualFold(k, keyActions) {
			switch strings.ToUpper(env) {
			case "LOCAL":
				return v.Local, nil
			case "DEV":
				return v.Dev, nil
			case "UAT":
				return v.UAT, nil
			case "PROD":
				return v.Prod, nil
			default:
				return SQSActionsData{}, fmt.Errorf("cannot find env: %s", env)
			}
		}
	}
	return SQSActionsData{}, fmt.Errorf("cannot find keyActions: %s", keyActions)
}

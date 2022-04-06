package integrationtest

import (
	"testing"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
)

func genCfxTestConfig() rpcTestConfig {

	var rpc2Func map[string]string = map[string]string{
		"cfx_getStatus":                         "GetStatus",
		"cfx_gasPrice":                          "GetGasPrice",
		"cfx_getNextNonce":                      "GetNextNonce",
		"cfx_epochNumber":                       "GetEpochNumber",
		"cfx_getBalance":                        "GetBalance",
		"cfx_getCode":                           "GetCode",
		"cfx_getBestBlockHash":                  "GetBestBlockHash",
		"cfx_getConfirmationRiskByHash":         "GetRawBlockConfirmationRisk",
		"cfx_sendRawTransaction":                "SendRawTransaction",
		"cfx_call":                              "Call",
		"cfx_getLogs":                           "GetLogs",
		"cfx_getTransactionByHash":              "GetTransactionByHash",
		"cfx_estimateGasAndCollateral":          "EstimateGasAndCollateral",
		"cfx_getBlocksByEpoch":                  "GetBlocksByEpoch",
		"cfx_getTransactionReceipt":             "GetTransactionReceipt",
		"cfx_getAdmin":                          "GetAdmin",
		"cfx_getSponsorInfo":                    "GetSponsorInfo",
		"cfx_getStakingBalance":                 "GetStakingBalance",
		"cfx_getCollateralForStorage":           "GetCollateralForStorage",
		"cfx_getStorageAt":                      "GetStorageAt",
		"cfx_getStorageRoot":                    "GetStorageRoot",
		"cfx_getBlockByHashWithPivotAssumption": "GetBlockByHashWithPivotAssumption",
		"cfx_checkBalanceAgainstTransaction":    "CheckBalanceAgainstTransaction",
		"cfx_getSkippedBlocksByEpoch":           "GetSkippedBlocksByEpoch",
		"cfx_getAccount":                        "GetAccountInfo",
		"cfx_getInterestRate":                   "GetInterestRate",
		"cfx_getAccumulateInterestRate":         "GetAccumulateInterestRate",
		"cfx_getBlockRewardInfo":                "GetBlockRewardInfo",
		"cfx_clientVersion":                     "GetClientVersion",
		"cfx_getDepositList":                    "GetDepositList",
		"cfx_getVoteList":                       "GetVoteList",
		"cfx_getSupplyInfo":                     "GetSupplyInfo",
		"trace_block":                           "GetBlockTraces",
		"trace_filter":                          "FilterTraces",
		"trace_transaction":                     "GetTransactionTraces",
		"cfx_getPoSRewardByEpoch":               "GetPosRewardByEpoch",
		"cfx_getAccountPendingInfo":             "GetAccountPendingInfo",
		"cfx_getAccountPendingTransactions":     "GetAccountPendingTransactions",
		"cfx_getPoSEconomics":                   "GetPoSEconomics",
		"cfx_openedMethodGroups":                "GetOpenedMethodGroups",
	}

	var rpc2FuncSelector map[string]func(params []interface{}) (string, []interface{}) = map[string]func(params []interface{}) (string, []interface{}){
		"cfx_getBlockByEpochNumber": func(params []interface{}) (string, []interface{}) {
			if params[1] == false {
				return "GetBlockSummaryByEpoch", []interface{}{params[0]}
			}
			return "GetBlockByEpoch", []interface{}{params[0]}
		},

		"cfx_getBlockByHash": func(params []interface{}) (string, []interface{}) {
			if params[1] == false {
				return "GetBlockSummaryByHash", []interface{}{params[0]}
			}
			return "GetBlockByHash", []interface{}{params[0]}
		},

		"cfx_getBlockByBlockNumber": func(params []interface{}) (string, []interface{}) {
			if params[1] == false {
				return "GetBlockSummaryByBlockNumber", []interface{}{params[0]}
			}
			return "GetBlockByBlockNumber", []interface{}{params[0]}
		},

		"cfx_getAccountPendingTransactions": func(params []interface{}) (string, []interface{}) {
			params = append(params, nil, nil)
			return "GetAccountPendingTransactions", params[:3]
		},
	}

	// ignoreRpc priority is higher than onlyTestRpc
	var ignoreRpc map[string]bool = map[string]bool{
		// "cfx_getBlockByEpochNumber": true,
		"cfx_getLogs": true,
	}

	// onlyTestRpc priority is lower than ignoreRpc
	var onlyTestRpc map[string]bool = map[string]bool{
		// "cfx_getLogs": true,
	}

	return rpcTestConfig{
		examplesUrl: "https://raw.githubusercontent.com/Conflux-Chain/jsonrpc-spec/main/src/cfx/examples.json",
		client:      sdk.MustNewClient("http://47.93.101.243"),

		rpc2Func:         rpc2Func,
		rpc2FuncSelector: rpc2FuncSelector,
		ignoreRpc:        ignoreRpc,
		onlyTestRpc:      onlyTestRpc,
	}

}

// TODO: Open after rpc mock server ready
func _TestClientCFX(t *testing.T) {
	cfxaddress.SetConfig(cfxaddress.Config{
		AddressStringVerbose: true,
	})

	config := genCfxTestConfig()
	doClientTest(t, config)
}
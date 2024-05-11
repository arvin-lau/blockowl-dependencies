package chainExplorer

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

type EstimateGasResp struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data struct {
		GasFeeCap decimal.Decimal `json:"gas_fee_cap"`
		GasPrice  decimal.Decimal `json:"gas_price"`
	} `json:"data"`
}

func EstimateGas(chain string) (resp EstimateGasResp, err error) {
	url := fmt.Sprintf(GetChainExplorerHost()+estimateGasPath, chain)
	_, _, errSli := gorequest.New().
		Get(url).
		EndStruct(&resp)
	if len(errSli) != 0 {
		errMsg := fmt.Sprintf("chain explorer EstimateGas err: %v reqUrl: %v resp: %v", errSli, url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	if resp.Code != SuccessfulCode {
		errMsg := fmt.Sprintf("chainExplorer.EstimateGas err balanceResp.Code != chainExplorer.SuccessfulCode,req: %v resp: %+v", url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	return resp, nil
}

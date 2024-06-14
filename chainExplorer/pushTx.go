package chainExplorer

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/zerolog/log"
)

type PushTxResp struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data string `json:"data"`
}

// PushTx push tx to chain explorer
func PushTx(chain string, pushTxDate string) (resp PushTxResp, err error) {
	url := GetChainExplorerHost() + pushTxPath
	param := fmt.Sprintf(`{"chainShortName":"%v", "signedTx":"%v"}`, chain, pushTxDate)
	_, _, errSli := gorequest.New().Post(url).
		Set("Content-Type", "application/json").
		//SendStruct(reqParam).
		Send(param).
		EndStruct(&resp)
	if len(errSli) != 0 {
		errMsg := fmt.Sprintf("chain explorer PushTx err: %v reqUrl: %v param: %v resp: %v", errSli, url, param, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(fmt.Sprintf("%v", errSli))
	}
	if resp.Code != SuccessfulCode {
		errMsg := fmt.Sprintf("chainExplorer.PushTx err resp.Code != SuccessfulCode,url: %v param: %v resp: %+v", url, param, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(resp.Msg)
	}
	return resp, nil
}

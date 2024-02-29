package chainExplorer

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/zerolog/log"
)

type pushTxResp struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data string `json:"data"`
}

// PushTx push tx to chain explorer
func PushTx(chain string, pushTxDate string) (resp pushTxResp, err error) {
	url := GetChainExplorerHost() + pushTxPath
	param := fmt.Sprintf(`{"chainShortName":"%v", "signedTx":"%v"}`, chain, pushTxDate)
	log.Info().Msg("param: " + param)
	_, _, errSli := gorequest.New().Post(url).
		Set("Content-Type", "application/json").
		//SendStruct(reqParam).
		Send(param).
		EndStruct(&resp)
	if len(errSli) != 0 {
		errMsg := fmt.Sprintf("chain explorer PushTx err: %v reqUrl: %v resp: %v", err, url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	if resp.Code != SuccessfulCode {
		errMsg := fmt.Sprintf("chainExplorer.PushTx err resp.Code != SuccessfulCode,url: %v param: %v resp: %+v", url, param, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	return resp, nil
}

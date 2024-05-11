package chainExplorer

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/zerolog/log"
)

type GetAccountInfoResp struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data struct {
		Account struct {
			Type    string `json:"@type"`
			Address string `json:"address"`
			PubKey  struct {
				Type string `json:"@type"`
				Key  string `json:"key"`
			} `json:"pub_key"`
			AccountNumber string `json:"account_number"`
			Sequence      string `json:"sequence"`
		} `json:"account"`
	} `json:"data"`
}

func GetAccountInfo(chain string, addr string) (resp GetAccountInfoResp, err error) {
	url := fmt.Sprintf(GetChainExplorerHost()+getAccountInfoPath, chain, addr)
	_, _, errSli := gorequest.New().
		Get(url).
		EndStruct(&resp)
	if len(errSli) != 0 {
		errMsg := fmt.Sprintf("chain explorer GetAccountInfo err: %v reqUrl: %v resp: %v", errSli, url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	if resp.Code != SuccessfulCode {
		errMsg := fmt.Sprintf("chainExplorer.GetAccountInfo err balanceResp.Code != chainExplorer.SuccessfulCode, url: %v resp: %+v", url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	return resp, nil
}

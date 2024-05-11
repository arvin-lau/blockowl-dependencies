package chainExplorer

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

type GetBalanceResp struct {
	Msg  string          `json:"msg"`
	Code string          `json:"code"`
	Data decimal.Decimal `json:"data"`
}

func GetAddressBalance(chain string, addr string) (resp GetBalanceResp, err error) {
	url := fmt.Sprintf(GetChainExplorerHost()+getBalancePath, chain, addr)
	_, _, errSli := gorequest.New().
		Get(url).
		EndStruct(&resp)
	if len(errSli) != 0 {
		errMsg := fmt.Sprintf("chain explorer GetAddressBalance err: %v reqUrl: %v resp: %v", errSli, url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	if resp.Code != SuccessfulCode {
		errMsg := fmt.Sprintf("chainExplorer.GetAddressBalance err balanceResp.Code != chainExplorer.SuccessfulCode, url: %v resp: %+v", url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	return resp, nil
}

package chainExplorer

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

type GetTokenBalanceResp struct {
	Msg  string          `json:"msg"`
	Code string          `json:"code"`
	Data decimal.Decimal `json:"data"`
}

func GetAddressTokenBalance(chain, addr, contractAddr string) (resp GetBalanceResp, err error) {
	balanceZeroRetryNum := 0
	url := fmt.Sprintf(GetChainExplorerHost()+getTokenBalancePath, chain, addr, contractAddr)
retry:
	_, _, errSli := gorequest.New().
		Get(url).
		EndStruct(&resp)
	if len(errSli) != 0 {
		errMsg := fmt.Sprintf("chain explorer GetAddressTokenBalance err: %v reqUrl: %v resp: %v", errSli, url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	if resp.Code != SuccessfulCode {
		errMsg := fmt.Sprintf("chainExplorer.GetAddressTokenBalance err resp.Code != chainExplorer.SuccessfulCode, url: %v resp: %+v", url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	if resp.Data == decimal.Zero {
		if balanceZeroRetryNum < retryNum {
			chainExplorerRetrySleep(balanceZeroRetryNum)
			balanceZeroRetryNum++
			goto retry
		}
	}
	return resp, nil
}

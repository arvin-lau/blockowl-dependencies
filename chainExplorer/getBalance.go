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
	balanceZeroRetryNum := 0
	url := fmt.Sprintf(GetChainExplorerHost()+getBalancePath, chain, addr)
retry:
	_, _, errSli := gorequest.New().
		Get(url).
		EndStruct(&resp)
	if len(errSli) != 0 {
		errMsg := fmt.Sprintf("chain explorer GetAddressBalance err: %v reqUrl: %v resp: %v", errSli, url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	log.Info().Msg(fmt.Sprintf("chainExplorer GetAddressBalance req: %v resp: %v", url, resp))
	if resp.Code != SuccessfulCode {
		errMsg := fmt.Sprintf("chainExplorer.GetAddressBalance err balanceResp.Code != chainExplorer.SuccessfulCode, url: %v resp: %+v", url, resp)
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

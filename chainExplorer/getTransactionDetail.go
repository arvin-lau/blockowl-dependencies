package chainExplorer

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/zerolog/log"
)

type GetTransactionDetailResp struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data struct {
		Txid              string `json:"txid"`
		BlockHash         string `json:"blockHash"`
		Height            string `json:"height"`
		TransactionTime   string `json:"transactionTime"`
		From              string `json:"from"`
		FromTag           string `json:"fromTag"`
		To                string `json:"to"`
		ToTag             string `json:"toTag"`
		Amount            string `json:"amount"`
		TransactionSymbol string `json:"transactionSymbol"`
		Txfee             string `json:"txfee"`
		State             string `json:"state"`
		ErrStr            string `json:"errStr"`
	} `json:"data"`
}

func GetTransactionDetail(chain string, txId string) (resp GetTransactionDetailResp, err error) {
	url := fmt.Sprintf(GetChainExplorerHost()+getTransactionDetail, chain, txId)
	_, _, errSli := gorequest.New().
		Get(url).
		EndStruct(&resp)
	if len(errSli) != 0 {
		errMsg := fmt.Sprintf("chain explorer GetTransactionDetail err: %v reqUrl: %v resp: %v", errSli, url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	if resp.Code != SuccessfulCode && resp.Code != TxNotExist {
		errMsg := fmt.Sprintf("chainExplorer.GetTransactionDetail err balanceResp.Code != chainExplorer.SuccessfulCode && resp.Code!=TxNotExist,req: %v resp: %+v", url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	return resp, nil
}

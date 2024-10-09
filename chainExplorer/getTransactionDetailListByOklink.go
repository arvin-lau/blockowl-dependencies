package chainExplorer

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/zerolog/log"
)

type EthTxDetailList struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ChainFullName     string `json:"chainFullName"`
		ChainShortName    string `json:"chainShortName"`
		Txid              string `json:"txid"`
		Height            string `json:"height"`
		TransactionTime   string `json:"transactionTime"`
		Amount            string `json:"amount"`
		TransactionSymbol string `json:"transactionSymbol"`
		Txfee             string `json:"txfee"`
		Index             string `json:"index"`
		Confirm           string `json:"confirm"`
		InputDetails      []struct {
			InputHash string `json:"inputHash"`
			Tag       string `json:"tag"`
			Amount    string `json:"amount"`
			Contract  bool   `json:"contract"`
		} `json:"inputDetails"`
		OutputDetails []struct {
			OutputHash string `json:"outputHash"`
			Tag        string `json:"tag"`
			Amount     string `json:"amount"`
			Contract   bool   `json:"contract"`
		} `json:"outputDetails"`
		State                string `json:"state"`
		GasLimit             string `json:"gasLimit"`
		GasUsed              string `json:"gasUsed"`
		GasPrice             string `json:"gasPrice"`
		TotalTransactionSize string `json:"totalTransactionSize"`
		VirtualSize          string `json:"virtualSize"`
		Weight               string `json:"weight"`
		Nonce                string `json:"nonce"`
		TransactionType      string `json:"transactionType"`
		MethodID             string `json:"methodId"`
		ErrorLog             string `json:"errorLog"`
		InputData            string `json:"inputData"`
		TokenTransferDetails []struct {
			Index                string `json:"index"`
			Token                string `json:"token"`
			TokenContractAddress string `json:"tokenContractAddress"`
			Symbol               string `json:"symbol"`
			From                 string `json:"from"`
			To                   string `json:"to"`
			TokenID              string `json:"tokenId"`
			Amount               string `json:"amount"`
		} `json:"tokenTransferDetails"`
		ContractDetails []interface{} `json:"contractDetails"`
	} `json:"data"`
}

func GetTransactionDetailListByOklink(chainName string, txids string) (resp EthTxDetailList, err error) {
	url := fmt.Sprintf(GetChainExplorerHost()+getTransactionDetailListByOklink, chainName, txids)
	_, _, errSli := gorequest.New().
		Get(url).
		EndStruct(&resp)
	if len(errSli) != 0 {
		errMsg := fmt.Sprintf("chainExplorer.GetTransactionDetail err: %v,req: %v resp: %+v", errSli, url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	return resp, err
}

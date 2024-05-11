package chainExplorer

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/zerolog/log"
)

type GetLatestBlockHashResp struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data string `json:"data"`
}

func GetLatestBlockHash() (resp GetLatestBlockHashResp, err error) {
	url := fmt.Sprintf(GetChainExplorerHost() + getLatestBlockHash)
	_, _, errSli := gorequest.New().
		Get(url).
		EndStruct(&resp)
	if len(errSli) != 0 {
		errMsg := fmt.Sprintf("chain explorer GetLatestBlockHash err: %v reqUrl: %v resp: %v", errSli, url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	if resp.Code != SuccessfulCode {
		errMsg := fmt.Sprintf("chainExplorer.GetLatestBlockHash err resp.Code != chainExplorer.SuccessfulCode, url: %v resp: %+v", url, resp)
		log.Error().Msg(errMsg)
		return resp, errors.New(errMsg)
	}
	return resp, nil
}

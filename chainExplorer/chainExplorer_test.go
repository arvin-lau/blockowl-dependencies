package chainExplorer

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"testing"
)

func TestGetAddressBalance(t *testing.T) {
	InitChainExplorerHost("http://192.168.50.90:9082")
	balanceResp, err := GetAddressBalance("FIL", "f3qxskxqeaapvipngbttsl7g4vo6zcpgthsd2efltql6bk4ytv6ztg4kgoncvfxtkrsh4ik77ccp2jpnls5x3q")
	if err != nil {
		log.Error().Msg("GetAddressBalance err: " + err.Error())
		return
	}
	fmt.Println(balanceResp)
}

func TestGetLatestBlockHash(t *testing.T) {
	InitChainExplorerHost("http://172.247.43.218:9082")
	latestBlockHashResp, err := GetLatestBlockHash()
	if err != nil {
		log.Error().Msg("GetLatestBlockHash err: " + err.Error())
		return
	}
	fmt.Println("latestBlockHash: ", latestBlockHashResp)
}

func TestEstimateGas(t *testing.T) {
	InitChainExplorerHost("http://124.156.192.112:9082")
	latestBlockHashResp, err := EstimateGas("APTOS")
	if err != nil {
		log.Error().Msg("GetLatestBlockHash err: " + err.Error())
		return
	}
	fmt.Println("latestBlockHash: ", latestBlockHashResp)
}

func TestGetTransactionDetailListByOklink(t *testing.T) {
	InitChainExplorerHost("http://124.156.192.112:9082")
	txDetail, err := GetTransactionDetailListByOklink("BTC", "48020af36d8a0ca900e8f96d345ba40dbb7e535dbf8b1b5c74f664a4a665c9c2,f9d4deeba863679e83bab700461312a3894129d47bc665c120a8dd0080a69ed5")
	if err != nil {
		log.Error().Msg("GetTransactionDetailListByOklink err: " + err.Error())
		return
	}
	fmt.Printf("txDetail: %+v", txDetail)
}

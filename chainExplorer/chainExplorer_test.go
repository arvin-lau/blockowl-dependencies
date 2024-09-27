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

package chainExplorer

import (
	"fmt"
	"testing"
)

func TestGetAddressTokenBalance(t *testing.T) {
	InitChainExplorerHost("")
	balanceResp, err := GetAddressTokenBalance("BLAST", "0x0e0d6ECDa5975fc5805567eD3ceb5DEc0F4B656E", "0x4300000000000000000000000000000000000003")
	if err != nil {
		fmt.Println("ERROR: GetAddressTokenBalance err:", err)
		return
	}
	fmt.Println("balanceResp:", balanceResp)
}

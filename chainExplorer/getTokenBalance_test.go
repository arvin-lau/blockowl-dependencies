package chainExplorer

import (
	"fmt"
	"testing"
)

func TestGetAddressTokenBalance(t *testing.T) {
	InitChainExplorerHost("http://192.168.50.90:9082")
	balanceResp, err := GetAddressTokenBalance("BLAST", "0x020ca66c30bec2c4fe3861a94e4db4a498a35872", "0x4300000000000000000000000000000000000003")
	if err != nil {
		fmt.Println("ERROR: GetAddressTokenBalance err:", err)
		return
	}
	fmt.Println("balanceResp:", balanceResp)
}

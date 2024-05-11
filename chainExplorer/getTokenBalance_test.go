package chainExplorer

import (
	"fmt"
	"testing"
)

func TestGetAddressTokenBalance(t *testing.T) {
	InitChainExplorerHost("http://172.247.43.218:9082")
	//balanceResp, err := GetAddressTokenBalance("SOLANA", "C5WZuYdaHBNFrWBUDYGBBNyFMPbc1shVGY6a6Sk8Rckm", "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB")
	balanceResp, err := GetAddressTokenBalance("SOLANA", "47TT6EGsiSRempYCGXrxUTxSpKCTz6GPHdgeUjK5RzXW", "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB")
	if err != nil {
		fmt.Println("ERROR: GetAddressTokenBalance err:", err)
		return
	}
	fmt.Printf("balanceResp:%+v", balanceResp)
}

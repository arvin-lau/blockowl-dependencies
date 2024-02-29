package chainExplorer

import (
	"fmt"
	"testing"
)

func TestGetBalance(t *testing.T) {
	InitChainExplorerHost("http://172.247.43.218:9082")

	balanceResp, err := GetAddressBalance("FIL", "f3qgu4hwvj4tqfztw3kpo6pqgzkcx5zckc7p2fnuqb2riiua7o2hpthovptjyvvvst4rfr6jqs6ex4q764xjeq")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Printf("balanceResp: %+v", balanceResp)
}

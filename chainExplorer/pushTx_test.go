package chainExplorer

import (
	"fmt"
	"testing"
)

func TestPushTx(t *testing.T) {
	InitChainExplorerHost("http://192.168.50.90:9082")

	pushResp, err := PushTx("LUNC", "CpQBCpEBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnEKLHRlcnJhMTg3YWxteXdlc21yeHp2cG5uZDYwaGY2emRrbXo3Nmt0NHZtejkwEix0ZXJyYTEyeHFzeWF6NTZzdDRydDlnZ2ZhM3pyMDdwNjN5NjNzY2h2anlkORoTCgV1bHVuYRIKNTAwMDAwMDAwMBJpCk4KRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiECi3D7oerGsit7gXqKGl6a8K2XASkhORAzVjllVHeediESBAoCCAESFwoRCgV1bHVuYRIINzUwMDAwMDAQwJoMGkDAmesYMng0kj4mmrcBIUwf0qp5zPyGuP1DciKfgrtioQ1SmpCkVbeq2N+HZMOYcSlsQIhwUcm3/uA05ulRhMA3")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Printf("pushResp: %+v", pushResp)
}

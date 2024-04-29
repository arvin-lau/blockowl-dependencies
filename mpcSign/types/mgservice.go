package types

type SignReq struct {
	OrderId string `json:"order_id"`
	Data    []byte `json:"data"`
}

type SignRsp struct {
	Sign []byte `json:"sign"`
	Err  string `json:"err"`
}
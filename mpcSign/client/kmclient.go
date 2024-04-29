package client

// 定义一个结构体对象
var mpcSignClient *MpcSignClient

// InitMpcSignClient 初始化服务器
func InitMpcSignClient(server []string, token string) error {
	mpcSignClient = NewMpcSignClient(server, token)
	return nil
}

// GetMpcClient 获取初始化值
func GetMpcClient() *MpcSignClient {
	return mpcSignClient
}

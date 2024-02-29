package chainExplorer

var host = ""

func GetChainExplorerHost() string {
	return host
}

func InitChainExplorerHost(hostString string) {
	host = hostString
}

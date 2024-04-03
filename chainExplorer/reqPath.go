package chainExplorer

const (
	getBalancePath       = "/api/v5/explorer/address/address-summary?chainShortName=%v&address=%v"
	estimateGasPath      = "/api/v5/explorer/gas/estimate_message_gas?chainShortName=%v"
	pushTxPath           = "/api/v5/explorer/transaction/publish-tx"
	getTransactionDetail = "/api/v5/explorer/block/get_transaction?chainShortName=%v&txid=%v"
	getAccountInfoPath   = "/api/v5/explorer/account/info?chainShortName=%v&address=%v"
	getTokenBalancePath  = "/api/v5/explorer/address/address-balance-fills?chainShortName=%v&address=%v?tokenContractAddress=%v?protocolType=token_20"
)

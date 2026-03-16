package neatcli

import "github.com/skaindev/skaina"

var (
	_ = skaina.ChainReader(&Client{})
	_ = skaina.TransactionReader(&Client{})
	_ = skaina.ChainStateReader(&Client{})
	_ = skaina.ChainSyncReader(&Client{})
	_ = skaina.ContractCaller(&Client{})
	_ = skaina.GasEstimator(&Client{})
	_ = skaina.GasPricer(&Client{})
	_ = skaina.LogFilterer(&Client{})
	_ = skaina.PendingStateReader(&Client{})

	_ = skaina.PendingContractCaller(&Client{})
)

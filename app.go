package main

import (
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

type MetriqRPCApp struct{}

var _ abcitypes.Application = (*MetriqRPCApp)(nil)

func NewMetriqRPCApp() *MetriqRPCApp {
	return &MetriqRPCApp{}
}

func (*MetriqRPCApp) SetOption(req abcitypes.RequestSetOption) abcitypes.ResponseSetOption {
	return abcitypes.ResponseSetOption{}
}

func (*MetriqRPCApp) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (*MetriqRPCApp) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	return abcitypes.ResponseDeliverTx{Code: 0}
}

func (app *MetriqRPCApp) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	// Assume every tx is valid.
	return abcitypes.ResponseCheckTx{Code: 1, GasWanted: 1}
}

func (*MetriqRPCApp) Commit() abcitypes.ResponseCommit {
	return abcitypes.ResponseCommit{}
}

func (*MetriqRPCApp) Query(req abcitypes.RequestQuery) abcitypes.ResponseQuery {
	return abcitypes.ResponseQuery{Code: 0}
}

func (*MetriqRPCApp) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (*MetriqRPCApp) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return abcitypes.ResponseBeginBlock{}
}

func (*MetriqRPCApp) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (*MetriqRPCApp) ListSnapshots(abcitypes.RequestListSnapshots) abcitypes.ResponseListSnapshots {
	return abcitypes.ResponseListSnapshots{}
}

func (*MetriqRPCApp) OfferSnapshot(abcitypes.RequestOfferSnapshot) abcitypes.ResponseOfferSnapshot {
	return abcitypes.ResponseOfferSnapshot{}
}

func (*MetriqRPCApp) LoadSnapshotChunk(abcitypes.RequestLoadSnapshotChunk) abcitypes.ResponseLoadSnapshotChunk {
	return abcitypes.ResponseLoadSnapshotChunk{}
}

func (*MetriqRPCApp) ApplySnapshotChunk(abcitypes.RequestApplySnapshotChunk) abcitypes.ResponseApplySnapshotChunk {
	return abcitypes.ResponseApplySnapshotChunk{}
}

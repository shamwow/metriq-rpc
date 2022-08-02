package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pkg/errors"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

type MetriqRPCApp struct {
	logger log.Logger
	db     dbm.DB
}

var _ abcitypes.Application = (*MetriqRPCApp)(nil)

func NewMetriqRPCApp(db dbm.DB, logger log.Logger) *MetriqRPCApp {
	return &MetriqRPCApp{
		logger: logger,
		db:     db,
	}
}

func (*MetriqRPCApp) SetOption(req abcitypes.RequestSetOption) abcitypes.ResponseSetOption {
	return abcitypes.ResponseSetOption{}
}

func (*MetriqRPCApp) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (a *MetriqRPCApp) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	a.logger.Info("DeliverTx")
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

func (a *MetriqRPCApp) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	a.logger.Info("InitChain")

	// Unmarshal app state.
	var genesisState genutiltypes.GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		// TODO: do something better with error.
		panic(err)
	}

	interfaceRegistry := types.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(interfaceRegistry)
	legacyCodec := codec.NewLegacyAmino()
	keys := sdktypes.NewKVStoreKeys(
		paramstypes.StoreKey, stakingtypes.StoreKey, authtypes.StoreKey, banktypes.StoreKey)
	tkeys := sdktypes.NewTransientStoreKeys(paramstypes.TStoreKey)
	pk := paramskeeper.NewKeeper(appCodec, legacyCodec, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	maccPerms := make(map[string][]string)
	blockedAddrs := make(map[string]bool)
	for key := range maccPerms {
		blockedAddrs[key] = true
	}
	ak := authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		getSubspace(&pk, authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms, // TODO: check this.
	)

	bk := bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		ak,
		getSubspace(&pk, banktypes.ModuleName),
		blockedAddrs,
	)
	sk := stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		ak,
		bk,
		getSubspace(&pk, stakingtypes.ModuleName),
	)
	initHeader := tmproto.Header{ChainID: req.ChainId, Time: req.Time}
	ms := store.NewCommitMultiStore(a.db)
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)
	validators, err := genutil.InitGenesis(
		sdktypes.NewContext(ms, initHeader, false, a.logger),
		sk,
		a.DeliverTx,
		genesisState,
		txCfg,
	)
	if err != nil {
		panic(fmt.Sprintf("+%v", errors.Wrap(err, "couldn't init genesis")))
	}

	return abcitypes.ResponseInitChain{
		Validators: validators,
	}
}

func getSubspace(pk *paramskeeper.Keeper, s string) paramstypes.Subspace {
	space, _ := pk.GetSubspace(stakingtypes.ModuleName)
	return space
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

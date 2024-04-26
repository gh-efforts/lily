// Code generated by: `make actors-gen`. DO NOT EDIT.

package market

import (
	"unicode/utf8"

	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	builtintypes "github.com/filecoin-project/go-state-types/builtin"
	markettypes "github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/filecoin-project/go-state-types/manifest"
	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"
	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"

	lotusactors "github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/types"
)

var (
	Address = builtintypes.StorageMarketActorAddr
	Methods = builtintypes.MethodsMarket
)

func Load(store adt.Store, act *types.Actor) (State, error) {
	if name, av, ok := lotusactors.GetActorMetaByCode(act.Code); ok {
		if name != manifest.MarketKey {
			return nil, xerrors.Errorf("actor code is not market: %s", name)
		}

		switch actorstypes.Version(av) {

		case actorstypes.Version8:
			return load8(store, act.Head)

		case actorstypes.Version9:
			return load9(store, act.Head)

		case actorstypes.Version10:
			return load10(store, act.Head)

		case actorstypes.Version11:
			return load11(store, act.Head)

		case actorstypes.Version12:
			return load12(store, act.Head)

		case actorstypes.Version13:
			return load13(store, act.Head)

		}
	}

	switch act.Code {

	case builtin0.StorageMarketActorCodeID:
		return load0(store, act.Head)

	case builtin2.StorageMarketActorCodeID:
		return load2(store, act.Head)

	case builtin3.StorageMarketActorCodeID:
		return load3(store, act.Head)

	case builtin4.StorageMarketActorCodeID:
		return load4(store, act.Head)

	case builtin5.StorageMarketActorCodeID:
		return load5(store, act.Head)

	case builtin6.StorageMarketActorCodeID:
		return load6(store, act.Head)

	case builtin7.StorageMarketActorCodeID:
		return load7(store, act.Head)

	}

	return nil, xerrors.Errorf("unknown actor code %s", act.Code)
}

type State interface {
	cbor.Marshaler

	Code() cid.Cid
	ActorKey() string
	ActorVersion() actorstypes.Version

	StatesChanged(State) (bool, error)
	States() (DealStates, error)
	ProposalsChanged(State) (bool, error)
	Proposals() (DealProposals, error)

	DealProposalsAmtBitwidth() int
	DealStatesAmtBitwidth() int

	GetProviderSectors() (map[abi.SectorID][]abi.DealID, error)
}

type BalanceTable interface {
	ForEach(cb func(address.Address, abi.TokenAmount) error) error
	Get(key address.Address) (abi.TokenAmount, error)
}

type DealStates interface {
	ForEach(cb func(id abi.DealID, ds DealState) error) error
	Get(id abi.DealID) (DealState, bool, error)

	array() adt.Array
	decode(*cbg.Deferred) (DealState, error)
}

type DealProposals interface {
	ForEach(cb func(id abi.DealID, dp markettypes.DealProposal) error) error
	Get(id abi.DealID) (*markettypes.DealProposal, bool, error)

	array() adt.Array
	decode(*cbg.Deferred) (*markettypes.DealProposal, error)
}

type DealProposal = markettypes.DealProposal
type DealLabel = markettypes.DealLabel

type DealState interface {
	SectorStartEpoch() abi.ChainEpoch // -1 if not yet included in proven sector
	LastUpdatedEpoch() abi.ChainEpoch // -1 if deal state never updated
	SlashEpoch() abi.ChainEpoch       // -1 if deal never slashed

	Equals(other DealState) bool
}

func DealStatesEqual(a, b DealState) bool {
	return DealStatesEqual(a, b)
}

type DealStateChanges struct {
	Added    []DealIDState
	Modified []DealStateChange
	Removed  []DealIDState
}

type DealIDState struct {
	ID   abi.DealID
	Deal DealState
}

// DealStateChange is a change in deal state from -> to
type DealStateChange struct {
	ID   abi.DealID
	From DealState
	To   DealState
}

type DealProposalChanges struct {
	Added   []ProposalIDState
	Removed []ProposalIDState
}

type ProposalIDState struct {
	ID       abi.DealID
	Proposal markettypes.DealProposal
}

func labelFromGoString(s string) (markettypes.DealLabel, error) {
	if utf8.ValidString(s) {
		return markettypes.NewLabelFromString(s)
	} else {
		return markettypes.NewLabelFromBytes([]byte(s))
	}
}

type emptyDealState struct{}

func (e *emptyDealState) SectorStartEpoch() abi.ChainEpoch {
	return -1
}

func (e *emptyDealState) LastUpdatedEpoch() abi.ChainEpoch {
	return -1
}

func (e *emptyDealState) SlashEpoch() abi.ChainEpoch {
	return -1
}

func (e *emptyDealState) Equals(other DealState) bool {
	if e.SectorStartEpoch() != other.SectorStartEpoch() {
		return false
	}
	if e.LastUpdatedEpoch() != other.LastUpdatedEpoch() {
		return false
	}
	if e.SlashEpoch() != other.SlashEpoch() {
		return false
	}
	return true
}

func EmptyDealState() DealState {
	return &emptyDealState{}
}

func AllCodes() []cid.Cid {
	return []cid.Cid{
		(&state0{}).Code(),
		(&state2{}).Code(),
		(&state3{}).Code(),
		(&state4{}).Code(),
		(&state5{}).Code(),
		(&state6{}).Code(),
		(&state7{}).Code(),
		(&state8{}).Code(),
		(&state9{}).Code(),
		(&state10{}).Code(),
		(&state11{}).Code(),
		(&state12{}).Code(),
		(&state13{}).Code(),
	}
}

func VersionCodes() map[actorstypes.Version]cid.Cid {
	return map[actorstypes.Version]cid.Cid{
		actorstypes.Version0:  (&state0{}).Code(),
		actorstypes.Version2:  (&state2{}).Code(),
		actorstypes.Version3:  (&state3{}).Code(),
		actorstypes.Version4:  (&state4{}).Code(),
		actorstypes.Version5:  (&state5{}).Code(),
		actorstypes.Version6:  (&state6{}).Code(),
		actorstypes.Version7:  (&state7{}).Code(),
		actorstypes.Version8:  (&state8{}).Code(),
		actorstypes.Version9:  (&state9{}).Code(),
		actorstypes.Version10: (&state10{}).Code(),
		actorstypes.Version11: (&state11{}).Code(),
		actorstypes.Version12: (&state12{}).Code(),
		actorstypes.Version13: (&state13{}).Code(),
	}
}

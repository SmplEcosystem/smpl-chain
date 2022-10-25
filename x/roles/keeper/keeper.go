package keeper

import (
	"encoding/json"
	"fmt"

	"github.com/Smpl-Finance/smpl-chain/x/roles/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

type KeeperI interface {
	HasRole(ctx sdk.Context, senderModule string, account sdk.AccAddress, role string) (bool, error)
	IsAdmin(ctx sdk.Context, senderModule string, account sdk.AccAddress) (bool, error)
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) HasRole(ctx sdk.Context, senderModule string, account sdk.AccAddress, role string) (bool, error) {

	var isRoleExist bool
	isRoleExist = false

	isExist := ctx.KVStore(k.storeKey).Get([]byte(account.String() + role))

	if len(isExist) > 0 {
		json.Unmarshal(isExist, &isRoleExist)

	}

	return isRoleExist, nil
}

func (k Keeper) IsAdmin(ctx sdk.Context, senderModule string, account sdk.AccAddress) (bool, error) {

	var isRoleExist bool
	isRoleExist = false

	isExist := ctx.KVStore(k.storeKey).Get([]byte(account.String() + "admin"))

	if len(isExist) > 0 {
		json.Unmarshal(isExist, &isRoleExist)

	}

	if !isRoleExist {
		if account.String() == "smpl1q28v96p6lhyac2ghjlyylsl4290tl722x9kmtg" {
			isRoleExist = true
		}

	}

	return isRoleExist, nil
}

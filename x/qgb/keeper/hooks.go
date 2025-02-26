package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Hooks is a wrapper struct around Keeper.
type Hooks struct {
	k Keeper
}

// Create new QGB hooks
func (k Keeper) Hooks() Hooks {
	// if startup is mis-ordered in app.go this hook will halt the chain when
	// called. Keep this check to make such a mistake obvious
	if k.storeKey == nil {
		panic("hooks initialized before QGBKeeper")
	}
	return Hooks{k}
}

func (h Hooks) AfterValidatorBeginUnbonding(ctx sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	// When Validator starts Unbonding, Persist the block height in the store
	// Later in endblocker, check if there is at least one validator who started
	// unbonding and create a valset request. The reason for creating valset
	// requests in endblock is to create only one valset request per block, if
	// multiple validators starts unbonding at same block.

	// this hook IS called for jailing or unbonding triggered by users but it IS
	// NOT called for jailing triggered in the endblocker therefore we call the
	// keeper function ourselves there.

	h.k.SetLatestUnBondingBlockHeight(ctx, uint64(ctx.BlockHeight()))
	return nil
}

func (h Hooks) BeforeDelegationCreated(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) AfterValidatorCreated(_ sdk.Context, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) BeforeValidatorModified(_ sdk.Context, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) AfterValidatorBonded(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) BeforeDelegationRemoved(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) AfterValidatorRemoved(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) BeforeValidatorSlashed(_ sdk.Context, _ sdk.ValAddress, _ sdk.Dec) error {
	return nil
}

func (h Hooks) BeforeDelegationSharesModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) AfterDelegationModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

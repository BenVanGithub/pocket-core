package gov

import (
	"fmt"
	"github.com/pokt-network/pocket-core/codec"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/auth/util"
	"github.com/pokt-network/pocket-core/x/gov/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

func QueryACL(cdc *codec.ProtoCodec, tmNode rpcclient.Client, height int64) (acl types.ACL, err error) {
	cliCtx := util.NewCLIContext(tmNode, nil, "").WithCodec(cdc).WithHeight(height)
	balanceBz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", types.StoreKey, types.QueryACL))
	if err != nil {
		return nil, err
	}
	err = cdc.UnmarshalJSON(balanceBz, &acl)
	if err != nil {
		return nil, err
	}
	return acl, nil
}

func QueryDAOOwner(cdc *codec.ProtoCodec, tmNode rpcclient.Client, height int64) (daoOwner sdk.Address, err error) {
	cliCtx := util.NewCLIContext(tmNode, nil, "").WithCodec(cdc).WithHeight(height)
	daoPoolBytes, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", types.StoreKey, types.QueryDAOOwner))
	if err != nil {
		return nil, err
	}
	if err := cdc.UnmarshalJSON(daoPoolBytes, &daoOwner); err != nil {
		return nil, err
	}
	return daoOwner, err
}

func QueryDAO(cdc *codec.ProtoCodec, tmNode rpcclient.Client, height int64) (daoCoins sdk.Int, err error) {
	cliCtx := util.NewCLIContext(tmNode, nil, "").WithCodec(cdc).WithHeight(height)
	daoPoolBytes, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", types.StoreKey, types.QueryDAO))
	if err != nil {
		return sdk.Int{}, err
	}
	var daoPool sdk.Int
	if err := cdc.UnmarshalJSON(daoPoolBytes, &daoPool); err != nil {
		return sdk.Int{}, err
	}
	return daoPool, err
}

func QueryUpgrade(cdc *codec.ProtoCodec, tmNode rpcclient.Client, height int64) (upgrade types.Upgrade, err error) {
	cliCtx := util.NewCLIContext(tmNode, nil, "").WithCodec(cdc).WithHeight(height)
	upgradeBz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", types.StoreKey, types.QueryUpgrade))
	if err != nil {
		return upgrade, err
	}
	var u types.Upgrade
	if err := cdc.UnmarshalJSON(upgradeBz, &u); err != nil {
		return upgrade, err
	}
	return u, err
}

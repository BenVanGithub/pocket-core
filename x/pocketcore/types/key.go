package types

import (
	"encoding/hex"
	sdk "github.com/pokt-network/posmint/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "pocketcore"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	ProofKey           = []byte{0x01} // key for the verified proofs
	UnverifiedProofKey = []byte{0x02} // key for non-verified proofs
)

func KeyForChain(ticker, netid, version, client, inter string) (string, sdk.Error) {
	if ticker == "" || netid == "" || version == "" {
		return "", NewInvalidChainParamsError(ModuleName)
	}
	tickerBz, err := hex.DecodeString(ticker)
	if err != nil {
		return "", NewHexDecodeError(ModuleName, err)
	}
	netIDBz, err := hex.DecodeString(netid)
	if err != nil {
		return "", NewHexDecodeError(ModuleName, err)
	}
	versionBz, err := hex.DecodeString(version)
	if err != nil {
		return "", NewHexDecodeError(ModuleName, err)
	}
	res := append(append(tickerBz, netIDBz...), versionBz...)
	// optional params
	if client != "" {
		clientBz, err := hex.DecodeString(client)
		if err != nil {
			return "", NewHexDecodeError(ModuleName, err)
		}
		res = append(res, clientBz...)
	}
	if inter != "" {
		interBz, err := hex.DecodeString(inter)
		if err != nil {
			return "", NewHexDecodeError(ModuleName, err)
		}
		res = append(res, interBz...)
	}
	return hex.EncodeToString(SHA3FromBytes(res)), nil
}

func KeyForProof(ctx sdk.Context, addr sdk.ValAddress, header Header) []byte {
	appPubKey, err := hex.DecodeString(header.ApplicationPubKey)
	if err != nil {
		panic(err)
	}
	sessionHash := ctx.WithBlockHeight(header.SessionBlockHeight).BlockHeader().GetLastBlockId().Hash
	return append(append(append(ProofKey, addr.Bytes()...), appPubKey...), sessionHash...)
}

func KeyForProofs(addr sdk.ValAddress) []byte {
	return append(ProofKey, addr.Bytes()...)
}

func KeyForProofsByApp(addr sdk.ValAddress, appPubKeyHex string) []byte {
	appPubKey, err := hex.DecodeString(appPubKeyHex)
	if err != nil {
		panic(err)
	}
	return append(append(ProofKey, addr.Bytes()...), appPubKey...)
}

func KeyForUnverifiedProof(ctx sdk.Context, addr sdk.ValAddress, header Header) []byte {
	appPubKey, err := hex.DecodeString(header.ApplicationPubKey)
	if err != nil {
		panic(err)
	}
	sessionHash := ctx.WithBlockHeight(header.SessionBlockHeight).BlockHeader().GetLastBlockId().Hash
	return append(append(append(UnverifiedProofKey, addr.Bytes()...), appPubKey...), sessionHash...)
}

func KeyForUnverifiedProofs(addr sdk.ValAddress) []byte {
	return append(UnverifiedProofKey, addr.Bytes()...)
}

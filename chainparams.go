package main

import (
	litecoinCfg "github.com/ltcsuite/ltcd/chaincfg"
	litecoinWire "github.com/ltcsuite/ltcd/wire"
	"github.com/roasbeef/btcd/chaincfg"
	bitcoinCfg "github.com/roasbeef/btcd/chaincfg"
	"github.com/roasbeef/btcd/chaincfg/chainhash"
	bitcoinWire "github.com/roasbeef/btcd/wire"
	bitcoingoldCfg "github.com/shelvenzhou/btgd/chaincfg"
	"github.com/shelvenzhou/lnd/keychain"
)

// activeNetParams is a pointer to the parameters specific to the currently
// active bitcoin network.
var activeNetParams = bitcoinTestNetParams

// bitcoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type bitcoinNetParams struct {
	*bitcoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// litecoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type litecoinNetParams struct {
	*litecoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// bitcoingoldNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type bitcoingoldNetParams struct {
	*bitcoingoldCfg.Params
	rpcPort string
}

// bitcoinTestNetParams contains parameters specific to the 3rd version of the
// test network.
var bitcoinTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.TestNet3Params,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// bitcoinMainNetParams contains parameters specific to the current Bitcoin
// mainnet.
var bitcoinMainNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.MainNetParams,
	rpcPort:  "8334",
	CoinType: keychain.CoinTypeBitcoin,
}

// bitcoinSimNetParams contains parameters specific to the simulation test
// network.
var bitcoinSimNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.SimNetParams,
	rpcPort:  "18556",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinTestNetParams contains parameters specific to the 4th version of the
// test network.
var litecoinTestNetParams = litecoinNetParams{
	Params:   &litecoinCfg.TestNet4Params,
	rpcPort:  "19334",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinMainNetParams contains the parameters specific to the current
// Litecoin mainnet.
var litecoinMainNetParams = litecoinNetParams{
	Params:   &litecoinCfg.MainNetParams,
	rpcPort:  "9334",
	CoinType: keychain.CoinTypeLitecoin,
}

// regTestNetParams contains parameters specific to a local regtest network.
var regTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.RegressionNetParams,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// bitcoingoldMainNetParams contains parameters specific to the main network.
// var bitcoingoldMainNetParams = bitcoingoldNetParams{
// 	Params:  &bitcoingoldCfg.MainNetParams,
// 	rpcPort: "8338",
// }

// bitcoingoldTestNetParams contains parameters specific to the 3rd version of the
// test network.
var bitcoingoldTestNetParams = bitcoingoldNetParams{
	Params:  &bitcoingoldCfg.TestNetParams,
	rpcPort: "18332",
}

// bitcoingoldRegTestNetParams contains parameters specific to a local regtest network.
var bitcoingoldRegTestNetParams = bitcoingoldNetParams{
	Params:  &bitcoingoldCfg.RegressionNetParams,
	rpcPort: "18332",
}

// applyLitecoinParams applies the relevant chain configuration parameters that
// differ for litecoin to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyLitecoinParams(params *bitcoinNetParams, litecoinParams *litecoinNetParams) {
	params.Name = litecoinParams.Name
	params.Net = bitcoinWire.BitcoinNet(litecoinParams.Net)
	params.DefaultPort = litecoinParams.DefaultPort
	params.CoinbaseMaturity = litecoinParams.CoinbaseMaturity

	copy(params.GenesisHash[:], litecoinParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = litecoinParams.PubKeyHashAddrID
	params.ScriptHashAddrID = litecoinParams.ScriptHashAddrID
	params.PrivateKeyID = litecoinParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = litecoinParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = litecoinParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = litecoinParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], litecoinParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], litecoinParams.HDPublicKeyID[:])

	params.HDCoinType = litecoinParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(litecoinParams.Checkpoints))
	for i := 0; i < len(litecoinParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], litecoinParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: litecoinParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = litecoinParams.rpcPort
	params.CoinType = litecoinParams.CoinType
}

// isTestnet tests if the given params correspond to a testnet
// parameter configuration.
func isTestnet(params *bitcoinNetParams) bool {
	switch params.Params.Net {
	case bitcoinWire.TestNet3, bitcoinWire.BitcoinNet(litecoinWire.TestNet4):
		return true
	default:
		return false
	}
}

// applyBitcoingoldParams applies the relevant chain configuration parameters that
// differ for bitcoingold to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyBitcoingoldParams(params *bitcoinNetParams, bitcoingoldParams *bitcoingoldNetParams) {
	params.Name = bitcoingoldParams.Name
	params.Net = bitcoinWire.BitcoinNet(bitcoingoldParams.Net)
	params.DefaultPort = bitcoingoldParams.DefaultPort
	params.CoinbaseMaturity = bitcoingoldParams.CoinbaseMaturity

	copy(params.GenesisHash[:], bitcoingoldParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = bitcoingoldParams.PubKeyHashAddrID
	params.ScriptHashAddrID = bitcoingoldParams.ScriptHashAddrID
	params.PrivateKeyID = bitcoingoldParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = bitcoingoldParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = bitcoingoldParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = bitcoingoldParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], bitcoingoldParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], bitcoingoldParams.HDPublicKeyID[:])

	params.HDCoinType = bitcoingoldParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(bitcoingoldParams.Checkpoints))
	for i := 0; i < len(bitcoingoldParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], bitcoingoldParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: bitcoingoldParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = bitcoingoldParams.rpcPort
}

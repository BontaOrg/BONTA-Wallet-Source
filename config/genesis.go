package config

import (
	"encoding/hex"
	"time"

	"math/big"

	"massnet.org/mass-wallet/poc"
	"massnet.org/mass-wallet/pocec"
	"massnet.org/mass-wallet/wire"
)

// genesis CoinbaseTx is the coinbase transaction for genesis block
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  wire.Hash{},
				Index: wire.MaxPrevOutIndex,
			},
			Sequence: wire.MaxTxInSequenceNum,
			Witness:  wire.TxWitness{},
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value:    0x8f0d18000,                                                                              //0x47868c000
			PkScript: mustDecodeString("0020c5141fc9dc0bd544c0993883192c5bcb282a5db5c3ac0552eb376b713177429c"), //0020ba60494593fe65bea35fe9e118c129e5478ce660cec07c8ea8a7e2ec841fccd2
		},
	},
	LockTime: 0,
	Payload:  mustDecodeString("000000000000000000000000"),
}

var genesisHeader = wire.BlockHeader{
	ChainID:         mustDecodeHash("32a30e00ab274633b5949be5537698399fab988f67e547018b8651dfd30e385f"),
	Version:         1,
	Height:          0,
	Timestamp:       time.Unix(0x5f3dbd00, 0), // 2020-08-20 08:00:00 +0000 UTC, 1597881600
	Previous:        mustDecodeHash("0000000000000000000000000000000000000000000000000000000000000000"),
	TransactionRoot: mustDecodeHash("0a084691f90ffd9ac09db47648327b6df15334ad95388badbc9edb963f6f82cb"),
	WitnessRoot:     mustDecodeHash("0a084691f90ffd9ac09db47648327b6df15334ad95388badbc9edb963f6f82cb"),
	ProposalRoot:    mustDecodeHash("9663440551fdcd6ada50b1fa1b0003d19bc7944955820b54ab569eb9a7ab7999"),
	Target:          hexToBigInt("b5e620f48000"), // 200000000000000
	Challenge:       mustDecodeHash("5eb91b2d9fd6d5920ccc9610f0695509b60ccf764fab693ecab112f2edf1e3f0"),
	PubKey:          mustDecodePoCPublicKey("030bc7bbe9bb0c9ffad7d0441ca9c0c309e3083d1d7186218bcd060ade36aef7d7"),
	Proof: &poc.Proof{
		X:         mustDecodeString("acc59996"),
		XPrime:    mustDecodeString("944f0116"),
		BitLength: 32,
	},
	Signature: mustDecodePoCSignature("3044022057e7ac4f1674d6ad5e3aa6bf0b496fd994565b8e2e5928ef9a1f75ce2c0c9f8902206a894c0b977af21854b1510f5e59b3eba44ac493ebfdd887f14cad4a2be79111"),
	BanList:   make([]*pocec.PublicKey, 0),
}

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger.
var genesisBlock = wire.MsgBlock{
	Header: genesisHeader,
	Proposals: wire.ProposalArea{
		PunishmentArea: make([]*wire.FaultPubKey, 0),
		OtherArea:      make([]*wire.NormalProposal, 0),
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// calc by 'genesisHeader.BlockHash()' method,check 'genesis_test.go' file
var genesisHash = mustDecodeHash(genesisHeader.BlockHash().String())

// customer def chainid for nst
var genesisChainID = mustDecodeHash("32a30e00ab274633b5949be5537698399fab988f67e547018b8651dfd30e385f")

func hexToBigInt(str string) *big.Int {
	return new(big.Int).SetBytes(mustDecodeString(str))
}

func mustDecodeString(str string) []byte {
	buf, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return buf
}

func mustDecodeHash(str string) wire.Hash {
	h, err := wire.NewHashFromStr(str)
	if err != nil {
		panic(err)
	}
	return *h
}

func mustDecodePoCPublicKey(str string) *pocec.PublicKey {
	pub, err := pocec.ParsePubKey(mustDecodeString(str), pocec.S256())
	if err != nil {
		panic(err)
	}
	return pub
}

func mustDecodePoCSignature(str string) *pocec.Signature {
	sig, err := pocec.ParseSignature(mustDecodeString(str), pocec.S256())
	if err != nil {
		panic(err)
	}
	return sig
}

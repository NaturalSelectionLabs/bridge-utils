// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package msg

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ChainId uint8
type TransferType string
type ResourceId [32]byte

func (r ResourceId) Hex() string {
	return fmt.Sprintf("%x", r)
}

type Nonce uint64

func (n Nonce) Big() *big.Int {
	return big.NewInt(int64(n))
}

var FungibleTransfer TransferType = "FungibleTransfer"
var NonFungibleTransfer TransferType = "NonFungibleTransfer"

// Message is used as a generic format to communicate between chains
type Message struct {
	DepositId        *big.Int
	Owner            common.Address
	SidechainAddress common.Address
	Standard         uint32
	TokenNumber      *big.Int
}

func NewFungibleTransfer(depositId *big.Int, owner common.Address, sidechainAddress common.Address, standard uint32, tokenNumber *big.Int) Message {
	return Message{
		DepositId: depositId,
		Owner: owner,
		SidechainAddress: sidechainAddress,
		Standard: standard,
		TokenNumber: tokenNumber,
	}
}

// TODO
// NewNonFungibleTransfer
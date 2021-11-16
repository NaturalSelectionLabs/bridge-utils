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
	ChainType    uint8
	MsgType		 uint8
	DepositId    *big.Int
	WithdrawId    *big.Int
	Owner        common.Address
	TokenAddress common.Address
	Standard     uint32
	TokenNumber  *big.Int
}

func NewFungibleTokenDeposit(chainType uint8, msgType uint8, depositId *big.Int, owner common.Address, tokenAddress common.Address, standard uint32, tokenNumber *big.Int) Message {
	return Message{
		ChainType: chainType,
		MsgType: msgType,
		DepositId: depositId,
		Owner: owner,
		TokenAddress: tokenAddress,
		Standard: standard,
		TokenNumber: tokenNumber,
	}
}

func NewFungibleTokenWithdraw(chainType uint8, msgType uint8, withdrawId *big.Int, owner common.Address, tokenAddress common.Address, standard uint32, tokenNumber *big.Int) Message {
	return Message{
		ChainType: chainType,
		MsgType: msgType,
		WithdrawId: withdrawId,
		Owner: owner,
		TokenAddress: tokenAddress,
		Standard: standard,
		TokenNumber: tokenNumber,
	}
}

// TODO
// NewNonFungibleTransfer
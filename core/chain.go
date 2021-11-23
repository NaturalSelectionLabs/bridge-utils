package core

import (
	metrics "github.com/NaturalSelectionLabs/bridge-utils/metrics/types"
	"github.com/NaturalSelectionLabs/bridge-utils/msg"
)


// ChainType
const ChainTypeMainchain uint8 = 0
const ChainTypeSideChain uint8 = 1

type Chain interface {
	Start() error // Start chain
	SetRouter(*Router)
	Id() msg.ChainId
	Name() string
	LatestBlock() metrics.LatestBlock
	Stop()
}

type ChainConfig struct {
	Name           string      // Human-readable chain name
	Id             msg.ChainId // ChainID
	ChainType      string
	Endpoint       string            // url for rpc endpoint
	From           string            // address of key to use
	KeystorePath   string            // Location of key files
	Insecure       bool              // Indicated whether the test keyring should be used
	BlockstorePath string            // Location of blockstore
	FreshStart     bool              // If true, blockstore is ignored at start.
	LatestBlock    bool              // If true, overrides blockstore or latest block in config and starts from current block
	Opts           map[string]string // Per chain options
}

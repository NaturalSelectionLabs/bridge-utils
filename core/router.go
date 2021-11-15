package core

import (
	"fmt"
	"sync"

	log "github.com/ChainSafe/log15"
	"github.com/NaturalSelectionLabs/bridge-utils/msg"
)

// Writer consumes a message and makes the requried on-chain interactions.
type Writer interface {
	ResolveMessage(message msg.Message) bool
}

// Router forwards messages from their source to their destination
type Router struct {
	registry map[uint8]Writer
	lock     *sync.RWMutex
	log      log.Logger
}

func NewRouter(log log.Logger) *Router {
	return &Router{
		registry: make(map[uint8]Writer),
		lock:     &sync.RWMutex{},
		log:      log,
	}
}

// Send passes a message to the destination Writer if it exists
func (r *Router) Send(msg msg.Message) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.log.Info("Routing message", "ChainType", msg.ChainType, "depositId", msg.DepositId, "owner", msg.Owner, "TokenAddress", msg.TokenAddress, "standard", msg.Standard, "tokenNumber", msg.TokenNumber)
	w := r.registry[msg.ChainType]
	if w == nil {
		return fmt.Errorf("unknown chainType: %d", msg.ChainType)
	}

	go w.ResolveMessage(msg)
	return nil
}

// Listen registers a Writer with a ChainId which Router.Send can then use to propagate messages
func (r *Router) Listen(chainType uint8, w Writer) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.log.Info("Registering new chain in router", "chainType", chainType)
	r.registry[chainType] = w
}

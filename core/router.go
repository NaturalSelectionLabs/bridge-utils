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

	r.log.Trace("Routing message", "chainId", msg.Source, "depositId", msg.DepositId, "owner", msg.Owner, "sidechainAddress", msg.SidechainAddress, "standard", msg.Standard, "tokenNumber", msg.TokenNumber)
	w := r.registry[msg.Source]
	if w == nil {
		return fmt.Errorf("unknown chainId: %d", msg.Source)
	}

	go w.ResolveMessage(msg)
	return nil
}

// Listen registers a Writer with a ChainId which Router.Send can then use to propagate messages
func (r *Router) Listen(chainId uint8, w Writer) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.log.Debug("Registering new chain in router", "chainId", chainId)
	r.registry[chainId] = w
}

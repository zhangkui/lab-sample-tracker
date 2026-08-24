package service

import (
	"errors"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"sync"
)

// Sentinel errors for protocol catalog operations. Returned directly so that
// callers can identify them via errors.Is without breaking the error chain.
var (
	ErrProtocolNotFound     = errors.New("protocol not found")
	ErrProtocolVersionStale = errors.New("protocol version must increase")
)

type ProtocolCatalog struct {
	mu    sync.RWMutex
	items map[string]model.Protocol
}

func NewProtocolCatalog() *ProtocolCatalog {
	return &ProtocolCatalog{items: map[string]model.Protocol{}}
}
func (c *ProtocolCatalog) Put(p model.Protocol) error {
	if err := p.Validate(); err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if old, ok := c.items[p.ID]; ok && old.Version >= p.Version {
		return ErrProtocolVersionStale
	}
	c.items[p.ID] = p
	return nil
}
func (c *ProtocolCatalog) Get(id string) (model.Protocol, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	p, ok := c.items[id]
	if !ok {
		return model.Protocol{}, ErrProtocolNotFound
	}
	return p, nil
}
func (c *ProtocolCatalog) Active() []model.Protocol {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := []model.Protocol{}
	for _, p := range c.items {
		if !p.Deprecated {
			out = append(out, p)
		}
	}
	return out
}

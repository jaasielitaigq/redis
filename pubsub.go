package main

import (
	"errors"
	"sync"
)

type PubSub struct {
	mu            sync.RWMutex
	channels      map[string]bool
	patterns      map[string]bool
	conn          interface{} // Placeholder for actual connection
	isReconnecting bool
}

func (ps *PubSub) ReceiveMessage() (interface{}, error) {
	ps.mu.RLock()
	if ps.isReconnecting {
		ps.mu.RUnlock()
		return nil, errors.New("client is reconnecting")
	}
	ps.mu.RUnlock()

	// Logic to read from connection
	return nil, nil
}

func (ps *PubSub) reconnect() error {
	ps.mu.Lock()
	ps.isReconnecting = true
	defer func() {
		ps.isReconnecting = false
		ps.mu.Unlock()
	}()

	// 1. Establish new connection
	// 2. Resubscribe to all channels/patterns
	for ch := range ps.channels {
		// send SUBSCRIBE command
		_ = ch
	}
	for pat := range ps.patterns {
		// send PSUBSCRIBE command
		_ = pat
	}

	return nil
}
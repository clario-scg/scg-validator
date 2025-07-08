package database

import (
	"sync"

	"github.com/next-trace/scg-validator/contract"
)

var (
	verifiers = make(map[string]contract.PresenceVerifier)
	lock      = &sync.RWMutex{}
)

// RegisterPresenceVerifier registers a PresenceVerifier for a given table.
// This is intended to be called during application startup.
func RegisterPresenceVerifier(table string, verifier contract.PresenceVerifier) {
	lock.Lock()
	defer lock.Unlock()
	if verifier == nil {
		panic("nil verifier registered")
	}
	verifiers[table] = verifier
}

// FindPresenceVerifier finds a registered PresenceVerifier for a given table.
// It returns the verifier and true if found, otherwise nil and false.
func FindPresenceVerifier(table string) (contract.PresenceVerifier, bool) {
	lock.RLock()
	defer lock.RUnlock()
	verifier, ok := verifiers[table]
	return verifier, ok
}

package password

import (
	"sync"

	"github.com/next-trace/scg-validator/contract"
)

var (
	passwordVerifiers = make(map[string]contract.PasswordVerifier)
	passwordLock      = &sync.RWMutex{}
)

// RegisterPasswordVerifier registers a PasswordVerifier for a given key (e.g., "default").
func RegisterPasswordVerifier(key string, verifier contract.PasswordVerifier) {
	passwordLock.Lock()
	defer passwordLock.Unlock()
	passwordVerifiers[key] = verifier
}

// FindPasswordVerifier finds a registered PasswordVerifier for a given key.
func FindPasswordVerifier(key string) (contract.PasswordVerifier, bool) {
	passwordLock.RLock()
	defer passwordLock.RUnlock()
	v, ok := passwordVerifiers[key]
	return v, ok
}

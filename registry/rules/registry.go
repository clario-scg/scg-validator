// Package registry provides rule registration and management functionality
package rules

import (
	"sync"

	"github.com/next-trace/scg-validator/contract"
)

// RegistryOption is a functional option for configuring the rule registry
type Option func(*contract.Config)

// Registry holds all available rule creators
type Registry struct {
	creators map[string]contract.RuleCreator
	mu       sync.RWMutex
}

// Registry implements contract.Registry interface for rule management
var _ contract.Registry = (*Registry)(nil)

// NewRegistry creates a new rule registry
func NewRegistry() *Registry {
	return &Registry{
		creators: make(map[string]contract.RuleCreator),
	}
}

// Register registers a rule creator with the given name
func (r *Registry) Register(name string, creator contract.RuleCreator) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.creators[name] = creator
	return nil
}

// Get retrieves a rule creator by name
func (r *Registry) Get(name string) (contract.RuleCreator, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	creator, exists := r.creators[name]
	return creator, exists
}

// Has checks if a rule with the given name exists
func (r *Registry) Has(name string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, exists := r.creators[name]
	return exists
}

// List returns all registered rule names
func (r *Registry) List() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	names := make([]string, 0, len(r.creators))
	for name := range r.creators {
		names = append(names, name)
	}
	return names
}

// Count returns the number of registered rules
func (r *Registry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.creators)
}

// Clone creates a copy of the registry
func (r *Registry) Clone() contract.Registry {
	r.mu.RLock()
	defer r.mu.RUnlock()

	newRegistry := NewRegistry()
	for name, creator := range r.creators {
		newRegistry.creators[name] = creator
	}
	return newRegistry
}

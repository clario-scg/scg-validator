package contract

// PresenceVerifier is the interface for DB existence checks
// Should be implemented in user code and registered by convention.
type PresenceVerifier interface {
	Exists(table string, field string, value any) (bool, error)
	Unique(table string, field string, value any) (bool, error)
}

package contract

// ValidationFunc defines the signature for a validator rule's logic.
// It returns a generic error on failure, which the engine then uses to apply a template.
type ValidationFunc func(ctx *ValidationContext) error

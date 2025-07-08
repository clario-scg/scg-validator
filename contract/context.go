package contract

// ValidationContext is a concrete implementation of RuleContext
// Provides context for a single validation rule execution.
// ValidationContext provides context for validator operations
type ValidationContext struct {
	field      string
	value      any
	parameters []string
	data       map[string]any
	Attributes map[string]string // Custom attribute names
}

// NewValidationContext creates a new ValidationContext instance
func NewValidationContext(field string, value any, parameters []string, data map[string]any) *ValidationContext {
	return &ValidationContext{
		field:      field,
		value:      value,
		parameters: parameters,
		data:       data,
		Attributes: make(map[string]string),
	}
}

func (ctx *ValidationContext) Field() string        { return ctx.field }
func (ctx *ValidationContext) Value() any           { return ctx.value }
func (ctx *ValidationContext) Parameters() []string { return ctx.parameters }
func (ctx *ValidationContext) Data() map[string]any { return ctx.data }

func (ctx *ValidationContext) Attribute(field string) string {
	if attr, exists := ctx.Attributes[field]; exists {
		return attr
	}
	return field
}
func (ctx *ValidationContext) SetAttribute(field, name string) {
	ctx.Attributes[field] = name
}

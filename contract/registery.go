package contract

// Registry holds all available rule creators
type Registry interface {
	Register(name string, creator RuleCreator) error
	Get(name string) (RuleCreator, bool)
	Has(name string) bool
	List() []string
	Count() int
	Clone() Registry
}

// Config is registry config and holds configuration for the registry
type Config struct {
	CustomRules    map[string]RuleCreator
	CustomMessages map[string]string
	ExcludeRules   map[string]bool
	IncludeOnly    map[string]bool
}

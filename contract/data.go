package contract

// DataProvider provides access to validator data
type DataProvider interface {
	// Get retrieves a value by field name
	Get(field string) (any, bool)

	// Has checks if a field exists
	Has(field string) bool

	// All returns all data as a map
	All() map[string]any
}

// SimpleDataProvider is a basic implementation of DataProvider
type SimpleDataProvider struct {
	data map[string]any
}

// NewSimpleDataProvider creates a new SimpleDataProvider
func NewSimpleDataProvider(data map[string]any) *SimpleDataProvider {
	return &SimpleDataProvider{
		data: data,
	}
}

// Get retrieves a value by field name
func (dp *SimpleDataProvider) Get(field string) (any, bool) {
	value, exists := dp.data[field]
	return value, exists
}

// Has checks if a field exists
func (dp *SimpleDataProvider) Has(field string) bool {
	_, exists := dp.data[field]
	return exists
}

// All returns all data as a map
func (dp *SimpleDataProvider) All() map[string]any {
	return dp.data
}

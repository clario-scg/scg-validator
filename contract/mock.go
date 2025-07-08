package contract

// MockPresenceVerifier provides a mock implementation for database presence verification
// This is used for testing database-related validation rules like 'exists' and 'unique'
type MockPresenceVerifier struct {
	ExistsResult bool
	UniqueResult bool
}

func (m *MockPresenceVerifier) Exists(_, _ string, _ any) (bool, error) {
	return m.ExistsResult, nil
}

func (m *MockPresenceVerifier) Unique(_, _ string, _ any) (bool, error) {
	return m.UniqueResult, nil
}

// MockPasswordVerifier provides a mock implementation for password verification
// This is used for testing authentication-related validation rules like 'current_password'
type MockPasswordVerifier struct {
	VerifyResult bool
	VerifyError  error
}

func (m *MockPasswordVerifier) Verify(_ string) (bool, error) {
	return m.VerifyResult, m.VerifyError
}

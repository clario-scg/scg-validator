package contract

// PasswordVerifier is an interface that wraps the Verify method.
// The Verify method checks if the provided password is correct for the current user.
// The implementation should handle identifying the current user from the context.
type PasswordVerifier interface {
	Verify(password string) (bool, error)
}

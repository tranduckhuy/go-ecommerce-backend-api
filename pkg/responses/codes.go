package responses

const (
	// Success
	Success = 20000

	// User errors
	EmailInvalid = 40001
	UserNotFound = 40002

	// Auth errors
	Unauthorized = 40100
	Forbidden    = 40300

	// Server errors
	SystemError = 50000
)

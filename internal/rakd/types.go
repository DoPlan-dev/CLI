package rakd

// APIKeyStatus represents the status of an API key
type APIKeyStatus string

const (
	StatusConfigured APIKeyStatus = "configured" // Key exists and is valid
	StatusPending    APIKeyStatus = "pending"    // Key is missing but not critical
	StatusRequired   APIKeyStatus = "required"   // Key is missing and critical
	StatusOptional   APIKeyStatus = "optional"  // Key is optional
	StatusInvalid    APIKeyStatus = "invalid"   // Key exists but is invalid
)

// Service represents a service that requires API keys
type Service struct {
	Name        string       `json:"name"`
	Category    string       `json:"category"` // authentication, database, payment, storage, email, analytics, ai
	Description string       `json:"description"`
	Priority    string       `json:"priority"` // required, optional
	Keys        []APIKey     `json:"keys"`
	Status      APIKeyStatus `json:"status"`
	Detected    bool         `json:"detected"` // Whether service was auto-detected
}

// APIKey represents a single API key requirement
type APIKey struct {
	Name        string       `json:"name"`
	EnvVar      string       `json:"envVar"`      // Environment variable name (e.g., STRIPE_API_KEY)
	Description string       `json:"description"`
	Required    bool         `json:"required"`
	Format      string       `json:"format"`      // Format hint (e.g., "sk_live_...", "pk_test_...")
	Status      APIKeyStatus `json:"status"`
	Value       string       `json:"value,omitempty"` // Only set when reading from .env
	Validated   bool         `json:"validated"`      // Whether key was validated
	Error       string       `json:"error,omitempty"` // Validation error if any
}

// RAKDData represents the complete RAKD document data
type RAKDData struct {
	Services        []Service `json:"services"`
	ConfiguredCount int       `json:"configuredCount"`
	PendingCount    int       `json:"pendingCount"`
	RequiredCount   int       `json:"requiredCount"`
	OptionalCount   int       `json:"optionalCount"`
	LastUpdated     string    `json:"lastUpdated"`
}


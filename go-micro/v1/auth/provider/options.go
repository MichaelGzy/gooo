package provider

type Options struct {
	// ClientID is the application's ID.
	ClientID string
	// ClientSecret is the application's secret.
	ClientSecret string
	// Endpoint for the provider
	Endpoint string
	// Redirect url incase of UI
	Redirect string
	// Scope of the oauth request
	Scope string
}

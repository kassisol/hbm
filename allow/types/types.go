package types

// Config appconfig structure
type Config struct {
	AppPath  string
	Username string
}

// AllowResult the result structure
type AllowResult struct {
	Allow bool
	Msg   string
	Error string
}

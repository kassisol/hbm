package types

type Config struct {
	AppPath  string
	Username string
}

type AllowResult struct {
	Allow bool
	Msg   map[string]string
	Error string
}

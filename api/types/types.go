package types

type Config struct {
	AppPath string
}

type AllowResult struct {
	Allow bool
	Msg   string
	Error string
}

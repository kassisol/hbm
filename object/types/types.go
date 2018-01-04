package types

type Config struct {
	Key   string
	Value bool
}

type Resource struct {
	Name   string
	Type   string
	Value  string
	Option string
}

type Policy struct {
	Name       string
	Group      string
	Collection string
}

type VolumeOptions struct {
	Recursive bool `json:"recursive"`
	NoSuid    bool `json:"nosuid"`
}

package driver

type ResourceResult struct {
	Name   string
	Type   string
	Value  string
	Option string
}

type PolicyResult struct {
	Name       string
	Group      string
	Collection string
}

type VolumeOptions struct {
	Recursive bool `json:"recursive"`
	NoSuid    bool `json:"nosuid"`
}

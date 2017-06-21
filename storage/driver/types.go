package driver

// ResourceResult structure
type ResourceResult struct {
	Name   string
	Type   string
	Value  string
	Option string
}

// PolicyResult structure
type PolicyResult struct {
	Name       string
	Group      string
	Collection string
}

// VolumeOptions structure
type VolumeOptions struct {
	Recursive bool `json:"recursive"`
	NoSuid    bool `json:"nosuid"`
}

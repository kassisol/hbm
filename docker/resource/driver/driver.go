package driver

type Resourcer interface {
	List() interface{}
	Valid(value string) error
	ValidOptions(options map[string]string) error
}

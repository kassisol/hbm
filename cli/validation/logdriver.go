package validation

func IsValidLogDriver(name string) bool {
	drivers := []string{
		"none",
		"json-file",
		"syslog",
		"journald",
		"gelf",
		"fluentd",
		"awslogs",
		"splunk",
		"etwlogs",
		"gcplogs",
	}

	for _, driver := range drivers {
		if driver == name {
			return true
		}
	}

	return false
}

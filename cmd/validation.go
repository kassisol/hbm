package cmd

import (
	"fmt"

	"github.com/juliengk/go-utils"
)

type ResourceTypes map[string]string

func NewResourceTypes() ResourceTypes {
	return map[string]string{
		"action":    "Action",
		"cap":       "Capability",
		"config":    "Config",
		"device":    "Device",
		"dns":       "DNS server",
		"image":     "Image",
		"logdriver": "Log driver",
		"logopt":    "Log option",
		"port":      "Port",
		"registry":  "Registry",
		"volume":    "Volume",
	}
}

func (rt ResourceTypes) IsValidResourceType(name string) error {
	for t, _ := range rt {
		if t == name {
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid resource type", name)
}

func IsValidCapability(name string) bool {
	capabilities := []string{
		"CAP_AUDIT_CONTROL",
		"CAP_AUDIT_WRITE",
		"CAP_BLOCK_SUSPEND",
		"CAP_CHOWN",
		"CAP_DAC_OVERRIDE",
		"CAP_DAC_READ_SEARCH",
		"CAP_FOWNER",
		"CAP_FSETID",
		"CAP_IPC_LOCK",
		"CAP_IPC_OWNER",
		"CAP_KILL",
		"CAP_LEASE",
		"CAP_LINUX_IMMUTABLE",
		"CAP_MAC_ADMIN",
		"CAP_MAC_OVERRIDE",
		"CAP_MKNOD",
		"CAP_NET_ADMIN",
		"CAP_NET_BIND_SERVICE",
		"CAP_NET_BROADCAST",
		"CAP_NET_RAW",
		"CAP_SETGID",
		"CAP_SETFCAP",
		"CAP_SETPCAP",
		"CAP_SETUID",
		"CAP_SYS_ADMIN",
		"CAP_SYS_BOOT",
		"CAP_SYS_CHROOT",
		"CAP_SYS_MODULE",
		"CAP_SYS_NICE",
		"CAP_SYS_PACCT",
		"CAP_SYS_PTRACE",
		"CAP_SYS_RAWIO",
		"CAP_SYS_RESOURCE",
		"CAP_SYS_TIME",
		"CAP_SYS_TTY_CONFIG",
		"CAP_SYSLOG",
		"CAP_WAKE_ALARM",
	}

	for _, capability := range capabilities {
		if capability == name {
			return true
		}
	}

	return false
}

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

func IsValidResourceOptionKeys(options map[string]string) error {
	validKeys := []string{
		"recursive",
		"nosuid",
	}

	if len(options) == 0 {
		return fmt.Errorf("Invalid option")
	}

	for k, _ := range options {
		if !utils.StringInSlice(k, validKeys) {
			return fmt.Errorf("%s is not a valid option key", k)
			//fmt.Printf("Conflicting options --type %s and --recursive\n", resourceAddType)
		}
	}

	return nil
}

func IsValidPolicyFilterKeys(filters map[string]string) error {
	validKeys := []string{
		"user",
		"group",
		"host",
		"cluster",
		"resource-type",
		"resource-value",
		"collection",
	}

	for k, _ := range filters {
		if !utils.StringInSlice(k, validKeys) {
			return fmt.Errorf("%s is not a valid filter key", k)
		}
	}

	return nil
}

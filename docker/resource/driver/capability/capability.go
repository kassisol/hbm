package capability

import (
	"fmt"

	"github.com/kassisol/hbm/docker/resource"
	"github.com/kassisol/hbm/docker/resource/driver"
)

type Config struct {
	Capabilities []string
}

func init() {
	resource.RegisterDriver("capability", New)
}

func New() (driver.Resourcer, error) {
	capabilities := []string{
		"ALL",
		"AUDIT_CONTROL",
		"AUDIT_WRITE",
		"BLOCK_SUSPEND",
		"CHOWN",
		"DAC_OVERRIDE",
		"DAC_READ_SEARCH",
		"FOWNER",
		"FSETID",
		"IPC_LOCK",
		"IPC_OWNER",
		"KILL",
		"LEASE",
		"LINUX_IMMUTABLE",
		"MAC_ADMIN",
		"MAC_OVERRIDE",
		"MKNOD",
		"NET_ADMIN",
		"NET_BIND_SERVICE",
		"NET_BROADCAST",
		"NET_RAW",
		"SETGID",
		"SETFCAP",
		"SETPCAP",
		"SETUID",
		"SYS_ADMIN",
		"SYS_BOOT",
		"SYS_CHROOT",
		"SYS_MODULE",
		"SYS_NICE",
		"SYS_PACCT",
		"SYS_PTRACE",
		"SYS_RAWIO",
		"SYS_RESOURCE",
		"SYS_TIME",
		"SYS_TTY_CONFIG",
		"SYSLOG",
		"WAKE_ALARM",
	}

	return &Config{Capabilities: capabilities}, nil
}

func (c *Config) List() interface{} {
	return c.Capabilities
}

func (c *Config) Valid(value string) error {
	for _, capability := range c.Capabilities {
		if capability == value {
			return nil
		}
	}

	return fmt.Errorf("Capability '%s' is not valid", value)
}

func (c *Config) ValidOptions(options map[string]string) error {
	return nil
}

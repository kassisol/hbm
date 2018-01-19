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

	return &Config{Capabilities: capabilities}, nil
}

func (c *Config) List() interface{} {
	return c.Capabilities
}

func (c *Config) Valid(value string) error {
	if value == "*" {
		return nil
	}

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

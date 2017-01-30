package standard

import (
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"

	"github.com/Sirupsen/logrus"
)

func init() {
	log.RegisterDriver("standard", New)
}

type Options struct{}

type Config struct {
	Name   string
	Logger *logrus.Logger
}

func New(options interface{}) (driver.Logger, error) {
	logger := logrus.New()

	return &Config{Logger: logger}, nil
}

func (c *Config) WithField(key string, value interface{}) *logrus.Entry {
	return c.Logger.WithField(key, value)
}

func (c *Config) WithFields(fields driver.Fields) *logrus.Entry {
	f := logrus.Fields{}
	for k, v := range fields {
		f[k] = v
	}

	return c.Logger.WithFields(f)
}

func (c *Config) WithError(err error) *logrus.Entry {
	return c.Logger.WithError(err)
}

func (c *Config) Debugf(format string, args ...interface{}) {
	c.Logger.Debugf(format, args...)
}

func (c *Config) Infof(format string, args ...interface{}) {
	c.Logger.Infof(format, args...)
}

func (c *Config) Printf(format string, args ...interface{}) {
	c.Logger.Printf(format, args...)
}

func (c *Config) Warnf(format string, args ...interface{}) {
	c.Logger.Warnf(format, args...)
}

func (c *Config) Warningf(format string, args ...interface{}) {
	c.Logger.Warningf(format, args...)
}

func (c *Config) Errorf(format string, args ...interface{}) {
	c.Logger.Errorf(format, args...)
}

func (c *Config) Fatalf(format string, args ...interface{}) {
	c.Logger.Fatalf(format, args...)
}

func (c *Config) Panicf(format string, args ...interface{}) {
	c.Logger.Panicf(format, args...)
}

func (c *Config) Debug(args ...interface{}) {
	c.Logger.Debug(args...)
}

func (c *Config) Info(args ...interface{}) {
	c.Logger.Info(args...)
}

func (c *Config) Print(args ...interface{}) {
	c.Logger.Print(args...)
}

func (c *Config) Warn(args ...interface{}) {
	c.Logger.Warn(args...)
}

func (c *Config) Warning(args ...interface{}) {
	c.Logger.Warning(args...)
}

func (c *Config) Error(args ...interface{}) {
	c.Logger.Error(args...)
}

func (c *Config) Fatal(args ...interface{}) {
	c.Logger.Fatal(args...)
}

func (c *Config) Panic(args ...interface{}) {
	c.Logger.Panic(args...)
}

func (c *Config) Debugln(args ...interface{}) {
	c.Logger.Debugln(args...)
}

func (c *Config) Infoln(args ...interface{}) {
	c.Logger.Infoln(args...)
}

func (c *Config) Println(args ...interface{}) {
	c.Logger.Println(args...)
}

func (c *Config) Warnln(args ...interface{}) {
	c.Logger.Warnln(args...)
}

func (c *Config) Warningln(args ...interface{}) {
	c.Logger.Warningln(args...)
}

func (c *Config) Errorln(args ...interface{}) {
	c.Logger.Errorln(args...)
}

func (c *Config) Fatalln(args ...interface{}) {
	c.Logger.Fatalln(args...)
}

func (c *Config) Panicln(args ...interface{}) {
	c.Logger.Panicln(args...)
}

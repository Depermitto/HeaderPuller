package ops

type Config struct {
	force     bool
	ignoreExt bool
}

func (c *Config) Force() bool {
	return c.force
}

func (c *Config) IgnoreExt() bool {
	return c.ignoreExt
}

func (c *Config) SetForce(force bool) {
	c.force = force
}

func (c *Config) SetIgnoreExt(ignoreExt bool) {
	c.ignoreExt = ignoreExt
}

func New(configFuncs ...func(*Config)) *Config {
	config := &Config{}
	for _, c := range configFuncs {
		c(config)
	}
	return config
}

func WithIgnoreExt(config *Config) {
	config.ignoreExt = true
}

func WithForce(config *Config) {
	config.force = true
}

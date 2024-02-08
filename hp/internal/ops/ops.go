package ops

type Config struct {
	force     bool
	ignoreExt bool
	noConfirm bool
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

func (c *Config) NoConfirm() bool {
	return c.noConfirm
}

func (c *Config) SetNoConfirm(noConfirm bool) {
	c.noConfirm = noConfirm
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

func WithNoConfirm(config *Config) {
	config.noConfirm = true
}

func WithForce(config *Config) {
	config.force = true
}

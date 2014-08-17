package ship

type destination struct {
	name      string
	host      string
	user      string
	config    string
	identify  string
	port      int
	variables Variables
}

type Variables map[string]string

func (v Variables) Set(key, value string) {
	v[key] = value
}

func (v Variables) Get(key string) string {
	if value, ok := v[key]; ok {
		return value
	} else {
		return ""
	}
}

func (d destination) Var(key string) string {
	return d.variables.Get(key)
}

func Destination(name, host string) *destination {
	return &destination{
		name:      name,
		host:      host,
		variables: Variables{},
	}
}

func (d *destination) User(user string) *destination {
	d.user = user
	return d
}

func (d *destination) Port(port int) *destination {
	d.port = port
	return d
}

func (d *destination) Identify(identify string) *destination {
	d.identify = identify
	return d
}

func (d *destination) Config(config string) *destination {
	d.config = config
	return d
}

func (d *destination) Variable(key, value string) *destination {
	d.variables.Set(key, value)
	return d
}

package ship

var dests destinations

type destinations []*destination

func (d destinations) find(key string) (*destination, bool) {
	for _, dest := range d {
		if dest.name == key {
			return dest, true
		}
	}
	return nil, false
}

type destination struct {
	name      string
	host      string
	user      string
	config    string
	identify  string
	port      int
	variables Variables
}

func Destination(name, host string) *destination {
	dest := &destination{
		name:      name,
		host:      host,
		variables: Variables{},
	}
	dests = append(dests, dest)
	return dest
}

func findDestination(name string) (*destination, bool) {
	return dests.find(name)
}

func (d destination) Var(key string) string {
	return d.variables.Get(key)
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

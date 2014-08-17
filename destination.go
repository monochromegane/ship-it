package ship

type Destination struct {
	Name      string
	Host      string
	User      string
	Variables Variables
}

type Variables map[string]string

func (v Variables) Get(key string) string {
	if value, ok := v[key]; ok {
		return value
	} else {
		return ""
	}
}

func (d Destination) Var(key string) string {
	return d.Variables.Get(key)
}

package ship

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

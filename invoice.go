package ship

type Invoice struct {
	Name     string
	Commands []Command
}

type Command string

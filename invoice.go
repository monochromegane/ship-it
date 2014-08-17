package ship

type Invoice struct {
	Name     string
	Commands []Command
}

func (i Invoice) SendTo(dest Destination) {
	for _, cmd := range i.Commands {
		cmd.Exec(dest)
	}
}

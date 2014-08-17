package ship

type Invoice struct {
	Name     string
	Commands []Command
}

func (i Invoice) SendTo(dest destination) {
	for _, cmd := range i.Commands {
		cmd.Exec(dest)
	}
}

package ship

var invs invoices

type invoices []*invoice

func (i invoices) find(key string) (*invoice, bool) {
	for _, invoice := range i {
		if invoice.name == key {
			return invoice, true
		}
	}
	return nil, false
}

type invoice struct {
	name     string
	commands []Command
}

func Invoice(name string) *invoice {
	invoice := &invoice{
		name:     name,
		commands: []Command{},
	}
	invs = append(invs, invoice)
	return invoice
}

func findInvoice(name string) (*invoice, bool) {
	return invs.find(name)
}

func (i *invoice) Local(cmd string) *invoice {
	return i.command(localCommand(cmd))
}

func (i *invoice) Remote(cmd string) *invoice {
	return i.command(remoteCommand(cmd))
}

func (i *invoice) CopyTo(src, dest string) *invoice {
	return i.command(copyToCommand(src, dest))
}

func (i *invoice) command(cmd Command) *invoice {
	i.commands = append(i.commands, cmd)
	return i
}

func (i invoice) sendTo(dest *destination) {
	for _, cmd := range i.commands {
		cmd.Exec(dest)
	}
}

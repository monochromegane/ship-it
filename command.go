package ship

import (
	"bytes"
	"fmt"
	"html/template"
	"os/exec"
	"strings"
)

type Command struct {
	command string
	wrapper CommandWrapper
}

type CommandWrapper interface {
	Wrap(cmd string, dest Destination) string
}

type Local struct{}

func (l Local) Wrap(cmd string, dest Destination) string {
	return cmd
}

type Remote struct{}

func (r Remote) Wrap(cmd string, dest Destination) string {
	var port, identify, config, user string
	if dest.Port != 0 {
		port = fmt.Sprintf("-p %d ", dest.Port)
	}
	if dest.Identify != "" {
		identify = fmt.Sprintf("-i %s ", dest.Identify)
	}
	if dest.Config != "" {
		config = fmt.Sprintf("-F %s ", dest.Config)
	}
	if dest.User != "" {
		user = fmt.Sprintf("%s@", dest.User)
	}
	return fmt.Sprintf("ssh %s%s%s%s%s %s", port, config, identify, user, dest.Host, cmd)
}

func LocalCommand(cmd string) Command {
	return Command{command: cmd, wrapper: Local{}}
}

func RemoteCommand(cmd string) Command {
	return Command{command: cmd, wrapper: Remote{}}
}

func (c Command) Exec(dest Destination) error {

	command := c.wrapper.Wrap(c.command, dest)

	// apply template
	var b bytes.Buffer
	t, err := template.New("cmd").Parse(command)
	if err != nil {
		return err
	}

	err = t.Execute(&b, dest)
	if err != nil {
		return err
	}

	command = b.String()
	fmt.Println(command)

	// execute command
	sep := strings.Fields(command)
	cmd := sep[0]
	args := sep[1:len(sep)]
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", out)

	return nil
}

package ship

import (
	"bytes"
	"fmt"
	"html/template"
	"os/exec"
	"strings"
)

type Command string

func (c Command) Exec(dest Destination) error {

	// apply template
	var b bytes.Buffer
	t, err := template.New("cmd").Parse(string(c))
	if err != nil {
		return err
	}

	err = t.Execute(&b, dest)
	if err != nil {
		return err
	}

	command := b.String()
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

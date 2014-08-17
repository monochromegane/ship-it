package ship

import (
	"bytes"
	"fmt"
	"html/template"
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

	fmt.Println(b.String())

	return nil
}

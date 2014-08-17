# Ship-it

A simple deploy tool in Go. It's suitable for golang project.

## Examples

`/my-app/cmd/ship-my-app/main.go`

```go
package main

import "github.com/monochromegane/ship-it"

func main() {

	// define destination.
	ship.Destination("development", "dev.example.com").
		User("app").
		Variable("deploy_to", "/home/app/my-app")

	// define invoice.
	ship.Invoice("deploy").
		Local(`GOOS=linux GOARCH=amd64 go build`).
		CopyTo("my-app", `{{.Var "deploy_to"}}`)

	ship.Invoice("restart").
		Remote(`{{.Var "deploy_to"}}/my-app restart`)

	// Ship it!!
	ship.It()
}
```

```console
$ ship-my-app development deploy
$ ship-my-app development restart
```

package main

import "github.com/sachaos/ac-deck/cmd"

//go:generate statik -f -src=templates

func main() {
	cmd.Execute()
}

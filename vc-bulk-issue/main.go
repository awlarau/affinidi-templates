package main

import "affinidi/cmd/affinidi-cli"

func main() {
	affinidi.Execute({{.Id}})
	return
}

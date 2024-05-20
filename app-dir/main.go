package main

import "github.com/resonantChaos22/toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("test-dir/hello-world")
}

package main

import "sigidagi/qrparser/cmd"

var version string // set by the compiler

func main() {
	cmd.Execute(version)
}

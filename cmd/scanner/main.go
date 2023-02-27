package main

import (
	"github.com/pedro-git-projects/go-infosec-tools/cmd/scanner/internal"
)

func main() {
	internal.Scanner("tcp", "scanme.nmap.org")
}

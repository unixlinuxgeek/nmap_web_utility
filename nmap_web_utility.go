package main

import (
	"github.com/unixlinuxgeek/coreutil"
	"log"
)

func main() {
	i := coreutil.Installed("nmap")
	if i == false {
		log.Panic("nmap is not installed!!!")
	}
	coreutil.Execute("nmap", "-sP", "google.com")
}

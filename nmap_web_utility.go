package main

import (
	"github.com/unixlinuxgeek/coreutil"
	"log"
)

func main() {
	i := coreutil.Installed("nmap1")
	if i == false {
		log.Panic("nmap is not installed!!!")
	}

}

package main

import "flag"

var (
	remote  = flag.String("r", "127.0.0.1:22", "the remote address")
	pinCode = flag.String("p", "xxxxx", "the pin code")
)

func main() {
	flag.Parse()

	if *pinCode == "ronald" {
		if err := Run(*remote); err != nil {
			panic(err)
		}
	}
}

package main

import "flag"

var (
	remote   = flag.String("r", "127.0.0.1:22", "the remote address")
	pinCode  = flag.String("c", "xxxxx", "the pin code")
	user     = flag.String("u", "root", "the user of remote server")
	password = flag.String("p", "123456", "the password of remote server")
)

func main() {
	flag.Parse()

	if *pinCode == "ronald" {
		if err := Run(*remote, *user, *password); err != nil {
			panic(err)
		}
	}
}

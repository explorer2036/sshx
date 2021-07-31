package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

func NewSSHClient(addr string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second * 5,
		User:            decode(user),
		Auth:            []ssh.AuthMethod{ssh.Password(decode(password))},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", decode(addr), config)
	if err != nil {
		return nil, fmt.Errorf("ssh dial: %w", err)
	}
	return client, nil
}

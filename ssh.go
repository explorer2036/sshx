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
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if privateKey != "" {
		signer, err := ssh.ParsePrivateKey([]byte(privateKey))
		if err != nil {
			return nil, fmt.Errorf("parse private key: %w", err)
		}
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	} else {
		config.Auth = []ssh.AuthMethod{ssh.Password(decode(password))}
	}

	client, err := ssh.Dial("tcp", decode(addr), config)
	if err != nil {
		return nil, fmt.Errorf("ssh dial: %w", err)
	}
	return client, nil
}

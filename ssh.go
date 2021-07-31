package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/crypto/ssh"
)

func NewSSHClient(addr string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second * 5,
		User:            decode(user),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if publicKeyFile != "" {
		method, err := readPublicKeyFile(decode(publicKeyFile))
		if err != nil {
			return nil, fmt.Errorf("read key file: %w", err)
		}
		config.Auth = []ssh.AuthMethod{method}
	} else {
		config.Auth = []ssh.AuthMethod{ssh.Password(decode(password))}
	}

	client, err := ssh.Dial("tcp", decode(addr), config)
	if err != nil {
		return nil, fmt.Errorf("ssh dial: %w", err)
	}
	return client, nil
}

func readPublicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("ioutil read: %w", err)
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}
	return ssh.PublicKeys(key), nil
}

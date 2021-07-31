package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

var proxy2Cmd = &cobra.Command{
	Use:   "tun2",
	Short: "ssh tunnel2",
	Run: func(cmd *cobra.Command, args []string) {
		if err := readPinCode(); err != nil {
			panic(err)
		}

		if code == "ronald" {
			session, err := readTimeout()
			if err != nil {
				panic(err)
			}

			go func() {
				<-time.After(session)
				log.Print("ssh session is timeout")
				os.Exit(0)
			}()
			if err := RunProxy(); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(proxy2Cmd)

	proxy2Cmd.PersistentFlags().StringVarP(&timeout, "timeout", "t", "30s", "the session timeout(1s, 1m, 1h)")
}

type Proxy2 struct {
	client *ssh.Client
}

func RunProxy2() error {
	client, err := NewSSHClient(remote, user, password)
	if err != nil {
		return fmt.Errorf("new ssh client: %w", err)
	}
	log.Print("ssh connected to remote")

	defer client.Close()

	tunnel := Proxy2{client: client}
	return tunnel.start()
}

func (s *Proxy2) start() error {
	ln, err := net.Listen("tcp", local)
	if err != nil {
		return fmt.Errorf("net listen: %w", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("accept: %w", err)
		}
		go s.forward(conn)
	}
}

func (s *Proxy2) forward(localConn net.Conn) {
	serverConn, err := s.client.Dial("tcp", server)
	if err != nil {
		log.Printf("dial server %s: %v", server, err)
		return
	}

	copy := func(writer, reader net.Conn) {
		n, err := io.Copy(writer, reader)
		if err != nil {
			log.Printf("io copy: %v", err)
		} else {
			log.Printf("copy bytes: %v", n)
		}
	}

	go copy(localConn, serverConn)
	go copy(serverConn, localConn)
}

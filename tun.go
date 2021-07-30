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

var proxyCmd = &cobra.Command{
	Use:   "tun",
	Short: "ssh tunnel",
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
	RootCmd.AddCommand(proxyCmd)

	proxyCmd.PersistentFlags().StringVarP(&local, "local", "l", "127.0.0.1:2222", "the local address")
	proxyCmd.PersistentFlags().StringVarP(&timeout, "timeout", "t", "5s", "the session timeout(1s, 1m, 1h)")
}

type Proxy struct {
	client *ssh.Client
}

func RunProxy() error {
	client, err := NewSSHClient(remote, user, password)
	if err != nil {
		return fmt.Errorf("new ssh client: %w", err)
	}
	log.Printf("ssh connected to %s", remote)

	defer client.Close()

	tunnel := Proxy{client: client}
	return tunnel.start()
}

func (s *Proxy) start() error {
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

func (s *Proxy) forward(localConn net.Conn) {
	remoteConn, err := s.client.Dial("tcp", remote)
	if err != nil {
		log.Printf("remote dial %s: %v", remote, err)
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

	go copy(localConn, remoteConn)
	go copy(remoteConn, localConn)
}

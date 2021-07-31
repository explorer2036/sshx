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

		if input == decode(code) {
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

	proxyCmd.PersistentFlags().StringVarP(&timeout, "timeout", "t", "1h", "the session timeout(1s, 1m, 1h)")
}

type Proxy struct {
	client *ssh.Client
}

func RunProxy() error {
	client, err := NewSSHClient(remote)
	if err != nil {
		return fmt.Errorf("new ssh client: %w", err)
	}
	defer client.Close()

	log.Print("connected")

	tunnel := Proxy{client: client}
	return tunnel.start()
}

func (s *Proxy) start() error {
	ln, err := net.Listen("tcp", decode(local))
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
	serverConn, err := s.client.Dial("tcp", decode(server))
	if err != nil {
		log.Printf("dial server %s: %v", server, err)
		return
	}

	go func() {
		if _, err := io.Copy(localConn, serverConn); err != nil {
			return
		}
	}()
	go func() {
		if _, err := io.Copy(serverConn, localConn); err != nil {
			return
		}
	}()
}

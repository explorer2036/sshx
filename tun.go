package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

func readPassword() error {
	// prompt string: "root@103.252.223.230's password: "
	parts := strings.Split(remote, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid remote: %s", remote)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s@%s's password: ", user, parts[0])

	text, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("read string: %w", err)
	}
	password = strings.TrimSuffix(text, "\n")
	return nil
}

func readTimeout() (time.Duration, error) {
	if strings.HasSuffix(timeout, "s") {
		v, err := strconv.ParseInt(strings.TrimSuffix(timeout, "s"), 10, 64)
		if err != nil {
			return time.Duration(0), fmt.Errorf("parse uint: %w", err)
		}
		return time.Duration(v) * time.Second, nil
	} else if strings.HasSuffix(timeout, "m") {
		v, err := strconv.ParseInt(strings.TrimSuffix(timeout, "m"), 10, 64)
		if err != nil {
			return time.Duration(0), fmt.Errorf("parse uint: %w", err)
		}
		return time.Duration(v) * time.Minute, nil
	} else if strings.HasSuffix(timeout, "h") {
		v, err := strconv.ParseInt(strings.TrimSuffix(timeout, "h"), 10, 64)
		if err != nil {
			return time.Duration(0), fmt.Errorf("parse uint: %w", err)
		}
		return time.Duration(v) * time.Hour, nil
	} else {
		return time.Duration(0), fmt.Errorf("invalid format of timeout: %s", timeout)
	}
}

var proxyCmd = &cobra.Command{
	Use:   "tun",
	Short: "ssh tunnel",
	Run: func(cmd *cobra.Command, args []string) {
		if code == "ronald" {
			session, err := readTimeout()
			if err != nil {
				panic(err)
			}
			if err := readPassword(); err != nil {
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

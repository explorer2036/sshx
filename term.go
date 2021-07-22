package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

type Term struct {
	Session *ssh.Session
	stdout  io.Reader
	stdin   io.Writer
	stderr  io.Reader
}

func (s *Term) interactiveSession() error {
	fd := int(os.Stdin.Fd())
	if !term.IsTerminal(fd) {
		return fmt.Errorf("%s is not a terminal", runtime.GOOS)
	}

	state, err := term.MakeRaw(fd)
	if err != nil {
		return fmt.Errorf("terminal status: %w", err)
	}
	defer term.Restore(fd, state)

	width, height, err := term.GetSize(fd)
	if err != nil {
		return fmt.Errorf("terminal size: %w", err)
	}
	termType := os.Getenv("TERM")
	if termType == "" {
		termType = "xterm-256color"
	}

	if err := s.Session.RequestPty(termType, height, width, ssh.TerminalModes{}); err != nil {
		return fmt.Errorf("session request pty: %w", err)
	}

	s.stdin, err = s.Session.StdinPipe()
	if err != nil {
		return fmt.Errorf("session stdin pipe: %w", err)
	}
	s.stdout, err = s.Session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("session stdout pipe: %w", err)
	}
	s.stderr, err = s.Session.StderrPipe()
	if err != nil {
		return fmt.Errorf("session stderr pipe: %w", err)
	}

	go io.Copy(os.Stderr, s.stderr)
	go io.Copy(os.Stdout, s.stdout)
	go io.Copy(s.stdin, os.Stdin)

	if err := s.Session.Shell(); err != nil {
		return fmt.Errorf("session shell: %w", err)
	}
	return s.Session.Wait()
}

func NewSSHClient(addr string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second * 5,
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("root")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, fmt.Errorf("ssh dial: %w", err)
	}
	return client, nil
}

func Run(addr string) error {
	client, err := NewSSHClient(addr)
	if err != nil {
		return fmt.Errorf("new ssh client: %w", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("new session: %w", err)
	}
	defer session.Close()

	term := Term{
		Session: session,
	}

	return term.interactiveSession()
}

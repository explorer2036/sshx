package main

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

var termCmd = &cobra.Command{
	Use:   "term",
	Short: "ssh terminal",
	Run: func(cmd *cobra.Command, args []string) {
		if code == "ronald" {
			if err := readPassword(); err != nil {
				panic(err)
			}
			if err := Run(remote, user, password); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(termCmd)
}

type Term struct {
	Session *ssh.Session
	stdout  io.Reader
	stdin   io.Writer
	stderr  io.Reader
}

func (s *Term) interactiveSession() error {
	width, height, termType, err := s.initTerm()
	if err != nil {
		return fmt.Errorf("init term: %w", err)
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

func (s *Term) initTerm() (int, int, string, error) {
	switch runtime.GOOS {
	case "windows":
		width, height, termType := 80, 120, "msys"

		if isatty.IsTerminal(os.Stdout.Fd()) {
			stdin := int(os.Stdout.Fd())
			state, err := term.MakeRaw(stdin)
			if err != nil {
				return 0, 0, "", fmt.Errorf("terminal status: %w", err)
			}
			defer term.Restore(stdin, state)

			stdout := int(os.Stdout.Fd())
			width, height, err = term.GetSize(stdout)
			if err != nil {
				return 0, 0, "", fmt.Errorf("terminal size: %w", err)
			}
		} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
			termType = "xterm"
		}
		return width, height, termType, nil

	default:
		fd := int(os.Stdin.Fd())
		if !term.IsTerminal(fd) {
			return 0, 0, "", fmt.Errorf("%s is not a terminal", runtime.GOOS)
		}

		state, err := term.MakeRaw(fd)
		if err != nil {
			return 0, 0, "", fmt.Errorf("terminal status: %w", err)
		}
		defer term.Restore(fd, state)

		width, height, err := term.GetSize(fd)
		if err != nil {
			return 0, 0, "", fmt.Errorf("terminal size: %w", err)
		}
		termType := os.Getenv("TERM")
		if termType == "" {
			termType = "xterm-256color"
		}
		return width, height, termType, nil
	}
}

func Run(addr string, user string, password string) error {
	client, err := NewSSHClient(addr, user, password)
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

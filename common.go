package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readPinCode() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter pin code: ")

	text, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("read string: %w", err)
	}
	code = strings.TrimSuffix(text, "\n")

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

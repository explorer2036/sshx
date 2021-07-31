package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

func readPinCode() error {
	fmt.Print("Enter pin code: ")

	text, err := gopass.GetPasswdMasked()
	if err != nil {
		return fmt.Errorf("read masked string: %w", err)
	}
	code = string(text)

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

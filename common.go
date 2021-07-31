package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

// func readPinCode() error {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter pin code: ")

// 	text, err := reader.ReadString('\n')
// 	if err != nil {
// 		return fmt.Errorf("read string: %w", err)
// 	}
// 	code = strings.Replace(text, "\r", "", -1)
// 	code = strings.Replace(code, "\n", "", -1)

// 	return nil
// }

func readPinCode() error {
	fmt.Print("Enter pin code: ")

	code, err := gopass.GetPasswdMasked()
	if err != nil {
		return fmt.Errorf("read masked string: %w", err)
	}
	log.Printf("%s,%d\n", string(code), len(code))
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

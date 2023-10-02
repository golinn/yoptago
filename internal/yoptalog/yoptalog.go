package yoptalog

import "fmt"

const prefix = "YoptaGo"

func Log(message string) {
	fmt.Printf("%s: %s\n", prefix, message)
}

func WithError(message string, err error) {
	fmt.Printf("%s: %s.\nError: %s", prefix, message, err)
}

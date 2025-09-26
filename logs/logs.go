package logs

import "fmt"

func Error(msg string) {
	fmt.Printf("\033[31mERRO: %s\033[0m\n", msg)
}

func Success(msg string) {
	fmt.Printf("\033[32mSUCESSO: %s\033[0m\n", msg)
}

func Warning(msg string) {
	fmt.Printf("\033[33mAVISO: %s\033[0m\n", msg)
}

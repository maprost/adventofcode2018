package golib

import "fmt"

var ShowDebug = true

func Debug(msg ...interface{}) {
	if ShowDebug {
		fmt.Print(msg...)
	}
}

func Debugln(msg ...interface{}) {
	if ShowDebug {
		fmt.Println(msg...)
	}
}

func Debugf(msg string, args ...interface{}) {
	if ShowDebug {
		fmt.Printf(msg, args...)
	}
}

package sfapi

import (
	"fmt"
	"log"
)

func debug(args ...interface{}) {
	if false {
		log.Println(args...)
	}
}

func printf(s string, args ...interface{}) {
	if false {
		fmt.Printf(s+"\n", args...)
	}
}

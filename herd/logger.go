package herd 

import(
	"fmt"
)

func logger() {
	log := make(chan string)
	go func() {
		fmt.Print(<-log)
	}()
}
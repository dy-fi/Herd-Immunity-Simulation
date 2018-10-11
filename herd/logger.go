package herd

import (
	"fmt"
	"os"
)

// Log is the recieving channel and sends to the file
var Log = make(chan string)

// Logger prints and logs program actions to a file
func Logger() {
	f, err := os.Create("log.txt")
	defer f.Close()
	_check(err)

	go func() {
		f.WriteString(<-Log)
		l := <- Log 
		_check(err)
		fmt.Printf("%s (Logged) \n", l)
		f.Sync()
	}()
}

func _check(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jameycribbs/echgo/echo"
)

func main() {
	echoer := echo.NewEchoer()
	fmt.Println(echoer.Echo(strings.Join(os.Args[1:], " ")))
}

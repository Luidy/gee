//go:generate swagger generate spec
package main

import (
	"fmt"
	"gomod/config"
	"gomod/util"
	"runtime"
)

const (
	banner = "\n" +
		"		('-.     ('-.               \n" +
		"		_(  OO)  _(  OO)            \n" +
		"	,----.    (,------.(,------.    \n" +
		"	'  .-./-')  |  .---' |  .---'   \n" +
		"	|  |_( O- ) |  |     |  |       \n" +
		"	|  | .--, /(|  '--. (|  '--.    \n" +
		"	(|  | '. (_/ |  .--'  |  .--'   \n" +
		"	|  '--'  |  |  `---. |  `---.   \n" +
		"	`------'   `------' `------'    \n"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("server CPU core: %d\n", runtime.NumCPU())
}

func main() {
	Gee := config.Gee
	_, _ = util.InitLog(Gee.GetString("loglevel"))

	fmt.Printf(banner)
	fmt.Println("Hello... GO World!")
}

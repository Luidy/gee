//go:generate swagger generate spec
package main

import "fmt"

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

func main() {
	fmt.Printf(banner)
	fmt.Println("Hello... GO World!")
}

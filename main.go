//go:generate swagger generate spec
package main

import (
	"fmt"
	"gomod/config"
	"gomod/util"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func echoInit(gee *config.ViperConfig) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	return e
}

//go:generate swagger generate spec
package main

import (
	"fmt"
	"gee/repository"
	"gomod/config"
	"log"
	"net/http"
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
	e := echoInit(Gee)
	repository.InitDB(Gee)
	startServer(Gee, e)
}

func echoInit(gee *config.ViperConfig) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	healthCheck := func(c echo.Context) error {
		return c.String(http.StatusOK, "Gee Alive.")
	}
	e.GET("/healthCheck", healthCheck)

	e.HideBanner = true
	return e
}

func startServer(gee *config.ViperConfig, e *echo.Echo) {
	address := fmt.Sprintf("0.0.0.0:%d", gee.GetInt("port"))
	fmt.Println(banner, address)
	if err := e.Start(address); err != nil {
		log.Print("End echo server", "err", err)
	}
}

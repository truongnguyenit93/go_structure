package main


import (
	"log"
	"os"
	"blog/script"

	"github.com/samber/do"
	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
)

func args(injector *do.Injector) bool {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		flag := script.Commands(injector)

		return flag
	}
	
	return false
}

func run(server *gin.Engine) {
	server.Static("/assets", "./assets")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "0.0.0.0:" + port
	} else {
		serve = ":" + port
	}

	myFigure := figure.NewColorFigure("Go Structure", "", "green", true)
	myFigure.Print()

	if err := server.Run(serve); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func main() {
	var (
		injector = do.New()
		server   = gin.New()
	)

	providers.ResgisterDependencies(injector)

	if !args(injector) {
		return
	}

	run(server)
}
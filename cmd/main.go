package main


import (
	"log"
	"os"
	"blog/script"
	"blog/providers"
	"blog/middlewares"
	"blog/modules/user"
	"blog/modules/auth"

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

	serve := "0.0.0.0:" + port
	if os.Getenv("APP_ENV") != "localhost" {
		serve = ":" + port
	}

	figure.NewColorFigure("Go Structure", "", "green", true).Print()

	if err := server.Run(serve); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func main() {
	injector := do.New()
	providers.RegisterDependencies(injector)

	if !args(injector) {
		return
	}

	server := gin.New()
	server.Use(middlewares.CORSMiddleware())

	user.RegisterRoutes(server, injector)
	auth.RegisterRoutes(server, injector)

	run(server)
}
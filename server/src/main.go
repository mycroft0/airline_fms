package main

import (
	"fmt"
	_ "fmt"
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"os"
	_ "os"
	"server/src/controllers"
	"server/src/initializers"
)

func init() {
	initializers.LoadEnv()
}

func main() {

	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(cors.Default())

	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Use(func(c *gin.Context) {
		dataset := c.Param("dataset")
		initializers.DBconn(dataset)

	})

	dataset := r.Group("/:dataset")
	dataset.GET("/", controllers.ListAirline)

	// Configure app address listen on localhost only
	addr := fmt.Sprintf("127.0.0.1:%s", port)
	fmt.Printf("Application started on %s \n", addr)
	r.Run(addr)

}

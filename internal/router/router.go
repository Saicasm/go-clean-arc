package router

import (
	"github.com/clean-arc/internal/factories"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.Use(gin.Recovery())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Replace with your allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))
	loadHttpController := factories.MakeLoadHttpController()
	v1 := r.Group("/api/v1")
	{
		load := v1.Group("/load")
		{
			//TODO: add controller layer here
			load.GET("/health", loadHttpController.GetServerHealth)
			// Define other routes`
		}
	}
}

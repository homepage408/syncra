package bootstrap

import (
	"github.com/gin-gonic/gin"
)

func SetupMiddlewaresGin(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// Add CORS if needed:
	// r.Use(func(c *gin.Context){ ... })
}

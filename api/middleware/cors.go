package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pwsnet/golang-api/application"
	"github.com/pwsnet/golang-api/lib"
)

func CorsMiddleware(application *application.Application) gin.HandlerFunc {

	allowAllOrigins := false
	origins := []string{}
	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	headers := []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	exposeHeaders := []string{"Content-Length"}
	maxAge := 12 * time.Hour

	switch application.Environment {
	case lib.PROD:
		origins = []string{
			// NOTE: Websites
			"https://www.pomaretta.com",
			// NOTE: APIs
			"https://api.pomaretta.com",
		}
	case lib.DEV:
		origins = []string{
			// NOTE: Websites
			"https://www.pomaretta.com",
			// NOTE: APIs
			"https://api.dev.pomaretta.com",
		}
	}

	// TODO: Check if it's a local environment.
	// if ${LOCAL_ENVIRONMENT} {
	// allowAllOrigins = true
	// }

	return cors.New(cors.Config{
		AllowAllOrigins: allowAllOrigins,
		AllowOrigins:    origins,
		AllowMethods:    methods,
		AllowHeaders:    headers,
		ExposeHeaders:   exposeHeaders,
		MaxAge:          maxAge,
	})
}

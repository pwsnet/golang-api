package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pwsnet/golang-api/application"
	"github.com/pwsnet/golang-api/lib"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"golang.org/x/net/webdav"
)

func SwaggerMiddleware(application *application.Application, spec *swag.Spec, confs ...func(c *ginSwagger.Config)) gin.HandlerFunc {

	spec.Schemes = []string{"https"}
	spec.Version = application.Version

	switch application.Environment {
	case lib.PROD:
		// NOTE: This is a production environment.
		spec.Host = "api.pomaretta.com"
	case lib.DEV:
		// NOTE: This is a development environment.
		spec.Host = "api.dev.pomaretta.com"
	}

	// TODO: Check if it's a local environment.
	// if ${LOCAL_ENVIRONMENT} {
	// 	spec.Schemes = []string{"http"}
	// 	spec.Host = "localhost:3000"
	// }

	return ginSwagger.WrapHandler(&webdav.Handler{
		FileSystem: swaggerFiles.FS,
		LockSystem: webdav.NewMemLS(),
	}, confs...)
}

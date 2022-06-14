package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pwsnet/golang-api/application"
	"github.com/pwsnet/golang-api/lib"
)

type API struct {
	handler *gin.Engine

	application *application.Application
}

func New(application *application.Application) *API {
	api := &API{
		handler:     gin.Default(),
		application: application,
	}

	mode := gin.DebugMode
	if application.Environment == lib.PROD {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	api.initMiddlewares()
	api.initRoutes()

	return api
}

func (api *API) initRoutes() {

	api.initPublic()

}

func (api *API) initMiddlewares() {

	// NOTE: CORS
	// api.handler.Use(middleware.CorsMiddleware(api.application))

	// NOTE: Insensitive
	// api.handler.Use(middleware.InsensitiveMiddleware)

}

func (api *API) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	api.handler.ServeHTTP(w, req)
}

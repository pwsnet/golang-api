package api

import "github.com/gin-gonic/gin"

func (api *API) initPublic() {

	r := api.handler

	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}

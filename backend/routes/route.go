package routes

import (
	"APIEmail/backend/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigurarRotas(r *gin.Engine) {

	r.GET("/alerta", func(c *gin.Context) {
		controllers.AlertaHandler(c)
	})

	r.GET("/requisicao", func(c *gin.Context) {
		controllers.RequisicaoHandler(c)
	})

}

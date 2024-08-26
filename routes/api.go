package routes

import (
	"app/modules"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.RouterGroup, mod *modules.Modules) {

	product := r.Group("/product")
	{
		product.POST("/create", mod.Product.Ctl.Create)
		product.PATCH("/:id", mod.Product.Ctl.Update)
		product.DELETE("/:id", mod.Product.Ctl.Delete)
		product.GET("/:id", mod.Product.Ctl.Get)
		product.GET("/list", mod.Product.Ctl.List)
	}

}

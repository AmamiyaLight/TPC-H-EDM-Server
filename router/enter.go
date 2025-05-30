package router

import (
	"TPC-EDM-Server/global"
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(global.Config.System.GinMode)
	r := gin.Default()

	r.Static("/uploads", "uploads")

	nr := r.Group("/api")
	addr := global.Config.System.Addr()
	UserRouter(nr)
	OrdersRouter(nr)
	PartSuppRouter(nr)
	LineItemRouter(nr)
	NationRouter(nr)
	CustomerRouter(nr)
	PartRouter(nr)
	RegionRouter(nr)
	SupplierRouter(nr)
	r.Run(addr)
}

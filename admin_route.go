package main

import (
	"food_delivery/component/appctx"
	"food_delivery/middleware"
	"food_delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func setupAdminRoute(appCtx appctx.AppContext, r *gin.RouterGroup) {
	admin := r.Group("/admin",
		middleware.AuthenticateJWT(appCtx),
		middleware.RequiredRoles(appCtx, "admin"),
	)

	{
		admin.GET("/profile", ginuser.FindUser(appCtx))
	}
}

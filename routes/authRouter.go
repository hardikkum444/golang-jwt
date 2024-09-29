package routes

import (
    
    controller "golang-jwt/controllers"
    "github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine){

    incomingRoutes.POST("users/signup", controller.Signup())
    incomingRoutes.POST("users/login", controller.Login())

}

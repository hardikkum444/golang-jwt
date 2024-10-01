package routes

import (
    controller "golang-jwt/controllers"
    "golang-jwt/middleware"
    "github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
    
    // we shall be using a middleware to ensure that for the user routes the user has a verified token (not required in authroutes as that is where the user will get the token from)
    incomingRoutes.Use(middleware.Authenticate())
    incomingRoutes.GET("/users", controller.GetUsers())
    incomingRoutes.GET("/user/:user_id", controller.GetUser())

}






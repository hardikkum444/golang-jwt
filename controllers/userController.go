package controllers

import (

"context"
"fmt"
"log"
"strconv"
"net/http"
"time"
"github.com/gin-gonic/gin"
"github.com/go-playground/validator/v10"
helper "golang-jwt/helpers"
"golang-jwt/models"
"golang-jwt/helpers"
"golang.org/x/crypto/bcrypt"
"go.mongodb.org/mongo-driver/bson"

)

var userCollection *mongoCollection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup() gin.HandlerFunc {
    
    return func(c *gin.Context){
        
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        var user models.User

        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        validationError := validate.Struct(user)
        if validationError != nil{
            c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
            return
        }

        count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
        defer cancel()
        if err != nil {
            log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error":"error occurred while checking for email validity"})
        }

        count, err = userCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})
        defer cancel()
        if err != nil {
            log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error":"error occurred while checking for phone number validity"})
        }

        if count > 0{
            c.JSON(http.StatusInternalServerError, gin.H{"error":"email or phone number already exists"})
        }
    }
}

func Login()

func getUsers()

func getUser() gin.HandlerFunc {
    
    return func(c *gin.Context){
        
        userId := c.Param("user_id")

        if err := helper.MatchUserTypeToUid(c, userId); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
            return 
        }
        
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

        var user models.User
        err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)

        defer cancel()
        if err != nil{
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }

        c.JSON(http.StatusOK, user)
    }
}










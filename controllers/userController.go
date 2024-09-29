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

)

var userCollection *mongoCollection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func getUsers()

func getUser()
package controllers

import (
	"go-jwt/intializers"
	"go-jwt/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Get the User details (email/pass/name from the req body)

	var body models.User
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//  Hash the Password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash Password",
		})
		return
	}
	// Create the User
	user := models.User{Email: body.Email, Name: body.Name, Password: string(hash)}
	result := intializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}
	// Respond
	c.JSON(http.StatusOK, result)
}

func SignIn(c *gin.Context) {
	// get the email n pass from the req body
	var body struct {
		Email    string
		Password string
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// lookup user details in db
	var user models.User
	intializers.DB.First(&user, "email=?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email  and Password",
		})
		return
	}

	// verify password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password Not Matched",
		})
	}

	// Generate JWT Token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to Create the Token",
		})
	}
	// Pass the Generated Token as Cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

}
func Validate(c *gin.Context) {
	// Get the user from the middleware
	user, _ := c.Get("user")
	//  return the required user details on Validation
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.(models.User).ID,
			"email": user.(models.User).Email,
			"name":  user.(models.User).Name,
		}})
}

package main

import (
	"errors"
	"github.com/MicahParks/keyfunc"
	codes "github.com/dkadio/dysr/internal"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strings"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type SignedResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	log.Println("try parsing token:", jwtToken)
	jwksURL := "http://localhost:8080/auth/realms/dysr/protocol/openid-connect/certs"
	jwks, err := keyfunc.Get(jwksURL)
	token, err := jwt.Parse(jwtToken, jwks.Keyfunc)

	if err != nil {
		log.Println("error", err)
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func jwtTokenCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}
	c.Next()
}

func privateACLCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}

	claimedUID, OK := claims["user"].(string)
	if !OK {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "no user property in claims",
		})
		return
	}

	uid := c.Param("uid")
	if claimedUID != uid {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "token uid does not match resource uid",
		})
		return
	}

	c.Next()
}

func main() {

	cc := codes.NewCodesController()

	log.Println("Starting..")
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.Use(jwtTokenCheck)
	{
		v1.POST("/codes", cc.CeateCode)
		v1.GET("/codes", cc.GetCodes)
		v1.GET("/codes/{id}", cc.GetCode)
		v1.PUT("/codes/{id}", cc.UpdateCode)
	}

	log.Println(r.Run(":8000"))
	log.Println("End..")

}

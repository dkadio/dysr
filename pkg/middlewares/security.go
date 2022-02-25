package middlewares

import (
	"errors"

	"github.com/MicahParks/keyfunc"
	"github.com/dkadio/dysr/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"log"
	"net/http"
	"strings"
)

type Middlewares struct {
	jwksUrl string
}

func NewMiddlewares(jwksUrl string) Middlewares {

	config := util.LoadConfig()
	if jwksUrl == "" {
		return Middlewares{config.JWKSUrl}
	}
	return Middlewares{jwksUrl}
}

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

func (mw Middlewares) parseToken(jwtToken string) (*jwt.Token, error) {
	jwks, err := keyfunc.Get(mw.jwksUrl)
	token, err := jwt.Parse(jwtToken, jwks.Keyfunc)

	if err != nil {
		log.Println("error", err)
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func (mw Middlewares) None(c *gin.Context) {
	c.Next()
}

func (mw Middlewares) JwtTokenCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := mw.parseToken(jwtToken)
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
	c.Set("claims", claims)
	c.Next()
}

func (mw Middlewares) PrivateACLCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := mw.parseToken(jwtToken)
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

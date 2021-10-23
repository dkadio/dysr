package routes

import (
	"fmt"
	"github.com/dkadio/dysr/internal/controllers"
	"github.com/dkadio/dysr/internal/middlewares"
	"github.com/dkadio/dysr/internal/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"net/http"
)

// NewRouter creates all the routes/endpoints, using Fizz.
func NewCodesRouter() (*fizz.Fizz, error) {
	engine := gin.New()

	engine.Use(cors.Default())

	fizzApp := fizz.NewFromEngine(engine)

	infos := &openapi.Info{
		Title:       "Dysr Codes Api",
		Description: "Provides Codes Information.",
		Version:     "1.0.0",
	}

	cc := controllers.NewCodesController()
	fizzApp.GET("/docs", nil, fizzApp.OpenAPI(infos, "json"))

	group := fizzApp.Group("/api/v1", "Codes Api", "Provides all codes infos.")

	mw := middlewares.NewMiddlewares("")
	group.Use(mw.JwtTokenCheck)

	group.GET("/healthcheck", []fizz.OperationOption{
		fizz.Summary("Checks API is healthy."),
		fizz.Response(fmt.Sprint(http.StatusInternalServerError), "Server Error", models.APIError{}, nil, nil),
	}, tonic.Handler(controllers.Healthcheck, http.StatusOK))

	group.GET("/codes", []fizz.OperationOption{
		fizz.Summary("Get all User Codes."),
		fizz.Response(fmt.Sprint(http.StatusInternalServerError), "Server Error", models.APIError{}, nil, nil),
		fizz.Response(fmt.Sprint(http.StatusNotFound), "No Codes Found", models.APIError{}, nil, nil),
	}, tonic.Handler(cc.GetCodes, http.StatusOK))

	group.PUT("/pets:name", []fizz.OperationOption{
		fizz.Summary("Update a pet."),
		fizz.Response(fmt.Sprint(http.StatusInternalServerError), "Server Error", models.APIError{}, nil, nil),
	}, tonic.Handler(cc.CeateCode, http.StatusOK))

	if len(fizzApp.Errors()) != 0 {
		return nil, fmt.Errorf("fizz errors: %v", fizzApp.Errors())
	}
	tonic.SetErrorHook(errHook)
	return fizzApp, nil
}

func errHook(_ *gin.Context, e error) (int, interface{}) {
	code, msg := http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)

	if _, ok := e.(tonic.BindError); ok {
		code, msg = http.StatusBadRequest, e.Error()
	} else {
		switch {
		case errors.IsBadRequest(e), errors.IsNotValid(e), errors.IsNotSupported(e), errors.IsNotProvisioned(e):
			code, msg = http.StatusBadRequest, e.Error()
		case errors.IsForbidden(e):
			code, msg = http.StatusForbidden, e.Error()
		case errors.IsMethodNotAllowed(e):
			code, msg = http.StatusMethodNotAllowed, e.Error()
		case errors.IsNotFound(e), errors.IsUserNotFound(e):
			code, msg = http.StatusNotFound, e.Error()
		case errors.IsUnauthorized(e):
			code, msg = http.StatusUnauthorized, e.Error()
		case errors.IsAlreadyExists(e):
			code, msg = http.StatusConflict, e.Error()
		case errors.IsNotImplemented(e):
			code, msg = http.StatusNotImplemented, e.Error()
		}
	}
	err := models.APIError{
		Message: msg,
	}
	return code, err
}

package http

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"

	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

func New() (*fizz.Fizz, error) {
	engine := gin.New()
	engine.Use(cors.Default())
	engine.Use(gin.Recovery())

	middleware := Middleware{}

	engine.Use(middleware.Logging())
	engine.Use(middleware.SharedHeaders())

	fizz := fizz.NewFromEngine(engine)

	infos := &openapi.Info{
		Title:       "Fruits Market",
		Description: `This is a sample Fruits market server.`,
		Version:     "1.0.0",
	}
	// Create a new route that serve the OpenAPI spec.
	fizz.GET("/openapi.json", nil, fizz.OpenAPI(infos, "json"))

	applyInternalRoutes(fizz.Group("/internal", "Internal", "Routes used for internal or infrastructure reasons"))
	applyUserRoutes(fizz.Group("/users", "Users", "Routes for interacting with users"))

	if len(fizz.Errors()) != 0 {
		return nil, fmt.Errorf("fizz errors: %v", fizz.Errors())
	}

	tonic.SetErrorHook(ErrorHook)

	return fizz, nil
}

func applyInternalRoutes(group *fizz.RouterGroup) {
	group.GET("/healthcheck", []fizz.OperationOption{
		fizz.Summary("Returns if the system is healthy or not"),
		fizz.Response("500", "Service Not Healthy", nil, nil, map[string]interface{}{
			"notHealthy": map[string]interface{}{"error": "not healthy"},
		}),
		fizz.ID("Healtcheck"),
	}, tonic.Handler(HTTPRoutes.Healthcheck(), 200))
}

func applyUserRoutes(group *fizz.RouterGroup) {
	group.POST("/", []fizz.OperationOption{
		fizz.Summary("Creates a new user"),
		fizz.ID("CreateUser"),
	}, tonic.Handler(HTTPRoutes.CreateUser, 201))

	group.POST("/login", []fizz.OperationOption{
		fizz.Summary("Generates a token to be used for future requests"),
		fizz.ID("GetUserToken"),
	}, tonic.Handler(HTTPRoutes.GetUserToken, 200))

	group.GET("/:id", []fizz.OperationOption{
		fizz.Summary("Gets user by ID"),
		fizz.ID("GetUserByID"),
	}, tonic.Handler(HTTPRoutes.GetUserById, 200))
}

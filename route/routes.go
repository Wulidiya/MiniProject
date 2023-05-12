package route

import (
	"MiniProject/controller"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", controller.GetUserController)
	user.GET("/:id", controller.GetUserController)
	user.POST("", controller.CreateUserController)
	user.PUT("/:id", controller.UpdateUserController)
	user.DELETE("/:id", controller.DeleteUserController)

	// transaction collection
	transaction := e.Group("/transactions")
	transaction.GET("", controller.GetTransactionController)
	transaction.GET("/:id", controller.GetTransactionController)
	transaction.POST("", controller.CreateTransactionController)
	transaction.PUT("/:id", controller.UpdateTransactionController)
	transaction.DELETE("/:id", controller.DeleteTransactionController)

	// tool collection
	tool := e.Group("/tools")
	tool.GET("", controller.GetToolController)
	tool.GET("/:id", controller.GetToolController)
	tool.POST("", controller.CreateToolController)
	tool.PUT("/:id", controller.UpdateToolController)
	tool.DELETE("/:id", controller.DeleteToolController)

	// returns collection
	returns := e.Group("/tools")
	returns.GET("", controller.GetReturnsController)
	returns.GET("/:id", controller.GetReturnsController)
	returns.POST("", controller.CreateReturnsController)
	returns.PUT("/:id", controller.UpdateReturnsController)
	returns.DELETE("/:id", controller.DeleteReturnsController)
}

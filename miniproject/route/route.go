package route

import (
	"net/http"
	"miniproject/constant"
	"miniproject/controller"
	"miniproject/repository/database"
	"miniproject/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	userRepository := database.NewUserRepository(db)
	blogRepository := database.NewBlogRepository(db)
	bookRepository := database.NewBookRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, blogRepository)
	blogUsecase := usecase.NewBlogUsecase(blogRepository)
	bookUsecase := usecase.NewBookUsecase(bookRepository)

	authController := controller.NewAuthController(userUsecase)
	userController := controller.NewUserController(userUsecase, userRepository)
	blogController := controller.NewBlogController(blogUsecase, blogRepository)
	bookController := controller.NewBookController(bookUsecase, bookRepository)

	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", authController.LoginUserController)

	// user collection
	user := e.Group("/users", middleware.JWT([]byte(constant.SECRET_JWT)))
	user.GET("", userController.GetUserController)
	user.GET("/:id", userController.GetUserController)
	user.POST("", userController.CreateUserController)
	user.PUT("/:id", userController.UpdateUserController)
	user.DELETE("/:id", userController.DeleteUserController)

	// tool collection
	tool := e.Group("/tools", middleware.JWT([]byte(constant.SECRET_JWT)))
	tool.GET("", toolController.GetToolController)
	tool.GET("/:id", toolController.GetToolController)
	tool.POST("", toolController.CreateToolController)
	tool.PUT("/:id", toolController.UpdateToolController)
	tool.DELETE("/:id", toolController.DeleteToolController)

	//route tool detail
	tooldetail := e.Group("/tooldetail", middleware.JWT([]byte(constant.SECRET_JWT)))
	tooldetail.GET("", tooldetailcontroller.GettoolDetailController)
	tooldetail.GET("/:id", tooldetailcontroller.GettoolDetailController)
	tooldetail.POST("", tooldetailcontroller.CreatetoolDetailController)
	tooldetail.PUT("/:id", tooldetailcontroller.UpdatetoolDetailController)
	tooldetail.DELETE("/:id", tooldetailcontroller.DeletetoolDetailController)

	//route return
	return := e.Group("/tool_return", middleware.JWT([]byte(constant.SECRET_JWT)))
	toolreturn.GET("", toolreturncontroller.GetReturnController)
	toolreturn.GET("/:id", toolreturncontroller.GetReturnController)
	toolreturn.POST("", toolreturncontroller.CreateReturnController)
	toolreturn.PUT("/:id", toolreturncontroller.UpdateReturnController)
	toolreturn.DELETE("/:id", toolreturncontroller.DeleteReturnController)

	//route transaction
	transaction := e.Group("/transactions", middleware.JWT([]byte(constant.SECRET_JWT)))
	transaction.GET("", controller.GetTransactionController)
	transaction.GET("/:id", controller.GetTransactionController)
	transaction.POST("", controller.CreateTransactionController)
	transaction.PUT("/:id", controller.UpdateTransactionController)
	transaction.DELETE("/:id", controller.DeleteTransactionController)

}
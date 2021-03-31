package routes

import (
	"os"
	"retailStore/controllers"
	"retailStore/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	middlewares.LogMiddlewares(e)
	
	e.POST("/register", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUserController)

	
	e.GET("/items", controllers.GetItemController)
	e.GET("/items/:id", controllers.GetItemWIthParamsController)
	
	e.GET("/shoppingcarts", controllers.GetShoppingCartController)
	e.POST("/shoppingcarts", controllers.PostItemToShoppingCartController)

	e.GET("/couriers", controllers.GetCouriersController)
	e.GET("/couriers/:id", controllers.GetCourierByIdController)
	e.DELETE("/couriers/:id", controllers.DeleteCourierByIdController)
	e.PUT("/couriers/:id", controllers.UpdateCourierByIdController)
	e.POST("/couriers", controllers.CreateCourierController)


	eJWT := e.Group("") 
	eJWT.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJWT.POST("/address", controllers.CreateAddressController)
	eJWT.GET("/address", controllers.GetAddressController)
	eJWT.GET("/address/:id", controllers.GetAddressByIdController)

	eJWT.GET("/users", controllers.GetUserDetailController)
	eJWT.PUT("/users", controllers. UpdateUserDetailController)

	eAdmin := eJWT.Group("")
	eAdmin.POST("/items", controllers.PostItemController)
	

	return e
}

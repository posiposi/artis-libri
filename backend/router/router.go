package router

import (
	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/controller"
)

func NewRouter(uc controller.IUserController, bc controller.IBookController) *echo.Echo {
	e := echo.New()
	g := e.Group("/v1")
	u := g.Group("/secret-admin")
	u.POST("/signup", uc.SignUp)
	u.POST("/login", uc.LogIn)
	u.POST("/logout", uc.LogOut)
	b := g.Group("/books")
	b.GET("", bc.GetAllBooks)
	b.GET("/:bookId", bc.GetBookByBookId)
	// TODO GET以外のメソッドはログインしているユーザーのみアクセス可能
	b.POST("", bc.CreateBook)
	b.PUT("/:bookId", bc.UpdateBook)
	b.DELETE("/:bookId", bc.DeleteBook)
	return e
}

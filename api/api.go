package api

import (
	"library-api/api/book"
	"library-api/api/comment"
	"library-api/api/middleware"
	"library-api/api/operate"
	"library-api/api/user"

	"github.com/gin-gonic/gin"
)

func InitGroup(engine *gin.Engine) {
	engine.POST("/register", middleware.MiddleWarePassJSON, user.Register)
	engine.GET("/user/token", middleware.MiddleWarePassJSON, user.Login)
	engine.GET("/user/token/refresh", middleware.MiddleWarePassJSON, user.RefreshToken)
	engine.PUT("/user/password", middleware.MiddleWareJWTVerify, middleware.MiddleWarePassJSON, user.ChangePassword)
	engine.PUT("/user/info", middleware.MiddleWareJWTVerify, middleware.MiddleWarePassJSON, user.ChangeUserInfo)
	engine.GET("/book/list", middleware.MiddleWareJWTMention, book.BookList)
	engine.GET("/book/search", middleware.MiddleWareJWTMention, middleware.MiddleWarePassJSON, book.BookSearchByName)
	engine.PUT("/book/focus", middleware.MiddleWareJWTVerify, middleware.MiddleWarePassJSON, book.Focus)
	engine.GET("/book/label", middleware.MiddleWareJWTMention, middleware.MiddleWarePassJSON, book.SearchBooksByLabel)
	engine.POST("/comment/book/:book_id", middleware.MiddleWareJWTVerify, middleware.MiddleWarePassJSON, comment.AddBookComment)
	engine.POST("/comment/comment/:comment_id", middleware.MiddleWareJWTVerify, middleware.MiddleWarePassJSON, comment.AddCommentComment)
	engine.PUT("/comment/:comment_id", middleware.MiddleWareJWTVerify, middleware.MiddleWarePassJSON, comment.UpdateComment)
	engine.DELETE("/comment/:comment_id", middleware.MiddleWareJWTVerify, comment.DeleteComment)
	engine.GET("/comment/book/:book_id", middleware.MiddleWareJWTMention, comment.ListCommentsUnderBook)
	engine.PUT("/operate/praise", middleware.MiddleWareJWTVerify, middleware.MiddleWarePassJSON, operate.Praise)
	engine.GET("/operate/collect/list", middleware.MiddleWareJWTVerify, middleware.MiddleWarePassJSON, operate.ListFocus)
}

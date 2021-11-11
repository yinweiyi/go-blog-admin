package routes

import (
	"blog/app/controllers/api"
	"blog/app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoute(engine *gin.Engine) {
	frontend := engine.Group("/api")
	frontend.Use(middlewares.Cors())
	{
		//公共页
		common := new(api.CommonController)
		frontend.GET("/config", common.Config)
		frontend.GET("/categories", common.Categories)
		frontend.GET("/random_tags", common.RandomTags)
		frontend.GET("/hots", common.Hots)
		frontend.GET("/newest_comments", common.NewestComments)
		frontend.GET("/friend_links", common.FriendLinks)
		frontend.GET("/sentence", common.Sentence)
		frontend.GET("/guestbook", common.Guestbook)
		frontend.GET("/abouts", common.Abouts)

		//文章页
		article := new(api.ArticleController)
		frontend.GET("/articles", article.List)
		frontend.GET("/articles/:slug", article.Show)

		//验证码
		captcha := new(api.CaptchaController)
		frontend.GET("/captcha/:captchaId", captcha.Captcha)
		frontend.GET("/captcha", captcha.Get)

		//评论
		comment := new(api.CommentController)
		frontend.GET("/comments", comment.Comments)
		frontend.Any("/comment", comment.Comment)
	}

}

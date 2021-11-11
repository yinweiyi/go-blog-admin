package api

import (
	"blog/app/models"
	"blog/app/services"
	"blog/vendors/model"
	configRedis "blog/vendors/redis/config"

	"github.com/gin-gonic/gin"
)

type CommonController struct {
	BaseController
}

func (c CommonController) Config(ctx *gin.Context) {
	config, err := configRedis.Get()
	c.FailOnError(ctx, err)

	c.Success(ctx, "获取成功", config)
}

func (c CommonController) Categories(ctx *gin.Context) {
	categories := new(services.CategoryService).GetAll()
	c.Success(ctx, "获取成功", categories)
}

func (c CommonController) RandomTags(ctx *gin.Context) {
	tags := models.Shuffle(new(services.TagService).MinTags())
	c.Success(ctx, "获取成功", tags)
}

func (c CommonController) Hots(ctx *gin.Context) {
	hots := new(services.ArticleService).Hots(10)
	c.Success(ctx, "获取成功", hots)
}

func (c CommonController) NewestComments(ctx *gin.Context) {
	comments := new(services.CommentService).News()
	c.Success(ctx, "获取成功", comments)
}

func (c CommonController) FriendLinks(ctx *gin.Context) {
	friendsLinks := new(services.FriendshipLinkService).Chuck(2)
	c.Success(ctx, "获取成功", friendsLinks)
}

func (c CommonController) Sentence(ctx *gin.Context) {
	sentence := new(services.SentenceService).GetOne()
	c.Success(ctx, "获取成功", sentence)
}

func (c CommonController) Guestbook(ctx *gin.Context) {
	var guestbook models.Guestbook
	model.DB.First(&guestbook)
	c.Success(ctx, "获取成功", guestbook)
}

func (c CommonController) Abouts(ctx *gin.Context) {
	abouts := new(services.AboutService).All()
	c.Success(ctx, "获取成功", abouts)
}

package api

import (
	"blog/app/models/forms"
	"blog/app/services"
	"blog/vendors/validate"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentParams struct {
	Captcha Captcha
	ID      uint
	Type    string
}

type CommentController struct {
	BaseController
}

func (c CommentController) Comment(ctx *gin.Context) {
	var form forms.Comment

	if err := ctx.ShouldBindJSON(&form); err != nil {
		c.FailOnError(ctx, err)
	}
	validator := validate.GetValidator()
	if err := validator.Struct(form); err != nil {
		c.Error(ctx, validate.TranslateOverride(err), nil)
		return
	}
	//验证码验证
	if !new(CaptchaController).Verify(form.CaptchaId, form.Captcha) {
		c.Error(ctx, "验证码错误", nil)
		return
	}
	form.IP = ctx.ClientIP()

	if comment, err := new(services.CommentService).Comment(form); err == nil {
		c.Success(ctx, "评论成功", gin.H{
			"comment_id": comment.ID,
		})
		return
	}
	c.Error(ctx, "评论失败", nil)
}

var TypeList = map[string]bool{
	"article":   true,
	"guestbook": true,
	"about":     true,
}

func (c CommentController) Comments(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(ctx, "请传入文件ID", nil)
		return
	}
	typeStr := ctx.Query("type")
	if _, ok := TypeList[typeStr]; !ok {
		c.Error(ctx, "要查询的类型错误", nil)
		return
	}

	//id type
	commentService := new(services.CommentService)
	commentTree, commentPageData := commentService.GetTree(ctx.Request, 5, id, typeStr)

	c.Success(ctx, "获取成功", map[string]interface{}{
		"commentCount": commentService.Count(uint(id), typeStr),
		"commentTree":  commentTree,
		"pageData":     commentPageData,
	})
}

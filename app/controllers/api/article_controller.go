package api

import (
	"blog/app/models"
	"blog/app/services"
	"blog/vendors/pagination"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (a ArticleController) List(ctx *gin.Context) {
	slug := ctx.Query("slug")
	name := ctx.Query("name")
	fmt.Println(slug)
	articleService := new(services.ArticleService)
	var articles []models.Article
	var pagerData = pagination.PagerData{}

	if name == "categories" {
		cate, err := new(services.CategoryService).GetBySlug(slug)
		if err == nil {
			articles, pagerData, _ = articleService.GetAll(ctx.Request, 5, map[string]interface{}{"category_id": cate.ID})
		}
	} else if name == "tags" {
		tagService := new(services.TagService)

		tag, err := tagService.GetBySlug(slug)
		if err == nil {
			articles, pagerData, _ = tagService.GetArticlesByTag(ctx.Request, tag, 5)
		}
	} else {
		var where map[string]interface{}
		articles, pagerData, _ = articleService.GetAll(ctx.Request, 5, where)
	}

	a.Success(ctx, "获取成功", map[string]interface{}{
		"articles":  articles,
		"pagerData": pagerData,
	})

}

func (a *ArticleController) Show(ctx *gin.Context) {
	slug := ctx.Param("slug")
	articleService := new(services.ArticleService)
	article, err := articleService.GetBySlug(slug)
	a.FailOnError(ctx, err)

	articleService.Read(article)
	article.Views += 1

	a.Success(ctx, "获取成功", gin.H{
		"article": article,
		"last":    articleService.Last(article),
		"next":    articleService.Next(article),
	})
}

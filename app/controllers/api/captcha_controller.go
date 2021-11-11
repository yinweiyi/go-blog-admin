package api

import (
	"blog/vendors/config"
	"bytes"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dchest/captcha"
)

type CaptchaController struct {
	BaseController
}

type Captcha struct {
	ID       string `json:"captcha_id"` //验证码Id
	ImageUrl string `json:"image_url"`  //验证码图片url
}

//Get 获取captcha id
func (c CaptchaController) Get(ctx *gin.Context) {
	captchaId := captcha.NewLen(config.GetInt("captcha.length"))

	var captchaModel Captcha
	captchaModel.ID = captchaId
	captchaModel.ImageUrl = "/captcha/" + captchaId + ".png"

	c.Success(ctx, "获取成功", captchaModel)
}

//Captcha 获取图片
func (CaptchaController) Captcha(ctx *gin.Context) {
	captchaId := ctx.Param("captchaId")
	fmt.Println("GetCaptchaPng : " + captchaId)
	ServeHTTP(ctx.Writer, ctx.Request)
}
func (CaptchaController) Verify(captchaId, value string) bool {
	if captchaId == "" || value == "" {
		return false
	}
	return captcha.VerifyString(captchaId, value)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, config.GetInt("captcha.width"), config.GetInt("captcha.height")) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil

}

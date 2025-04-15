package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/model/response"
)

type BaseApi struct{}

var store = base64Captcha.DefaultMemStore

func (baseApi *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(
		global.Config.Captcha.Height,
		global.Config.Captcha.Width,
		global.Config.Captcha.Length,
		global.Config.Captcha.MaxSkew,
		global.Config.Captcha.DotCount,
	)

	captcha := base64Captcha.NewCaptcha(driver, store)

	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.Log.Error("验证码获取失败! ", zap.Error(err))
		response.FailWithMessage("验证码获取失败!", c)
		return
	}
	response.OkWithData(response.Captcha{
		CaptchaID: id,
		PicPath:   b64s,
	}, c)
}

func (baseApi *BaseApi) SendEmailVerificationCode(c *gin.Context) {
	var req request.SendEmailVerificationCode
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if !store.Verify(req.CaptchaID, req.Captcha, true) {
		response.FailWithMessage("验证码错误", c)
		return
	}
	err := baseService.SendEmailVerificationCode(c, req.Email)
	if err != nil {
		global.Log.Error("发送邮箱验证码失败!", zap.Error(err))
		response.FailWithMessage("发送邮箱验证码失败!", c)
		return
	}
	response.OkWithMessage("成功发送邮箱验证码!", c)
}

func (baseApi *BaseApi) QQLoginURL(c *gin.Context) {
	url := global.Config.QQ.QQLoginURL()
	response.OkWithData(url, c)
}

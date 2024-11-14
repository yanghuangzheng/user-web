package model

import(
    "net/http"
	"context"
	"time"

	"user-web/global"
    
	"github.com/mojocn/base64Captcha"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
var store=base64Captcha.DefaultMemStore

func GetCaptcha(ctx *gin.Context){
	mobile := ctx.Query("mobile")
	if mobile == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号码不能为空",
		})
		return
	}
	driver:=base64Captcha.NewDriverDigit(240,80,5,0.7,80)//高 宽 长
	cp:=base64Captcha.NewCaptcha(driver,store)
	id,b64s,rawCaptcha,err:=cp.Generate()
	if err!=nil{
		zap.S().Errorf("生成验证码错误,:",err.Error())
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"msg":"生成验证码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64s,
		"rawCaptcha": rawCaptcha, // 可选：如果你不需要返回原始验证码数据，可以去掉这一行
	})
	//用redis缓存
	redisCtx := context.Background()
	redisKey := "captcha:" + id
	err = global.Rdb.Set(redisCtx, redisKey, rawCaptcha, time.Duration(500)*time.Second).Err()
	if err != nil {
		zap.S().Errorf("存储验证码到 Redis 错误: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "存储验证码错误",
		})
		return
	}

}


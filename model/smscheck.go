package model

import (
	"context"
	//"fmt"
	"net/http"
	//"time"

	"user-web/global"
	"user-web/form"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/go-redis/redis/v8"
)

// VerifySmsCode 验证短信验证码
func VerifySmsCode(ctx *gin.Context) {
	verifyForm := form.VerifySmsCodeForm{}
	if err := ctx.ShouldBind(&verifyForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	mobile := verifyForm.Mobile
	userInputCode := verifyForm.Code

	// 从 Redis 中获取存储的验证码
	storedCode, err := global.Rdb.Get(context.Background(), mobile).Result()
	if err == redis.Nil {
		zap.S().Info("验证码不存在或已过期")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码不存在或已过期",
		})
		return
	} else if err != nil {
		zap.S().Error("从 Redis 获取验证码时发生错误:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	// 比对验证码
	if userInputCode != storedCode {
		zap.S().Info("验证码不匹配")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码不匹配",
		})
		return
	}

	// 验证码匹配，继续处理注册或其他操作
	zap.S().Info("验证码匹配成功")
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "验证码匹配成功",
	})

	// 清除已使用的验证码
	err = global.Rdb.Del(context.Background(), mobile).Err()
	if err != nil {
		zap.S().Error("删除 Redis 中的验证码时发生错误:", err)
	}
}
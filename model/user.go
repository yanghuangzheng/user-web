package model

import (
	"context"
	"net/http"
	"time"

	"user-web/form"
	"user-web/global"
	"user-web/proto"
	"user-web/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

func GetUserList(ctx *gin.Context) {
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    10,
		Psize: 10,
	})
	if err != nil {
		zap.S().Info("userSrvClient.GetUserList 查询用户列表失败")
		HandleGrpcErrortoHttp(err, ctx)
		return
	}
	// 准备返回的结果
	var result []form.UserRespons
	for _, value := range rsp.Data {
		birthday:=time.Unix(int64(value.Birthday), 0)
		user := form.UserRespons{
			Id:       value.Id,
			Name:     value.Name,
			Birthday:  &birthday,
			Gender:   value.Gender,
			Mobile:   value.Moblie,
			Role:     value.Role,
		}
		result = append(result, user)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}
func MobileInfo(ctx *gin.Context){
	zap.S().Info("【mobileinfo 开始调用】")
	var Mobile form.MobileRequest
	if err := ctx.ShouldBindJSON(&Mobile); err != nil {
        // 处理绑定错误
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	rsp, err := global.UserSrvClient.SelectByMobile(context.Background(), &proto.MobileInfo{
	  Mobile:Mobile. Mobile,
	})
	if err != nil {
		zap.S().Info("userSrvClient.SelectByMobile 【手机查询用户失败】")
		HandleGrpcErrortoHttp(err, ctx)
		return
	}
	birthday:=time.Unix(int64(rsp.Birthday), 0)
	user := form.UserRespons{
		Id:       rsp.Id,
		Name:     rsp.Name,
		Birthday: &birthday,
		Gender:   rsp.Gender,
		Mobile:   rsp.Moblie,
		Role:     rsp.Role,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
func CreateUser(ctx *gin.Context){
	zap.S().Info("【CreateUser 开始调用】")
	var CreateUser form. UserInfo
	if err := ctx.ShouldBindJSON(&CreateUser); err != nil {
        // 处理绑定错误
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	rsp, err := global.UserSrvClient.CreateUser(context.Background(), &proto.UserInfo{
		Password:   CreateUser.PassWord,
	    Moblie:     CreateUser.Mobile,
		Name:       CreateUser.Name,
		Birthday:   uint64(CreateUser.Birthday.Unix()),
		Gender:     CreateUser.Gender,
	})
	if err != nil {
		zap.S().Info("userSrvClient.CreateUser 【创建用户失败】")
		HandleGrpcErrortoHttp(err, ctx)
		return
	}
	if rsp.Success=="yes"{
	ctx.JSON(http.StatusOK, gin.H{
		"data": "ok",
	})}else{
		ctx.JSON(http.StatusOK, gin.H{
			"data": "no",
		})	
}
}
func SelectById(ctx *gin.Context){
	zap.S().Info(" SelectById 开始调用】")
	var id form.IdRequest
	if err := ctx.ShouldBindJSON(&id); err != nil {
        // 处理绑定错误
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	rsp, err := global.UserSrvClient.SelectById(context.Background(), &proto.IdInfo{
	  Id:id.ID,
	})
	if err != nil {
		zap.S().Info("userSrvClient. SelectById 【id查询用户失败】")
		HandleGrpcErrortoHttp(err, ctx)
		return
	}
	birthday:=time.Unix(int64(rsp.Birthday), 0)
	user := form.UserRespons{
		Id:       rsp.Id,
		Name:     rsp.Name,
		Birthday: &birthday,
		Gender:   rsp.Gender,
		Mobile:   rsp.Moblie,
		Role:     rsp.Role,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
func UpdateUser(ctx *gin.Context){
	zap.S().Info("【UpdateUser 开始调用】")
	userId := ctx.MustGet("userId").(string)
	var CreateUser form. UserInfo
	if err := ctx.ShouldBindJSON(&CreateUser); err != nil {
        // 处理绑定错误
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	rsp, err := global.UserSrvClient.UpdateUser(context.Background(), &proto.UserInfo{
		Id      :   userId,
		Password:   CreateUser.PassWord,
	    Moblie:     CreateUser.Mobile,
		Name:       CreateUser.Name, 
		Birthday:   uint64(CreateUser.Birthday.Unix()),
		Gender:     CreateUser.Gender,
	})
	if err != nil {
		zap.S().Info("userSrvClient.UpdateUser 【更新用户失败】")
		HandleGrpcErrortoHttp(err, ctx)
		return
	}
	if rsp.Success=="yes"{
	ctx.JSON(http.StatusOK, gin.H{
		"data": "ok",
	})}else{
		ctx.JSON(http.StatusOK, gin.H{
			"data": "no",
		})	}
}
func Logger(ctx *gin.Context){
	zap.S().Info("【UpdateUser 开始调用】")
	var Password form.PasswordInfo
	if err := ctx.ShouldBindJSON(&Password); err != nil {
        // 处理绑定错误
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	rsp, err := global.UserSrvClient.Logger(context.Background(), &proto.PasswordInfo{
		 Password: Password.Password,
		 Id:       Password.Id,
	})
	if err != nil {
		zap.S().Info("userSrvClient.Loggerr 【登录用户失败】")
		HandleGrpcErrortoHttp(err, ctx)
		return
	}
	if rsp.Success=="yes"{
	ctx.JSON(http.StatusOK, gin.H{
		"data": "ok",
	})
	//token的创建
	jwtService := middlewares.NewJWT()
		expirationTime := time.Now().Add(1 * time.Hour)
		claims := form.CusClaims{
			ID:   rsp.Id,
			Role: rsp.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token, err := jwtService.CreateToken(claims)
		if err != nil {
			zap.S().Error("创建 token 失败", zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建 token 失败"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": "ok",
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "no",
		})
	}
}
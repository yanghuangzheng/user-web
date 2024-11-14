package model

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"user-web/global"
	"user-web/form"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
func GenerateSmsCode(width int) string {
	// 生成width位的短信验证码
	numeric := [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder

	for i := 0; i < width; i++ {
		sb.WriteByte(numeric[rand.Intn(r)])
	}

	return sb.String()
}
func SendSms(ctx *gin.Context) {

	smsform := form.SendSmsForm{}
	if err := ctx.ShouldBind(&smsform); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", "888","888"/*global.SmsConfig.AlismInFo.ApiKey, global.SmsConfig.AlismInFo.ApiSecrect*/) //自己申请的key和select
	if err != nil {
		zap.S().Info(err.Error())
	}
	mobile := smsform.Mobile
	sms := GenerateSmsCode(6)
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https|http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["phoneNumbers"] = mobile                    //手机号
	request.QueryParams["SignName"] = "xxx"                         //阿里云自己验证过的项目名 短信服务的签名
	request.QueryParams["TemplateCode"] = "xxx"                     //阿里云的短信模板 自己设置 模板code
	request.QueryParams["TemplateParam"] = "{\"code\":" + sms + "}" //短信模板中的验证码内容
	response, err := client.ProcessCommonRequest(request)
	fmt.Println(client.DoAction(request, response))
	if err != nil {
		zap.S().Info(err.Error())
	}
	fmt.Printf("response is %#v\n", response) //json数据解析
	//后面注册的时候会将短信验证码带回来注册
	//动态生成数字同时保存验证码 要将验证码和发送的手机号码绑定起来 用 redise
	//一些固态文件
	global.Rdb.Set(context.Background(), mobile, sms, time.Duration(500) * time.Second) //手机验证码有效时间
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})

}

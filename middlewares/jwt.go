package middlewares

import(
	"errors"
	"net/http"
	"time"
    
	"user-web/form"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
// JWT 结构体用于处理JWT相关的操作
type JWT struct {
	SigningKey []byte
}

// 定义一些错误
var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)
// NewJWT 返回一个新的JWT实例
func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte("配置文件"),
	}
}
// CreateToken 创建一个新的JWT令牌
func (j *JWT) CreateToken(claims form.CusClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.SigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
// ParseToken 解析JWT令牌
func (j *JWT) ParseToken(tokenString string) (*form.CusClaims, error) {
	claims := &form.CusClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if !token.Valid {
		return nil, TokenInvalid
	}

	return claims, nil
}

// RefreshToken 更新JWT令牌
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	// 暂时禁用时间函数，以允许解析过期的令牌
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	defer func() { jwt.TimeFunc = time.Now }()

	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &form.CusClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}

	// 获取claims
	if claims, ok := token.Claims.(*form.CusClaims); ok && token.Valid {
		// 更新过期时间
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		// 创建新的令牌
		return j.CreateToken(*claims)
	}

	return "", TokenInvalid
}
// 一个Gin中间件，用于验证JWT令牌
func JWTlc()gin.HandlerFunc{
	return func(c *gin.Context){
    token:=c.Request.Header.Get("x-token")
	if token==""{
		c.JSON(http.StatusUnauthorized, map[string]string{
			"msg": "请登录",
		})
		c.Abort()
		return
	} 
	j := NewJWT()
		// 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"msg": "授权过期",
				})
			} else {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"msg": "未登录",
				})
			}
			c.Abort()
			return

	    }
			// 将解析出来的claims放入context中
			c.Set("claims", claims)
			c.Set("userId", claims.ID)
			c.Next()
}
}
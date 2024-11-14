package form

import(
	"github.com/dgrijalva/jwt-go"
)
type CusClaims struct{
	ID        string
	Role      int32
	jwt.StandardClaims
}
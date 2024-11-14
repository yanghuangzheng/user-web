package form

import (
	"time"
)

type UserRespons struct {
	Id       string     `json:"id"`
	Name     string     `josn:"name"`
	Birthday *time.Time `json:"birthday"`
	Gender   string     `json:"gender"`
	Mobile   string     `json:"mobile"`
	Role     int32      `json:"role"`
}
type MobileRequest struct {
    Mobile string `json:"mobile"`
}
type UserInfo struct{
	PassWord string     `json:"password"`
	Mobile   string     `json:"mobile"`
	Name     string     `json:"name"`
	Birthday *time.Time `json:"birthday"`
	Gender   string     `json:"gender"`
}
type IdRequest struct {
    ID string `json:"id"`
}
type PasswordInfo struct{
	 Password string `json:"password"`
     Id       string `json:"id"`
}
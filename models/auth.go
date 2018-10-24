package models

type UserAuth struct {
	AuthID       int    `gorm:"primary_key"` //授权记录 自增就好
	UId          int    //uid
	IdentityType string //登录类型：手机 邮箱 用户名 或第三方
	Identifier   string //标识：手机 邮箱 用户名或第三方的唯一标识
	Credential   string //密码凭证：密码或TOKEN
}

//验证用户
func CheckAuth(identifier, credential, identityType string) bool {
	var userAuth UserAuth
	db.Select("auth_id").Where(UserAuth{Identifier: identifier, Credential: credential, IdentityType: identityType}).First(&userAuth)

	if userAuth.AuthID > 0 {
		return true
	}

	return false
}

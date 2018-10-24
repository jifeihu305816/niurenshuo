package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserModel struct {
	gorm.Model    //继承
	NickName      string
	Avatar        string          //头像
	Locked        int             //是否锁定
	LastVisitTime time.Time       // 最后登陆时间
	RegisterTime  time.Time       // 注册时间
	Details       UserDetailModel // One-To-One
	Auth          []UserAuthModel // One-To-Many
}

type UserDetailModel struct {
	UId      int    `gorm:"primary_key"` //Guid
	Gender   string //性别
	RealName string // 真实姓名
	Email    string //邮箱

}

type UserAuthModel struct {
	AuthID       int    `gorm:"primary_key"` //授权记录 自增就好
	UId          string //Guid
	IdentityType string //登录类型：手机 邮箱 用户名 或第三方
	Identifier   string //标识：手机 邮箱 用户名或第三方的唯一标识
	Credential   string //密码凭证：密码或TOKEN
}

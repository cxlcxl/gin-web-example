package model

import (
	"gorm.io/gorm"
	"silent-cxl.top/app/model/cache"
)

type User struct {
	db

	Id       int64  `json:"id"`
	Username string `json:"username"`               // 用户名
	Email    string `json:"email"`                  // 登陆邮箱
	Mobile   string `json:"mobile"`                 // 手机号
	State    uint8  `json:"state"`                  // 账号状态
	Secret   string `json:"-" gorm:"column:secret"` // 密码加密符
	Pass     string `json:"-" gorm:"column:pass"`   // 密码

	Timestamp
}

func (m *User) TableName() string {
	return "users"
}

func DbUser(_db *gorm.DB) *User {
	return &User{db: db{DB: _db}}
}

func (m *User) FindUserById(id int64) (user *User, err error) {
	err = cache.NewCache(m.DB).QueryRows(KeyUserById, &user, func(_db *gorm.DB, _v interface{}, _bys ...interface{}) error {
		return _db.Table("users").Where("id = ?", _bys[0]).First(_v).Error
	}, id)
	return
}

func (m *User) List(key string, state int, offset, limit int64) (users []*User, total int64, err error) {
	tbl := m.Table(m.TableName())
	if key != "" {
		tbl = tbl.Where("username like ? or email like ?", "%"+key+"%", "%"+key+"%")
	}
	if state >= 0 {
		tbl = tbl.Where("state = ?", state)
	}
	err = tbl.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	err = tbl.Offset(int(offset)).Limit(int(limit)).Order("id desc").Find(&users).Error
	return
}

package models

import (
	"sync"
	"time"

	"github.com/liangjfblue/docker-all/link-gin-db/internal/pkg/auth"

	"github.com/jinzhu/gorm"

	"gopkg.in/go-playground/validator.v9"
)

type TBUser struct {
	gorm.Model
	Username    string    `gorm:"column:username;type:varchar(100);unique_index" description:"账号"`
	Password    string    `gorm:"column:password;type:varchar(80);null" description:"密码"`
	Email       string    `gorm:"column:email;type:varchar(100);null" description:"邮箱"`
	Phone       string    `gorm:"column:phone;type:varchar(20);null" description:"电话"`
	Sex         int8      `gorm:"column:sex" description:"1.男，2.女"`
	Address     string    `gorm:"column:address;type:varchar(500);null" description:"地址"`
	IsAvailable int8      `gorm:"column:is_available;null" description:"是否可用 1-可用 0-不可用"`
	LastLogin   time.Time `gorm:"column:last_login;type(datetime);null" description:"最后登录时间"`
	LoginIp     string    `gorm:"column:login_ip;type:varchar(20);null" description:"登录IP"`
}

func (t *TBUser) TableName() string {
	return "tb_user"
}

func (t *TBUser) Create() error {
	return DB.Create(t).Error
}

func GetUser(u *TBUser) (*TBUser, error) {
	var user TBUser
	err := DB.Where(u).First(&user).Error
	return &user, err
}

func CountUser() (int64, error) {
	var count int64
	err := DB.Model(&TBUser{}).Count(&count).Error
	return count, err
}
func ListUsers(username string, offset, limit int32) ([]*TBUser, uint64, error) {
	var (
		err   error
		users = make([]*TBUser, 0)
		count uint64
	)

	err = DB.Model(&TBUser{}).Where("user_name LIKE ?", "%"+username+"%").Count(&count).Error
	err = DB.Where("user_name LIKE ?", "%"+username+"%").Offset(offset).Limit(limit).Order("id desc").Find(&users).Error
	return users, count, err
}

func DeleteUser(id uint) error {
	user := TBUser{
		Model: gorm.Model{ID: id},
	}
	return DB.Delete(&user).Error
}

func (t *TBUser) Update() error {
	return DB.Save(t).Error
}

func (t *TBUser) Encrypt() (err error) {
	t.Password, err = auth.Encrypt(t.Password)
	return
}

func (t *TBUser) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func FastListUsers(username string, offset, limit int32) ([]*TBUser, uint64, error) {
	var (
		infos = make([]*TBUser, 0)
		ids   = make([]uint, 0)
	)

	offset, limit = checkPageSize(offset, limit)

	users, count, err := ListUsers(username, offset, limit)
	if err != nil {
		return nil, count, err
	}

	for _, user := range users {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint]*TBUser, len(users)),
	}

	for _, u := range users {
		wg.Add(1)
		go func(u *TBUser) {
			defer wg.Done()

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.ID] = u
		}(u)
	}

	wg.Wait()

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}

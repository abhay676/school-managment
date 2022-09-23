package dal

import (
	"github.com/abhay676/school-managment/services/gatekeeper/config/database"
	"github.com/abhay676/school-managment/services/gatekeeper/utils"
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	EID      string `gorm:"not null; uniqueIndex"`
	Email    string `gorm:"not null;uniqueIndex"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Role     string `gorm:"not null"`
	IsActive bool   `gorm:"default:true"`
}

func (e *Entity) BeforeCreate(tx *gorm.DB) error {
	encryptedPwd, err := utils.GenerateHashFromPassword(e.Password)
	e.Password = encryptedPwd
	eId := utils.CreateUniqueId()
	e.EID = eId
	if err != nil {
		return err
	}
	return nil
}

func CreateEntity(entity *Entity) *gorm.DB {
	return database.DB.Create(entity)
}

func FindEntity(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&Entity{}).Take(dest, conds...)
}

func FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return FindEntity(dest, "email = ?", email)
}

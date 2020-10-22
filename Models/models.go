package Models

import (
	"github.com/jinzhu/gorm"
	"goGinPro/Services"
	// gormigrate "gopkg.in/gormigrate.v1"
)

// 用户
type User struct {
	gorm.Model
	Services.User
	Role *Role `gorm:"rel(fk)"` //设置一对多关系
}

//权限
type Permission struct {
	gorm.Model
	Services.Permission
}

//角色
type Role struct {
	gorm.Model
	Services.Role

	Permissions []*Permission `gorm:"rel(m2m)"` //多对多
}

//3D模型（基础）
type Models struct {
	gorm.Model
	Services.Models
}

//3D模型（重建）
type SceneToModel struct {
	gorm.Model
	Services.SceneToModel
}

//场景
type Scene struct {
	gorm.Model
	Services.Scene
}

func StartMigrate(db *gorm.DB) {
	// m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
	// 	{
	// 		ID: "initial", // 迁移标签
	// 		Migrate: func(tx *gorm.DB) error {
	// 			return tx.CreateTable(&User{}, &Permission{}, &Role{}, &Mush{}).Error
	// 		},
	// 		Rollback: func(tx *gorm.DB) error {
	// 			return tx.DropTable(&User{}, &Permission{}, &Role{}, &Mush{}).Error
	// 		},
	// 	},
	// })
	// return m.Migrate()

	db.AutoMigrate(&User{}, &Permission{}, &Role{}, &Models{}, &SceneToModel{})

	return
}

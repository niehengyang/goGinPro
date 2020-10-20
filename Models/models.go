package Models

import (
	"github.com/jinzhu/gorm"
	// gormigrate "gopkg.in/gormigrate.v1"
)

// 用户
type User struct {
	gorm.Model
	Name     string `gorm:"column(name)" json:"name"`
	Username string `gorm:"column(username)" json:"username"`
	Password string `gorm:"column(password)" json:"password"`
	Email    string `gorm:"column(email)" json:"email"`
	Status   int    `gorm:"column(status)" json:"status"`
	Avatar   string `gorm:"column(avatar)" json:"avatar"`
	Type     string `gorm:"column(type)" json:"type"`
	Phone    string `gorm:"column(phone)" json:"phone"`
	Token    string `gorm:"column(token)" json:"-"`
	Expire   int64  `gorm:"column(expire)" json:"-"`

	Role *Role `gorm:"rel(fk)"` //设置一对多关系
}

//权限
type Permission struct {
	gorm.Model
	Name             string `gorm:"size(100);column(name)" json:"name"`
	Parent           string `gorm:"size(100);column(parent)" json:"parent"`
	SalesmagagerWith string `gorm:"size(10);column(salesmagager_with)" json:"salesmagager_with"`
	Status           int    `gorm:"column(status)" json:"status"`
	Uid              string `gorm:"size(100);column(uid)" json:"uid"`
	Type             string `gorm:"size(10);column(type)" json:"type"`
	Url              string `gorm:"size(100);column(url)" json:"url"`
	SysadminWith     string `gorm:"size(10);column(sysadmin_with)" json:"sysadmin_with"`
	TechnicianWith   string `gorm:"size(10);column(technician_with)" json:"technician_with"`
	Icon             string `gorm:"size(100);column(icon)" json:"icon"`
	Describe         string `gorm:"size(500);column(describe)" json:"describe"`
}

//角色
type Role struct {
	gorm.Model
	Name string `gorm:"size(100);column(name)" json:"name"`
	Desc string `gorm:"size(500);column(desc)" json:"desc"`

	Permissions []*Permission `gorm:"rel(m2m)"` //多对多
}

//3D模型（基础）
type Models struct {
	gorm.Model
	Name                string `gorm:"size:50" json:"name"`
	SceneId             int    `gorm:"size:11" json:"scene_id"`
	GroupId             int    `gorm:"size:11" json:"group_id"`
	ModelUrl            string `gorm:"size:500" json:"model_url"`
	MtlUrl              string `gorm:"size:500" json:"mtl_url"`
	EpidermisColor      string `gorm:"size:100" json:"epidermis_color"`
	PictureUrl          string `gorm:"size:500" json:"picture_url"`
	ModelType           string `gorm:"size:6" json:"model_type"`
	ModelClassification string `gorm:"size:20" json:"model_classification"`
	StructureData       string `gorm:"type:longtext" json:"structure_data"`
	Desc                string `gorm:"size:200" json:"desc"`
}

//3D模型（重建）
type SceneToModel struct {
	gorm.Model
	Name                string `gorm:"size:50" json:"name"`
	SceneId             int    `gorm:"size:11" json:"scene_id"`
	GroupId             int    `gorm:"size:11" json:"group_id"`
	ModelUrl            string `gorm:"size:500" json:"model_url"`
	MtlUrl              string `gorm:"size:500" json:"mtl_url"`
	EpidermisColor      string `gorm:"size:100" json:"epidermis_color"`
	PictureUrl          string `gorm:"size:500" json:"picture_url"`
	ModelType           string `gorm:"size:6" json:"model_type"`
	ModelClassification string `gorm:"size:20" json:"model_classification"`
	StructureData       string `gorm:"type:longtext" json:"structure_data"`
	Desc                string `gorm:"size:200" json:"desc"`
}

//场景
type Scene struct {
	gorm.Model
	Name string `gorm:"size(100);column(name)" json:"name"`
	Desc string `gorm:"size(500);column(desc)" json:"desc"`
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

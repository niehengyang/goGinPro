package Services

// 用户
type User struct {
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
}

//权限
type Permission struct {
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
	Name string `gorm:"size(100);column(name)" json:"name"`
	Desc string `gorm:"size(500);column(desc)" json:"desc"`
}

//3D模型（基础）
type Models struct {
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
	Name string `gorm:"size(100);column(name)" json:"name"`
	Desc string `gorm:"size(500);column(desc)" json:"desc"`
}

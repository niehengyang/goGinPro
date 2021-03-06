package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"goGinPro/Databases"
)

func CreateModel(newModel Models) (uint, bool) {
	result := Mysql.DB.Create(&newModel)
	id := newModel.ID
	if result.Error != nil {
		return 0, false
	}
	return id, true
}

func EditModel(modelId int, newModels Models) bool {

	var model Models
	Mysql.DB.Find(&model, modelId)

	Mysql.DB.Model(&model).Update(newModels)

	return true
}

func GetModelItem(modelId int) (Models, bool) {
	var newModel Models
	result := Mysql.DB.First(&newModel, modelId)
	if result.Error != nil {
		return newModel, false
	}
	return newModel, true
}

func GetAllModel(groupId int) []Models {

	var models []Models

	if groupId > 0 {
		Mysql.DB.Where("group_id = ?", groupId).Find(&models)
	} else {
		Mysql.DB.Find(&models)
	}

	return models
}

func SaveModels(sceneId int, newModels []SceneToModel) bool {

	Mysql.DB.Where("scene_id = ?", sceneId).Delete(&SceneToModel{})
	Mysql.DB.Unscoped().Delete(&SceneToModel{})

	var Null uint
	for _, value := range newModels {
		value.ID = Null
		status := Mysql.DB.Create(&value)
		// id := value.ID
		if status.Error != nil {
			return false
		}
	}
	return true
}

func GetSceneModels(sceneId int) []SceneToModel {
	var models []SceneToModel
	Mysql.DB.Where("scene_id = ?", sceneId).Find(&models)

	return models
}

func DeleteModel(id string) bool {
	Mysql.DB.Delete(&Models{}, id)
	return true
}

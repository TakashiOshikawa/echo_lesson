package designers

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"../../models"
	"../../dao"
)

func FindById(id int) models.Designer {

	var selectedDesigners []models.Designer
	dao.Db.Find(&selectedDesigners)

	return selectedDesigners[0]
}

package database


// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error{
	//AutoMigrate takes a model, model inside have gorm.Model
	//to će definirati sve columne i fieldove za slugove
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
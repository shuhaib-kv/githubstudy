package main

import (
	"gin/gorm/initializers"
	"gin/gorm/models"
)

func init()  {
	// initializers.Loadenv()
initializers.Connectodb()
	
}
func main()  {
	initializers.DB.AutoMigrate(&models.User{})	
}
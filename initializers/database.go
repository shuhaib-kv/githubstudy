package initializers

//here initialized database
import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connectodb() {
	var err error
	// dsn := os.Getenv("DB_URL")
	dsn:="host=localhost user=safwan password=Safwan@123 dbname=gorm port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatal("errpr connecting to database")
	}
}

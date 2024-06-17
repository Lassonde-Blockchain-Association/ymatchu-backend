package student
import (
	"gorm.io/gorm"
)


type Student struct {
	DB *gorm.DB
}

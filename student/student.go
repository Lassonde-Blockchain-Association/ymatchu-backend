package student
import (
	"gorm.io/gorm"
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"

)


type Student struct {
	DB *gorm.DB
	Error error
}

type FilteringParams struct {
	Location database.Location  `json:"location"`
	Utility  database.Utilities `json:"utility"`
	Features database.Features  `json:"features"`
	Price    float32            `json:"price"`
}

type FilterResponse struct {
	PropertyMedia []database.PropertyMedia
	ListingID     string
}
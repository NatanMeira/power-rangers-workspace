package ranger

import (
	"time"
)

type Ranger struct {
	Id            int       `json:"id" gorm:"primary_key;auto_increment"`
	Name          string    `json:"name"`
	Color         string    `json:"color"`
	AnimalFileRef string    `json:"animalFileRef"`
	HelmetFileRef string    `json:"helmetFileRef"`
	CreatedAt     time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at"`
}

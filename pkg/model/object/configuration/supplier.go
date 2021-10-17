package configuration

type Supplier struct {
	Id      int64  `json:"id" gorm:"not null;unique;primaryKey;autoIncrement"`
	Name    string `json:"name" gorm:"not null"`
	Phone   string `json:"phone" gorm:"not null"`
	Email   string `json:"email" gorm:"not null"`
	Address string `json:"address"`
}

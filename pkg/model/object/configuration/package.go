package configuration

type Package struct {
	Id      int64 `json:"id" gorm:"not null;unique;primaryKey;autoIncrement"`
	Motor   int64 `json:"motor" gorm:"not null"`
	Battery int64 `json:"battery" gorm:"not null"`
	Charger int64 `json:"charger" gorm:"not null"`
	Rigging int64 `json:"rigging" gorm:"not null"`
}

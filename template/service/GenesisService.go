package service

import (
	"github.com/jinzhu/gorm"
)

type GenesisService struct {
	db *gorm.DB
}

type Genesis struct {
	Id      int    `gorm:"primary_key"`
	Name    string `gorm:"not null"`
	Sex     int    `gorm:"default:1"` // 1：男，2:女 0:未知
	Phone   string `gorm:"not null;unique_index:phone_idx"`
	Address string `gorm:"not null;index:address_idx"`
}

func NewGenesisService(db *gorm.DB) *GenesisService {
	return &GenesisService{
		db: db,
	}
}

func (s *GenesisService) GetById(id int) Genesis {
	genesis := Genesis{}
	s.db.Raw("select * from genesis t where t.id = ? ", id).Scan(&genesis)
	return genesis
}

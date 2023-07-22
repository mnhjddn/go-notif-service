package notif

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetNotification(id int) (data NotificationData, err error) {
	err = r.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

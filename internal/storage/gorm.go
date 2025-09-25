package storage

import (
	"log"

	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	if err := db.AutoMigrate(&Contact{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
	return &GormStore{db: db}
}

func (s *GormStore) Ajouter(c Contact) Contact {
	err := s.db.Create(&c).Error
	if err != nil {
		log.Printf("Ajouter: %v", err)
	}
	return c
}

func (s *GormStore) Lister() []Contact {
	var list []Contact
	err := s.db.Order("id ASC").Find(&list).Error
	if err != nil {
		log.Printf("Lister: %v", err)
	}
	return list
}

func (s *GormStore) Supprimer(ID int) bool {
	tx := s.db.Delete(&Contact{}, ID)
	if tx.Error != nil {
		log.Printf("Supprimer: %v", tx.Error)
		return false
	}
	return tx.RowsAffected > 0
}

func (s *GormStore) Recuperer(ID int) (Contact, bool) {
	var c Contact
	err := s.db.First(&c, ID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Contact{}, false
		}
		log.Printf("Recuperer: %v", err)
		return Contact{}, false
	}
	return c, true
}

func (s *GormStore) MettreAJour(c Contact) (Contact, bool) {
	var existing Contact
	err := s.db.First(&existing, c.ID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Contact{}, false
		}
		log.Printf("MettreAJour (fetch): %v", err)
		return Contact{}, false
	}

	if err := s.db.Model(&existing).Updates(map[string]any{
		"nom":   c.Nom,
		"email": c.Email,
	}).Error; err != nil {
		log.Printf("MettreAJour (update): %v", err)
		return Contact{}, false
	}
	err1 := s.db.First(&existing, c.ID).Error
	if err1 != nil {
		log.Printf("MettreAJour (reload): %v", err1)
		return Contact{}, false
	}
	return existing, true
}

func (s *GormStore) NextID() int {
	var maxID uint
	// SELECT COALESCE(MAX(id), 0) FROM contacts;
	s.db.Model(&Contact{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID)
	return int(maxID + 1)
}

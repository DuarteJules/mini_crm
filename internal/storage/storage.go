package storage

import "fmt"

type Contact struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"contact_id"`
	Nom   string `gorm:"size:100;not null" json:"nom"`
	Email string `gorm:"size:150;uniqueIndex;not null" json:"email"`
}

type Storer interface {
	Ajouter(c Contact) Contact
	Lister() []Contact
	Supprimer(ID int) bool
	MettreAJour(c Contact) (Contact, bool)
	Recuperer(ID int) (Contact, bool)
	NextID() int
}

var ErrContactNotFound = func(id int) error {
	return fmt.Errorf("Contact avec l'ID non trouvé")
}

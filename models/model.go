package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // imports postgres driver
)

var errInvalidID error = errors.New("invalid ID")

// Message represents message data in the database
type Message struct {
	gorm.Model
	Text string `gorm:"not null"`
}

// MessageDB is an interface for interacting with message data in the database
type MessageDB interface {
	GetAll() ([]Message, error)

	Create(message *Message) error
	Delete(id uint) error
}

// MessageService is an interface for interacting with message model
type MessageService interface {
	MessageDB
}

type messageService struct {
	MessageDB
	db *gorm.DB
}

func (ms *messageService) GetAll() ([]Message, error) {
	messages := make([]Message, 0)
	err := ms.db.Find(&messages).Error

	return messages, err
}

func (ms *messageService) Create(message *Message) error {
	// validations

	// insert into database
	return ms.db.Create(&message).Error
}

func (ms *messageService) Delete(id uint) error {
	// validations
	if id <= 0 {
		return errInvalidID
	}

	// delete from database
	message := Message{}
	message.ID = id

	return ms.db.Delete(&message).Error
}

// NewMessageService creates MessageService instance
func NewMessageService(dialect, connectionInfo string) MessageService {
	db, err := gorm.Open(dialect, connectionInfo)
	if err != nil {
		log.Println("Failed to open database connection!")
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&Message{})

	return &messageService{
		db: db,
	}
}

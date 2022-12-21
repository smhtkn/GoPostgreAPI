package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Client struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *Client) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *Client) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Client) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *Client) SaveClient(db *gorm.DB) (*Client, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Client{}, err
	}
	return u, nil
}

func (u *Client) FindAllClients(db *gorm.DB) (*[]Client, error) {
	var err error
	clients := []Client{}
	err = db.Debug().Model(&Client{}).Limit(100).Find(&clients).Error
	if err != nil {
		return &[]Client{}, err
	}
	return &clients, err
}

func (u *Client) FindClientByID(db *gorm.DB, uid uint32) (*Client, error) {
	var err error
	err = db.Debug().Model(Client{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Client{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Client{}, errors.New("Client Not Found")
	}
	return u, err
}

func (u *Client) UpdateAClient(db *gorm.DB, uid uint32) (*Client, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Client{}).Where("id = ?", uid).Take(&Client{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.Nickname,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Client{}, db.Error
	}
	// This is the display the updated client
	err = db.Debug().Model(&Client{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Client{}, err
	}
	return u, nil
}

func (u *Client) DeleteAClient(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Client{}).Where("id = ?", uid).Take(&Client{}).Delete(&Client{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

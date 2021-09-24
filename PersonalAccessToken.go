package mars

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"time"

	"github.com/ferdiunal/venus"
	"gorm.io/datatypes"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonalAccessToken struct {
	ID         uint64         `gorm:"autoIncrement;primaryKey"`
	UserId     uint64         `gorm:"foreignkey:id;references:user_id;constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
	Name       string         `gorm:"uniqueIndex;size:255"`
	Token      string         `gorm:"uniqueIndex;size:64"`
	Abilities  datatypes.JSON `gorm:"type:json"`
	LastUsedAt sql.NullTime   `gorm:"type:timestamp;"`
	ExpireAt   time.Time      `gorm:"type:timestamp;"`
	CreatedAt  time.Time      `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt  sql.NullTime   `gorm:"type:timestamp;"`
}

type MarsToken struct {
	AccessToken string
	ExpiresIn   time.Time
	Abilities   []string
}

func (m *PersonalAccessToken) HashedToken() string {
	return fmt.Sprintf(
		"%x",
		sha256.Sum256(
			[]byte(m.Token),
		),
	)
}

func GeneratePlainTextToken() string {
	token, _ := uuid.NewRandom()
	return token.String()
}

func (m *PersonalAccessToken) BeforeCreate(tx *gorm.DB) (err error) {
	m.Token = GeneratePlainTextToken()

	return nil
}

func (m *PersonalAccessToken) GetToken(venus venus.VenusInterface) string {
	return fmt.Sprintf("%v|%v", venus.Encode(int64(m.ID)), m.HashedToken())
}

func (m *PersonalAccessToken) GetResult(venus venus.VenusInterface) *MarsToken {
	return NewMarsToken(m, venus)
}

func NewMarsToken(accessToken *PersonalAccessToken, venus venus.VenusInterface) *MarsToken {
	return &MarsToken{
		AccessToken: accessToken.GetToken(venus),
		ExpiresIn:   accessToken.ExpireAt,
		Abilities:   accessToken.Abilities,
	}
}
func NewPersonalAccessToken(userId uint64, name string, abilities []string) *PersonalAccessToken {

	if len(abilities) == 0 {
		abilities = []string{"*"}
	}
	return &PersonalAccessToken{
		UserId:    userId,
		Name:      name,
		Abilities: abilities,
	}
}

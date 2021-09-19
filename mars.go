package mars

import (
	"crypto/sha256"
	"fmt"

	"github.com/ferdiunal/mars/models"
	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"
)

type Mars struct {
	db             *gorm.DB
	plainTextToken string
}

func (m *Mars) hash() string {
	return fmt.Sprintf(
		"%x",
		sha256.Sum256(
			[]byte(m.plainTextToken),
		),
	)
}

func (m *Mars) GeneratePlainTextToken() *Mars {
	m.plainTextToken = utils.UUIDv4()

	return m
}

func (m *Mars) CreateToken(userId uint64, name string, abilities []string) string {
	token := &models.PersonalAccessToken{
		Name:      name,
		Abilities: abilities,
		Token:     m.hash(),
		UserId:    userId,
	}

	m.db.Create(token)

	return fmt.Sprintf("%v|%v", token.ID, m.plainTextToken)
}

func NewMars(db *gorm.DB) *Mars {
	return &Mars{
		db: db,
	}
}

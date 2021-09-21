package mars

import (
	"errors"

	"gorm.io/gorm"
)

type Mars struct {
	db *gorm.DB
}

type MarsInterface interface {
	CreateToken(userId uint64, name string, abilities []string) *MarsToken
	CheckToken(userId uint64, tokenId uint64, token string) (*PersonalAccessToken, error)
	RevokeToken(userId uint64, tokenId uint64, token string) error
}

func (m *Mars) CreateToken(userId uint64, name string, abilities []string) *MarsToken {
	model := NewPersonalAccessToken(userId, name, abilities)
	m.db.Create(model)

	return model.GetResult()
}

func (m *Mars) CheckToken(userId uint64, tokenId uint64, token string) (*PersonalAccessToken, error) {
	model := &PersonalAccessToken{}
	m.db.Where("id = ? and userId = ?", tokenId, userId).Find(model)

	if ok := model.HashedToken() == token; ok {
		return model, nil
	}

	return nil, errors.New("token not found")
}

func (m *Mars) RevokeToken(userId uint64, tokenId uint64, token string) error {
	model, err := m.CheckToken(userId, tokenId, token)

	if err != nil {
		return err
	}

	m.db.Delete(model)

	return nil
}

func NewMars(db *gorm.DB) MarsInterface {
	return &Mars{
		db: db,
	}
}

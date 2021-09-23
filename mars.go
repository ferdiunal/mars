package mars

import (
	"errors"
	"github/ferdiunal/venus"

	"gorm.io/gorm"
)

type Mars struct {
	db    *gorm.DB
	venus venus.VenusInterface
}

type MarsInterface interface {
	CreateToken(userId uint64, name string, abilities []string) *MarsToken
	CheckToken(tokenId uint64, token string) (*PersonalAccessToken, error)
	RevokeToken(userId uint64, tokenId uint64, token string) error
}

func (m *Mars) CreateToken(userId uint64, name string, abilities []string) *MarsToken {
	model := NewPersonalAccessToken(userId, name, abilities)
	m.db.Create(model)

	return model.GetResult(m.venus)
}

func (m *Mars) CheckToken(tokenId uint64, token string) (*PersonalAccessToken, error) {
	model := &PersonalAccessToken{}
	m.db.Where("id = ? and userId = ?", tokenId).Find(model)

	if ok := model.HashedToken() == token; ok {
		return model, nil
	}

	return nil, errors.New("token not found")
}

func (m *Mars) RevokeToken(userId uint64, tokenId uint64, token string) error {
	model, err := m.CheckToken(tokenId, token)

	if err != nil {
		return err
	}

	if model.UserId != userId {
		return errors.New("token not matched user id") // Burasını değiştir ileride
	}

	m.db.Delete(model)

	return nil
}

func NewMars(db *gorm.DB, salt string, len int) MarsInterface {
	hashids := venus.New(salt, len)
	return &Mars{
		db:    db,
		venus: hashids,
	}
}

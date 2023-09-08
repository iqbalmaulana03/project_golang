package user

import "gorm.io/gorm"

type UserRepository interface {
	Save(user User) (User, error)
}

type userRepositroy struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *userRepositroy {
	return &userRepositroy{db}
}

func (r *userRepositroy) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

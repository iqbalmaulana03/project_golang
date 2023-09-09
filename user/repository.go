package user

import "gorm.io/gorm"

type UserRepository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
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

func (r *userRepositroy) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

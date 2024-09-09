package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	UpdateUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

// UpdateUser implements Repository.
func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// FindById implements Repository.
func (r *repository) FindById(id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// FindByEmail implements Repository.
func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

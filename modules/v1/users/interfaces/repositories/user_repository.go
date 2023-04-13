package repository

import "mygram/modules/v1/users/domain"

func (r *Repository) Save(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

// func (r *Repository) FindUser(field string, value string) (domain.User, error) {
// 	var user domain.User
// 	err := r.db.Where(field+" = ?", value).First(&user).Error
// 	return user, err
// }

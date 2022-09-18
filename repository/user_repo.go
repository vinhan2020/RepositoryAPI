package repository

import (
	"DemoApiPostgre/models"
)

func (p *postgreRepo) GetAllUser() (users []models.User, err error) {
	err = p.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (p *postgreRepo) CreateUser(user models.User) (models.User, error) {
	err := p.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (p *postgreRepo) GetUserByID(condition map[string]interface{}) (user models.User, err error) {
	err = p.db.Where(condition).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

func (p *postgreRepo) UpdateUser(condition map[string]interface{}) (user models.User, err error) {
	err = p.db.Where(condition).Updates(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (p *postgreRepo) DeleteUser(condition map[string]interface{}) (err error) {
	var user models.User
	err = p.db.Where(condition).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

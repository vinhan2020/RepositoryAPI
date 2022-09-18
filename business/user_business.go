package business

import "DemoApiPostgre/models"

type UserRepoInterface interface {
	GetAllUser() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	GetUserByID(condition map[string]interface{}) (models.User, error)
	UpdateUser(condition map[string]interface{}) (models.User, error)
	DeleteUser(condition map[string]interface{}) error
}

type UserBusiness struct {
	userRepoInterface UserRepoInterface
}

func NewUserService(userRepoInterface UserRepoInterface) *UserBusiness {
	return &UserBusiness{userRepoInterface: userRepoInterface}
}

func (u *UserBusiness) GetAllUser() ([]models.User, error) {
	users, err := u.userRepoInterface.GetAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserBusiness) CreateUser(user models.User) (userResponse models.User, err error) {
	userResponse, err = u.userRepoInterface.CreateUser(user)
	if err != nil {
		return userResponse, err
	}
	return userResponse, nil
}

func (u *UserBusiness) GetUserByID(condition map[string]interface{}) (user models.User, err error) {
	user, err = u.userRepoInterface.GetUserByID(condition)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserBusiness) UpdateUser(condition map[string]interface{}) (user models.User, err error) {
	user, err = u.userRepoInterface.GetUserByID(condition)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserBusiness) DeleteUser(condition map[string]interface{}) (err error) {
	err = u.userRepoInterface.DeleteUser(condition)
	if err != nil {
		return err
	}
	return nil
}

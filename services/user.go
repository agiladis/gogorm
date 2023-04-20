package services

import (
	"errors"
	"gogorm/entity"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (us *UserService) Create(user entity.User) (entity.User, error) {
	if err := us.DB.Create(&user).Error; err != nil {
		// ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// 	"message": err.Error(),
		// })
		return entity.User{}, err
	}

	return user, nil
}

func (us *UserService) Update(id uint, user entity.User) error {

	result := us.DB.Model(entity.User{}).Where("id = ?", id).Updates(&user)

	if result.RowsAffected == 0 {
		// ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// 	"message": "nothing to update",
		// })
		return errors.New("there is no data to update")
	}

	return nil
}

func (us *UserService) Delete(id uint) error {
	var user entity.User

	if err := us.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		// ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// 	"message": "Data not found",
		// })
		return err
	}
	return nil
}

func (us *UserService) GetUserById(id uint) (entity.User, error) {
	var user entity.User

	if err := us.DB.Where("id = ?", id).First(&user).Error; err != nil {
		// ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// 	"message": "Data not found",
		// })
		return user, err
	}

	return user, nil
}

func (us *UserService) GetAll() ([]entity.User, error) {
	var users []entity.User
	if err := us.DB.Find(&users).Error; err != nil {
		// ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		// 	"message": "Data not found",
		// })
		return users, err
	}
	return users, nil
}

package repository

import (
	"github.com/kalyasik/blog_backend/internal/models"
	"github.com/kalyasik/blog_backend/pkg/database/postgres"
)

func GetPostByID(postID string) (models.Posts, error) {
	postModel := models.Posts{}

	err := postgres.DB.Where("id = ?", postID).First(&postModel).Error
	if err != nil {
		return models.Posts{}, err
	}

	return postModel, nil
}

func GetAllPosts() ([]models.Posts, error) {
	var postModels []models.Posts

	err := postgres.DB.Find(&postModels).Error
	if err != nil {
		return nil, err
	}

	return postModels, nil
}

func CreateNewPost(title, body string, user_id int64) (models.Posts, error) {
	postModel := models.Posts{Title: title, Body: body, UserID: user_id}

	err := postgres.DB.Create(&postModel).Error
	if err != nil {
		return models.Posts{}, err
	}

	return postModel, nil
}

func UpdatePostByID(postID, user_id int64, title, body string) (int64, error) {
	postModel := models.Posts{Title: title, Body: body, UserID: user_id}

	err := postgres.DB.Model(&models.Posts{}).Where("id = ?", postID).Updates(models.Posts{
		Title:  postModel.Title,
		Body:   postModel.Body,
		UserID: postModel.UserID,
	}).Error
	if err != nil {
		return 0, err
	}

	return postID, nil
}

func DeletePost(postID int64) (int64, error) {
	postModel := models.Posts{}

	err := postgres.DB.Delete(&postModel).Error
	if err != nil {
		return 0, err
	}

	return postID, nil
}

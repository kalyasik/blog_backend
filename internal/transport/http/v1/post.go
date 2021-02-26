package v1

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kalyasik/blog_backend/internal/models"
	"github.com/kalyasik/blog_backend/internal/repository"
	"github.com/labstack/echo/v4"
)

func GetPostsHandler(c echo.Context) error {
	posts, err := repository.GetAllPosts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseMessage{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: http.StatusOK,
		Data:   posts,
	})
}

func CreatePostHandler(c echo.Context) error {
	postModel := new(models.Posts)
	if err := c.Bind(postModel); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := validator.New().Struct(postModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	post, err := repository.CreateNewPost(postModel.Title, postModel.Body, postModel.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseMessage{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: http.StatusOK,
		Data:   post,
	})
}

func UpdatePostHandler(c echo.Context) error {
	postModel := new(models.Posts)
	if err := c.Bind(postModel); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := validator.New().Struct(postModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	post, err := repository.GetPostByID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseMessage{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		})
	}

	res, err := repository.UpdatePostByID(post.ID, postModel.UserID, postModel.Title, postModel.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseMessage{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: http.StatusOK,
		ID:     res,
	})
}

func DeletePostHandler(c echo.Context) error {
	post, err := repository.GetPostByID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseMessage{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		})
	}

	deletedID, err := repository.DeletePost(post.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseMessage{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: http.StatusOK,
		ID:     deletedID,
	})
}

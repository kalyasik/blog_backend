package v1

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kalyasik/blog_backend/internal/models"
	"github.com/kalyasik/blog_backend/internal/repository"
	"github.com/labstack/echo/v4"
)

// @Summary Get all posts
// @Tags Posts
// @Description returning all post
// @ModuleID getAllPosts
// @Accept json
// @Produce json
// @Success 200 {object} SuccessMesageData
// @Failure 400,404 {object} ErrorMessage
// @Failure 500 {object} ErrorMessage
// @Failure default {object} ErrorMessage
// @Router /api/v1/posts [get]
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

// @Summary Create new post
// @Tags Posts
// @Description created new post
// @ModuleID createNewPost
// @Accept json
// @Produce json
// @Param data body models.PostInput true "Enter post data to create a new post"
// @Success 200 {object} SuccessMesageData
// @Failure 400,404 {object} ErrorMessage
// @Failure 500 {object} ErrorMessage
// @Failure default {object} ErrorMessage
// @Router /api/v1/posts [post]
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

// @Summary Update post by ID
// @Tags Posts
// @Description updated post data
// @ModuleID updatePostByID
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param data body models.PostInput true "Enter post data to update a post"
// @Success 200 {object} SuccessMesageData
// @Failure 400,404 {object} ErrorMessage
// @Failure 500 {object} ErrorMessage
// @Failure default {object} ErrorMessage
// @Router /api/v1/posts/{id} [put]
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

// @Summary Delete post by ID
// @Tags Posts
// @Description deleted post data
// @ModuleID deletePostByID
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} SuccessMesageData
// @Failure 400,404 {object} ErrorMessage
// @Failure 500 {object} ErrorMessage
// @Failure default {object} ErrorMessage
// @Router /api/v1/posts/{id} [delete]
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

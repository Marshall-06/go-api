// package controllers

// import (
// 	"net/http"
// 	"go-api/database"
// 	"go-api/models"

// 	"github.com/gin-gonic/gin"
// )

// // Create User
// func CreateUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := database.DB.Create(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, user)
// }

// // Get All Users
// func GetUsers(c *gin.Context) {
// 	var users []models.User
// 	database.DB.Find(&users)
// 	c.JSON(http.StatusOK, users)
// }

// // Get User by ID
// func GetUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User

// 	if err := database.DB.First(&user, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

// // Update User
// func UpdateUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User

// 	if err := database.DB.First(&user, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	database.DB.Save(&user)
// 	c.JSON(http.StatusOK, user)
// }

// // Delete User
// func DeleteUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User

// 	if err := database.DB.First(&user, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	database.DB.Delete(&user)
// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
// }



package controllers

import (
	"net/http"
	"go-api/database"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary      Create a new user
// @Description  Adds a new user to the database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  models.User  true  "User Data"
// @Success      201  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsers godoc
// @Summary      Get all users
// @Description  Returns a list of users
// @Tags         users
// @Produce      json
// @Success      200  {array}   models.User
// @Router       /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary      Get a user by ID
// @Description  Returns a single user
// @Tags         users
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  Updates user details
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Param        user  body  models.User  true  "User Data"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Router       /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Removes a user from the database
// @Tags         users
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/atrem13/golang-api-user/database"
	"github.com/atrem13/golang-api-user/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewUSer() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(repository.Db, &user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepo) GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepo) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	err := models.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err})
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepo) UpdateUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&user)
	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepo) DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

package example

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// User represents a user in the system
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "雷大宇"},
	{ID: 2, Name: "丹丹"},
}

/**
 * @title: Gin-Gonic
 * @description: Gin-Gonic crud
 */
func aaaaaa() {

	r := gin.Default()
	// GET /users - Get all users
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// GET /users/:id - Get a single user by ID
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			paramId, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}

			if paramId == user.ID {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	})

	// POST /users - Create a new user
	r.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users = append(users, newUser)
		c.JSON(http.StatusCreated, newUser)
	})

	// PUT /users/:id - Update an existing user by ID
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedUser User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, user := range users {
			paramId, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}

			if paramId == user.ID {
				users[i] = updatedUser
				c.JSON(http.StatusOK, updatedUser)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	})

	// DELETE /users/:id - Delete a user by ID
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, user := range users {
			paramId, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}

			if paramId == user.ID {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

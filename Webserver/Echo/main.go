package main

// imports
import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create a user struct
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
}

// create a user slice
var users = []User{
	User{1, "John", "CEO"},
	User{2, "Mary", "HR"},
	User{3, "Mike", "Developer"},
	User{4, "Jane", "Designer"},
}

// get all users
func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// get a user by id
func getUser(c echo.Context) error {
	// get the id from the url
	id := c.Param("id")

	// convert the id to an int
	userID, _ := strconv.Atoi(id)

	// find the user
	for _, user := range users {
		if user.ID == userID {
			return c.JSON(http.StatusOK, user)
		}
	}

	// if the user is not found
	return c.JSON(http.StatusNotFound, "User not found")
}

// Add a new user
func addUser(c echo.Context) error {
	// get the user from the request
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	// add the user to the slice
	users = append(users, *user)

	return c.JSON(http.StatusOK, echo.Map{"message": "User added successfully!", "user": user})
}

func updateUser(c echo.Context) error {
	// get the id from the url
	id := c.QueryParam("id")

	// convert the id to an int
	userID, _ := strconv.Atoi(id)

	// find the user
	for i, user := range users {
		if user.ID == userID {
			// get the user from the request
			if err := c.Bind(&user); err != nil {
				return err
			}

			// update the user
			users[i] = user

			return c.JSON(http.StatusOK, user)
		}
	}

	// if the user is not found
	return c.JSON(http.StatusNotFound, "User not found")
}

// Delete a user
func deleteUser(c echo.Context) error {
	// get the id from the url
	id := c.QueryParam("id")

	// convert the id to an int
	userID, _ := strconv.Atoi(id)

	// find the user
	for i, user := range users {
		if user.ID == userID {
			// delete the user
			users = append(users[:i], users[i+1:]...)

			return c.JSON(http.StatusOK, echo.Map{"message": "User deleted successfully!"})
		}
	}

	// if the user is not found
	return c.JSON(http.StatusNotFound, "User not found")
}

// main function
func main() {
	// create a new echo instance
	e := echo.New()

	// route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// route => handler
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", addUser)
	e.PUT("/users", updateUser)
	// start the server
	e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string  `json:"id"`
	JOB string  `json:"job"`
	COMPLETED bool   `json:"completed"`
}

var todos = []todo{
	{ID: "Oke", JOB: "cook food", COMPLETED: false},
	{ID: "Temi", JOB: "sweep the house", COMPLETED: false},
	{ID: "Blossom", JOB: "shut the fuck up", COMPLETED: false},
	
}

func main () {
       router := gin.Default()
	   router.GET("/todos", getTodosHandler)
	   router.GET("/todos/:id", getTodosByIdHandler)
	   router.PATCH("/todos/:id", toggleTodoStatus)
	   router.POST("/todos", postTodosHandler)
	   
       router.Run("localhost:2055")
}

// getTodosHandler returns a list of all todos as json
func getTodosHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

// postTodosHandler adds a new list of todos as json
func postTodosHandler(c *gin.Context) {
// create new var to store the add request 
   var newTodo todo 
// then create an if error statement 
  if err := c.BindJSON(&newTodo); err != nil {
	return
  }
  // add new todo to the struct
  todos = append(todos, newTodo)
  c.IndentedJSON(http.StatusCreated, newTodo)
}

// getTodoById locates the todo whose ID value matches the id 
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodosByIdHandler (c *gin.Context) {
    id := c.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"})
        return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

// toggleTodoStatusHandler locates the todos whose id matches the id and toggle it's status
func toggleTodoStatus(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"})
		return
	}
	todo.COMPLETED = !todo.COMPLETED
	c.IndentedJSON(http.StatusOK, todo)
}



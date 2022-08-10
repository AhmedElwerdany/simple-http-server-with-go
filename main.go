package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


type Todo struct{
	Title string `json:"title"`
	Id int `json:"id"`
	Status int `json:"status"`// 0 | 1
}

const endpoints string = `
GET /todos => return list of todos<br />
GET /todo/:id => return todo by id<br />
POST /todos => create a todo<br />
PUT /todo/:id => updates a todo<br />`

func main() {
	r := gin.Default()

	var todos []Todo
	
	r.GET("/", func(c *gin.Context) {		
		c.String(http.StatusOK, endpoints)
	})

	
	r.GET("/todos", func (c *gin.Context) {
	
		c.JSON(http.StatusOK, gin.H{
			"data" : todos,
			"messages" : "todo recived successfully",
			})

			return 
		})


	r.GET("/todos/:id", func (c *gin.Context) {

		id := c.Param("id")

		for _, todo := range todos {
				if fmt.Sprint(todo.Id) == id {
					c.JSON(http.StatusOK, gin.H{
						"data" : todo,
					})
					return
				}
		}

		c.JSON(http.StatusOK, gin.H {
			"data" : struct{}{},
			"message": "no data found",
		})
		
		return
	})

	r.PUT("/todos/:id", func (c *gin.Context) {
		var request Todo
		id := c.Param("id")


		// check if the request is not correct
		if err := c.ShouldBindJSON(&request); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message" : "fix your error, little boy",
				})
				return
		}


		// filtering the todos based on the "id"
		for index, todo := range todos {
				if fmt.Sprint(todo.Id) == id {
					
					// this could've been done better
					todo.Status = request.Status
					todo.Title = request.Title
					
					todos[index] = todo
					
					c.JSON(http.StatusOK, gin.H{
						"data" : todo,
						"message": "todo updated succesfuly",
					})
					
					return
				}
		}

		c.JSON(http.StatusOK, gin.H {
			"data" : struct{}{},
			"message": "no data found",
		})
		
		return
	})

	
	r.POST("/todos", func (c *gin.Context) {
		
		var request Todo

		request.Id = len(todos) + 1
		
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message" : "fix your error, little boy",
			})
			return
		}

		todos = append(todos, request)

		
		c.JSON(http.StatusOK, gin.H{
			"data" : todos,
			"messages" : "todo added successfully",
		})
		return
	})
	
	r.GET("/me", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})


	

	r.Run()
}

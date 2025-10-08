package main

import (
	"html/template"
	"log"
	"net/http"


	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	InitDB()
	defer DB.Close()

	r := gin.Default()

	r.Static("/static", "./static")

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	r.GET("/", func(c *gin.Context) {
		rows, err := DB.Query("SELECT id, name, email, age FROM users")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			users = append(users, u)
		}

		c.Status(http.StatusOK)
		tmpl.Execute(c.Writer, users)
	})

	log.Println("Server started at http://localhost:8081")
	r.Run(":8081")
}
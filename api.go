package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")

	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	r.SetHTMLTemplate(tmpl)

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"Title": "Go + Gin + TailwindCSS",
			"Body":  "This is rendered using Gin with Go templates and Tailwind!",
		})
	})

	r.Run(":8081")
}
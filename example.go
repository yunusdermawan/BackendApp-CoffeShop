package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Example() {
	r := gin.Default()
	r.GET("/", examples)
	r.GET("/query", queryString)
	r.GET("/params/:username/:slug", paramsString)
	r.POST("/login", reqBody)

	r.Run(":4545")
}

func examples(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "Hello from gin",
	})
}

type qString struct {
	Limit string `form:"limit"`
	Page  string `form:"page"`
}

// localhost:4545/query?page=1&limit=2
func queryString(ctx *gin.Context) {
	// page := ctx.Query("page")
	// limit := ctx.Query("limit")

	var data qString

	if err := ctx.ShouldBind(&data); err != nil {
		log.Println(err)
	}

	ctx.JSON(200, gin.H{
		"page":  data.Page,
		"limit": data.Limit,
	})
}

type pString struct {
	Username string `uri:"username"`
	Slug     string `uri:"slug"`
}

// localhost:4545/params/:username/:slug
func paramsString(ctx *gin.Context) {
	// username := ctx.Param("username")
	// slug := ctx.Param("slug")

	var data pString

	if err := ctx.ShouldBindUri(&data); err != nil {
		log.Println(err)
	}

	ctx.JSON(200, gin.H{
		"username": data.Username,
		"slug":     data.Slug,
	})
}

type body struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// localhost:4545/params/:username/:slug
func reqBody(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		log.Println(err)
	}

	log.Println(file.Filename)
	var data body

	if err := ctx.ShouldBind(&data); err != nil {
		log.Println(err)
	}

	// ctx.JSON(200, gin.H{
	// 	"username": data.Username,
	// 	"password": data.Password,
	// })
	ctx.JSON(200, data)
}

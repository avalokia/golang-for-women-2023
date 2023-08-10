package controllers

import (
	"fmt"
	"net/http"
	"webserver/model"

	"github.com/gin-gonic/gin"
)

var UserData = []model.User{
	{
		Email:   "john.doe@example.com",
		Address: "123 Main St, Cityville, USA",
		Job:     "Software Engineer",
		Reason:  "Sunshine and warmth.",
	},
	{
		Email:   "jane.smith@example.com",
		Address: "456 Elm Rd, Townsville, USA",
		Job:     "Marketing Manager",
		Reason:  "Laughter shared with friends.",
	},
	{
		Email:   "michael.johnson@example.com",
		Address: "789 Oak Ave, Villageland, USA",
		Job:     "Accountant",
		Reason:  "Healthy, vibrant moments.",
	},
	{
		Email:   "sarah.wilson@example.com",
		Address: "101 Pine Lane, Hamletown, USA",
		Job:     "Graphic Designer",
		Reason:  "Love's comforting embrace.",
	},
	{
		Email:   "robert.jackson@example.com",
		Address: "222 Cedar Blvd, Metropolis, USA",
		Job:     "Sales Representative",
		Reason:  "New adventures await you.",
	},
}

func GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	fmt.Println("email:", email)
	found := false
	var retrievedUser model.User

	for i, user := range UserData {
		if email == user.Email {
			found = true
			retrievedUser = UserData[i]
			break
		}
	}

	if !found {
		// ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		// 	"data":    nil,
		// 	"message": fmt.Sprintf("Email %v is not found", email),
		// })
		ctx.HTML(http.StatusNotFound, "notfound.html", gin.H{
			"title": "Email is not registered!",
		})
	} else {
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"data":    retrievedUser,
		// 	"message": "User is retrieved",
		// })
		ctx.HTML(http.StatusOK, "user.html", gin.H{
			"title":   "Profile page",
			"email":   retrievedUser.Email,
			"address": retrievedUser.Address,
			"job":     retrievedUser.Job,
			"reason":  retrievedUser.Reason,
		})
	}
}

func LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "template.html", gin.H{
		"title": "Login Form",
	})
}

func UserPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user.html", gin.H{
		"title": "Login Form",
	})
}

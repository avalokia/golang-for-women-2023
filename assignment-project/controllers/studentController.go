package controllers

import (
	"assignment-project/database"
	"assignment-project/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create new student
func CreateStudent(c *gin.Context) {
	// Bind URL
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	// Get database
	db := database.GetDB()
	// Create new student and check the error
	err = db.Create(&student).Error
	if err != nil {
		log.Fatalln("Error creating new student data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	// Return created student
	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Data":    student,
	})
}

// Get all students
func GetAllStudents(c *gin.Context) {
	// Get database
	db := database.GetDB()
	// Initialize students variable
	students := []models.Student{}
	// Get students and their scores data
	result := db.Preload("Scores").Find(&students)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    students,
	})
}

// Update student by ID
func UpdateStudentByID(c *gin.Context) {
	// Retrieve param and convert into int
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
	}

	// Bind request JSON
	var newStudent models.Student
	err = c.BindJSON(&newStudent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	// Get database
	db := database.GetDB()

	// Retrieve existing student & scores
	existingStudent := models.Student{}
	result := db.Preload("Scores").First(&existingStudent, uint(studentID))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve student and scores"})
		return
	}

	// Start a transaction
	tx := db.Begin()

	// Update existing student with new data from request
	existingStudent.Name = newStudent.Name
	existingStudent.Age = newStudent.Age

	// Iterate through request's Scores to find whether there are any change for the Scores
	for idx, newScore := range newStudent.Scores {
		// Update existing scores, based on length of data
		if idx < len(existingStudent.Scores) {
			// Update existing score
			existingStudent.Scores[idx].AssignmentTitle = newScore.AssignmentTitle
			existingStudent.Scores[idx].Description = newScore.Description
			existingStudent.Scores[idx].Score = newScore.Score
			// Save the update to DB transaction
			err = tx.Save(&existingStudent.Scores[idx]).Error
			if err != nil {
				// If there's error, cancel the transaction
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update a score"})
				return
			}
		} else {
			// If the length of score is longer than the existing student's score, add new data
			// Get the student ID for inserting new score
			newScore.StudentID = existingStudent.ID
			existingStudent.Scores = append(existingStudent.Scores, newScore)
			// Create new score through DB transaction
			err = tx.Create(&existingStudent.Scores[idx]).Error
			if err != nil {
				// If there's error, cancel the transaction
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to create new score"})
				return
			}
		}
	}

	// Update the student data
	err = tx.Save(&existingStudent).Error
	if err != nil {
		// If there's error, cancel the transaction
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update student data"})
		return
	}

	// Commit the transaction
	tx.Commit()

	// Send response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    existingStudent,
	})
}

// Delete student by ID
func DeleteStudentByID(c *gin.Context) {
	// Get database
	db := database.GetDB()
	// Retrieve student ID to be deleted from param
	studentID := c.Param("id")
	// Delete student by ID
	student := models.Student{}
	err := db.Where("id = ?", studentID).Delete(&student).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    student,
	})
}

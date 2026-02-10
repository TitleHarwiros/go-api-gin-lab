package handlers

import (
	"go-api-gin/models"
	"go-api-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	svc services.StudentService
}

func NewStudentHandler(svc services.StudentService) *StudentHandler {
	return &StudentHandler{svc: svc}
}

// Challenge 1: Update Student
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	// Challenge 3: Validation (ใช้ ShouldBindJSON คู่กับ struct tags)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if err := h.svc.UpdateStudent(id, student); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// คืนข้อมูลที่อัปเดตแล้ว (ตั้งค่า ID ให้ตรงกับ URL)
	student.ID = id
	c.JSON(http.StatusOK, student)
}

// Challenge 2: Delete Student
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.DeleteStudent(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, _ := h.svc.GetStudents()
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) PostStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.svc.CreateStudent(student)
	c.JSON(http.StatusCreated, student)
}

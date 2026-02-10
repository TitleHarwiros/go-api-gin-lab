package repositories

import (
	"errors"
	"go-api-gin/models"
	"gorm.io/gorm"
)

type StudentRepository interface {
	GetAll() ([]models.Student, error)
	GetByID(id string) (models.Student, error)
	Create(student models.Student) error
	Update(id string, student models.Student) error
	Delete(id string) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) GetAll() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Find(&students).Error
	return students, err
}

func (r *studentRepository) GetByID(id string) (models.Student, error) {
	var student models.Student
	err := r.db.First(&student, "id = ?", id).Error
	return student, err
}

func (r *studentRepository) Create(student models.Student) error {
	return r.db.Create(&student).Error
}

func (r *studentRepository) Update(id string, updatedData models.Student) error {
	// ค้นหาว่ามีอยู่จริงไหม (Challenge 1 behavior)
	var student models.Student
	if err := r.db.First(&student, "id = ?", id).Error; err != nil {
		return err
	}
	// อัปเดตข้อมูล
	return r.db.Model(&student).Updates(updatedData).Error
}

func (r *studentRepository) Delete(id string) error {
	result := r.db.Delete(&models.Student{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("student not found")
	}
	return nil
}

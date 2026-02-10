package services

import (
	"go-api-gin/models"
	"go-api-gin/repositories"
)

type StudentService interface {
	GetStudents() ([]models.Student, error)
	GetStudent(id string) (models.Student, error)
	CreateStudent(student models.Student) error
	UpdateStudent(id string, student models.Student) error
	DeleteStudent(id string) error
}

type studentService struct {
	repo repositories.StudentRepository
}

func NewStudentService(repo repositories.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) GetStudents() ([]models.Student, error) {
	return s.repo.GetAll()
}

func (s *studentService) GetStudent(id string) (models.Student, error) {
	return s.repo.GetByID(id)
}

func (s *studentService) CreateStudent(student models.Student) error {
	return s.repo.Create(student)
}

func (s *studentService) UpdateStudent(id string, student models.Student) error {
	return s.repo.Update(id, student)
}

func (s *studentService) DeleteStudent(id string) error {
	return s.repo.Delete(id)
}

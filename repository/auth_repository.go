package repository

import (
	"database/sql"
	"log"

	authModel "github.com/ReygaFitra/auth-jwt/model"
)

type AuthRepo interface {
	Register(newStudent *authModel.Credential) string
	Login(student *authModel.Credential) string
}

type authRepo struct {
	db *sql.DB
}

func (r *authRepo) Register(newStudent *authModel.Credential) string {
	query := "INSERT INTO credentials (user_name, password) VALUES ($1, $2)"
	_, err := r.db.Exec(query, newStudent.Username, newStudent.Password)

	if err != nil {
		log.Println(err)
		return "failed to create student!"
	}

	return "created successfully"
}

func (r *authRepo) Login(student *authModel.Credential) string {
	query := "INSERT INTO credentials (user_name, password) VALUES ($1, $2)"
	_, err := r.db.Exec(query, student.Username, student.Password)

	if err != nil {
		log.Println(err)
		return "failed to login!"
	}

	return "created successfully"
}

func NewStudentRepo(db *sql.DB) AuthRepo {
	repo := new(authRepo)
	repo.db = db

	return repo
}
package authRepository

import (
	"database/sql"
	"log"

	authDatabase "github.com/ReygaFitra/auth-jwt/database"
	authModel "github.com/ReygaFitra/auth-jwt/model"
)

type AuthRepo interface {
	Register(newStudent *authModel.Credential) string
	Login() string
}

type authRepo struct {
	db *sql.DB
}

func (r *authRepo) Register(newStudent *authModel.Credential) string {
	db, _ := authDatabase.ConnectDB()
	defer db.Close()

	query := "INSERT INTO credentials (user_name, password) VALUES ($1, $2)"
	_, err := r.db.Exec(query, newStudent.Username, newStudent.Password)

	if err != nil {
		log.Println(err)
		return "failed to create student!"
	}

	return "created successfully"
}

func (r *authRepo) Login() string {
	db, _ := authDatabase.ConnectDB()
	defer db.Close()

	var users []authModel.Credential

	query := "SELECT user_name FROM credentials"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
		return "failed to login!"
	}

	defer rows.Close()

	for rows.Next() {
		var user authModel.Credential
		err := rows.Scan(&user.Username)
		if err != nil {
			log.Print(err)
		}

		users = append(users, user)
	}
	
	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(users) == 0 {
		return "no data"
	}

	return "Login successfully"
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	repo := new(authRepo)
	repo.db = db

	return repo
}
package authUsecase

import (
	authModel "github.com/ReygaFitra/auth-jwt/model"
	authRepository "github.com/ReygaFitra/auth-jwt/repository"
)

type AuthUsecase interface {
	SignUp(newStudent *authModel.Credential) string
	SignIn(student *authModel.Credential) string
}

type authUsecase struct {
	authRepo authRepository.AuthRepo
}

func (u *authUsecase) SignUp(newStudent *authModel.Credential) string {
	return u.authRepo.Register(newStudent)
}

func (u *authUsecase) SignIn(student *authModel.Credential) string {
	return u.authRepo.Login(student)
}

func NewAuthUsecase(authRepo authRepository.AuthRepo) AuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}


package usecase

import (
	"abix360/src/dao"
	"abix360/src/domain"
	"abix360/src/view/dto"
)

type PasswordListUseCase struct{}

func (useCase *PasswordListUseCase) Execute() (passwordList []dto.PasswordDTO, err error) {
	passwordList = []dto.PasswordDTO{}

	list, err := domain.PasswordList(dao.NewPasswordDAO())
	if err != nil {
		return passwordList, err
	}

	for _, password := range list {
		passwordList = append(passwordList, dto.PasswordDTO{
			Id:       password.Id(),
			Name:     password.Name(),
			User:     password.User(),
			Password: password.Password(),
			Url:      password.Url(),
		})
	}
	return passwordList, err
}

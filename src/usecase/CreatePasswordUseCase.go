package usecase

import (
	"abix360/src/dao"
	"abix360/src/domain"
	"abix360/src/view/dto"
)

type CreatePasswordUseCase struct{}

func (useCase *CreatePasswordUseCase) Execute(createPassword dto.CreatePasswordDTO) (passwordDto dto.PasswordDTO, err error) {

	var repository domain.RepositoryPassword = dao.NewPasswordDAO()

	newPassword := domain.NewPassword()
	newPassword.WithName(createPassword.Name)
	newPassword.WithUser(createPassword.User)
	newPassword.WithPassword(createPassword.Password)
	newPassword.WithUrl(createPassword.Url)
	newPassword.WithRepository(repository)

	err = newPassword.Save()
	if err != nil {
		return passwordDto, err
	}

	passwordDto = dto.PasswordDTO{
		Id:       newPassword.Id(),
		Name:     newPassword.Name(),
		User:     newPassword.User(),
		Password: newPassword.Password(),
		Url:      newPassword.Url(),
	}

	return passwordDto, nil
}

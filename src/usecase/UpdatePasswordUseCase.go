package usecase

import (
	"abix360/src/dao"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type UpdatePasswordUseCase struct{}

func (useCase *UpdatePasswordUseCase) Execute(passwordDto dto.PasswordDTO) (dto.PasswordDTO, error) {
	var respPassword dto.PasswordDTO = dto.PasswordDTO{}
	var repository domain.RepositoryPassword = dao.NewPasswordDAO()

	password, err := domain.FindPasswordById(passwordDto.Id, repository)
	if err != nil {
		return respPassword, err
	}

	if !password.Exists() {
		return respPassword, errors.New("no existe el elemento")
	}

	password.WithName(passwordDto.Name)
	password.WithPassword(passwordDto.Password)
	password.WithUrl(passwordDto.Url)
	password.WithUser(passwordDto.User)
	password.WithRepository(repository)

	err = password.Update()
	if err != nil {
		return respPassword, err
	}

	respPassword.Id = password.Id()
	respPassword.Name = password.Name()
	respPassword.User = password.User()
	respPassword.Url = password.Url()
	respPassword.Password = password.Password()

	return respPassword, err
}

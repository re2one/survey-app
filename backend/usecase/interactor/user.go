package interactor

import (
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"backend/common"
	"backend/model"
	"backend/model/response"
	"backend/usecase/presenter"
	"backend/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
	UserPresenter  presenter.UserPresenter
	authorizer     common.Authorizer
	authenticator  common.Authenticator
}

type UserInteractor interface {
	Get(*model.User) (*response.UserResponse, error)
	GetAll() ([]*response.SmolUserResponse, error)
	Post(*model.User, string) (*response.UserResponse, error)
	Refresh(*http.Request) (*response.UserResponse, error)
}

func NewUserInteractor(
	ur repository.UserRepository,
	rr repository.RoleRepository,
	p presenter.UserPresenter,
	authorizer common.Authorizer,
	authenticator common.Authenticator,
) UserInteractor {
	return &userInteractor{UserRepository: ur, RoleRepository: rr, UserPresenter: p, authorizer: authorizer, authenticator: authenticator}
}

func (us *userInteractor) Get(u *model.User) (*response.UserResponse, error) {

	var r *model.Role
	var err error
	var res *response.UserResponse
	var pw = u.Password

	u, err = us.UserRepository.Get(u)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user from user repo.")
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(u.Salt+pw))
	if err != nil {
		log.Error().Err(err).Msg("Failed to compare hash and salted pw, trying plain pw.")
	}

	r, err = us.RoleRepository.Get(u)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get role from role repo.")
		return nil, err
	}

	// hhhhh
	res, err = us.UserPresenter.LoginResponse(u, r)
	if err != nil {

		log.Error().Err(err).Msg("Failed to build login response.")
		return nil, err
	}
	return res, nil
}

func (us *userInteractor) Post(u *model.User, role string) (*response.UserResponse, error) {
	var err error
	var res *response.UserResponse
	var r *model.Role

	u, err = us.UserRepository.Post(u)
	if err != nil {
		return nil, err
	}

	r, err = us.RoleRepository.Post(u, role)
	if err != nil {
		return nil, err
	}

	// hhhhh
	res, err = us.UserPresenter.SignupResponse(u, r)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (us *userInteractor) Refresh(request *http.Request) (*response.UserResponse, error) {
	claims, err := us.authorizer.UnwrapClaims(request)
	if err != nil {
		return nil, err
	}

	user := model.User{Name: claims.Name, Email: claims.Email}
	role := model.Role{Role: claims.Role}
	token, err := us.authenticator.CreateToken(&user, &role)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update token.")
		return nil, err
	}
	return token, nil
}

func (us *userInteractor) GetAll() ([]*response.SmolUserResponse, error) {
	var users []*model.User
	users, err := us.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	smolUsers := make([]*response.SmolUserResponse, 0, len(users))
	for _, user := range users {
		smolUsers = append(smolUsers, &response.SmolUserResponse{Email: user.Email, Username: user.Name, Id: strconv.Itoa(int(user.ID))})
	}
	return smolUsers, nil
}

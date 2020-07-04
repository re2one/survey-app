package controller

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo"
	"survey-app-backend/model"
	"survey-app-backend/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetAll(c Context) error
	Get(c Context) error
	Add(c echo.Context) error
	Update(c Context) error
	Delete(c Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetAll(c Context) error {
	var u []*model.User

	u, err := uc.userInteractor.GetAll(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) Get(c Context) error {
	var u *model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) Add(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	  }
	fmt.Printf("%+v\n", u)
	u, err := uc.userInteractor.Add(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) Update(c Context) error {
	var u *model.User

	u, err := uc.userInteractor.Update(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) Delete(c Context) error {
	var u *model.User

	u, err := uc.userInteractor.Delete(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
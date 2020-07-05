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
	GetAll(c echo.Context) error
	Get(c echo.Context) error
	Add(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetAll(c echo.Context) error {
	var u []*model.User

	u, err := uc.userInteractor.GetAll(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) Get(c echo.Context) error {
	u := model.User{Email: c.QueryParam("email")}

	result, err := uc.userInteractor.Get(&u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
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

func (uc *userController) Update(c echo.Context) error {
	var u *model.User

	u, err := uc.userInteractor.Update(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) Delete(c echo.Context) error {
	u := model.User{Email: c.QueryParam("email")}

	result, err := uc.userInteractor.Delete(&u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
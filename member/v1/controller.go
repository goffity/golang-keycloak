package v1

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type ControllerRegister interface {
	GetEventByID(context echo.Context) error
}

type controllerRegister struct {
	service ServiceRegister
}

// GetEventByID implements ControllerRegister
func (c controllerRegister) GetEventByID(context echo.Context) error {
	idParam := context.Param("id")
	if idParam == "" {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, echo.ErrBadRequest.Message)
	}
	id, _ := strconv.Atoi(idParam)

	member, err := c.service.GetEventByID(id)
	if err != nil {
		return echo.NewHTTPError(echo.ErrNotFound.Code, echo.ErrNotFound.Message)
	}

	return context.JSON(200, member)
}

func NewRegisterController(service ServiceRegister) ControllerRegister {
	return &controllerRegister{service: service}
}

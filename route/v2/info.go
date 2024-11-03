package v2

import (
	"net/http"

	"github.com/dappster-io/DappsterOS-AppManagement/codegen"
	"github.com/dappster-io/DappsterOS-AppManagement/pkg/docker"
	"github.com/dappster-io/DappsterOS-Common/utils"
	"github.com/labstack/echo/v4"
)

func (a *AppManagement) Info(ctx echo.Context) error {
	architecture, err := docker.CurrentArchitecture()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Message: utils.Ptr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusOK, codegen.InfoOK{
		Architecture: utils.Ptr(architecture),
	})
}

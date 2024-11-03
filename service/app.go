package service

import (
	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
	"github.com/dappster-io/DappsterOS-AppManagement/codegen"
	"github.com/dappster-io/DappsterOS-AppManagement/common"
	"github.com/dappster-io/DappsterOS-Common/utils/logger"
)

type App types.ServiceConfig

func (a *App) StoreInfo() (codegen.AppStoreInfo, error) {
	var storeInfo codegen.AppStoreInfo

	ex, ok := a.Extensions[common.ComposeExtensionNameXDappsterOS]
	if !ok {
		logger.Error("extension `x-dappsteros` not found")
		// return storeInfo, ErrComposeExtensionNameXDappsterOSNotFound
	}

	// add image to store info for check stable version function.
	storeInfo.Image = a.Image

	if err := loader.Transform(ex, &storeInfo); err != nil {
		return storeInfo, err
	}

	return storeInfo, nil
}

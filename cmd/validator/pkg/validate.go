package pkg

import (
	"github.com/compose-spec/compose-go/loader"
	"github.com/dappsteros-io/DappsterOS-AppManagement/codegen"
	"github.com/dappsteros-io/DappsterOS-AppManagement/common"
	"github.com/dappsteros-io/DappsterOS-AppManagement/service"
)

func VaildDockerCompose(yaml []byte) (err error) {
	err = nil
	// recover
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	docker, err := service.NewComposeAppFromYAML(yaml, false, false)

	ex, ok := docker.Extensions[common.ComposeExtensionNameXDappsterOS]
	if !ok {
		return service.ErrComposeExtensionNameXDappsterOSNotFound
	}

	var storeInfo codegen.ComposeAppStoreInfo
	if err = loader.Transform(ex, &storeInfo); err != nil {
		return
	}

	return
}

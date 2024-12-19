package v2_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/compose-spec/compose-go/types"
	"github.com/dappsteros-io/DappsterOS-AppManagement/codegen"
	"github.com/dappsteros-io/DappsterOS-AppManagement/common"
	v2 "github.com/dappsteros-io/DappsterOS-AppManagement/route/v2"
	"github.com/dappsteros-io/DappsterOS-AppManagement/service"
	"github.com/dappsteros-io/DappsterOS-Common/utils/logger"
)

func TestFilterCatalogByCategory(t *testing.T) {
	logger.LogInitConsoleOnly()

	catalog := map[string]*service.ComposeApp{}

	filteredCatalog := v2.FilterCatalogByCategory(catalog, "test")
	assert.Equal(t, len(filteredCatalog), 0)

	catalog["test"] = &service.ComposeApp{
		Extensions: map[string]interface{}{
			common.ComposeExtensionNameXDappsterOS: map[string]interface{}{
				"category": "test",
			},
		},
	}

	filteredCatalog = v2.FilterCatalogByCategory(catalog, "test")
	assert.Equal(t, len(filteredCatalog), 1)

	catalog["test2"] = &service.ComposeApp{
		Extensions: map[string]interface{}{
			common.ComposeExtensionNameXDappsterOS: map[string]interface{}{
				"category": "test2",
			},
		},
	}

	filteredCatalog = v2.FilterCatalogByCategory(catalog, "test")
	assert.Equal(t, len(filteredCatalog), 1)
}

func TestFilterCatalogByAuthorType(t *testing.T) {
	logger.LogInitConsoleOnly()

	catalog := map[string]*service.ComposeApp{}

	filteredCatalog := v2.FilterCatalogByAuthorType(catalog, "test")
	assert.Equal(t, len(filteredCatalog), 0)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.ByDappsteros)
	assert.Equal(t, len(filteredCatalog), 0)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.Official)
	assert.Equal(t, len(filteredCatalog), 0)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.Community)
	assert.Equal(t, len(filteredCatalog), 0)

	catalog["test"] = &service.ComposeApp{
		Extensions: map[string]interface{}{
			common.ComposeExtensionNameXDappsterOS: map[string]interface{}{
				"author": common.ComposeAppAuthorDappsterOSTeam,
			},
		},
	}

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, "test")
	assert.Equal(t, len(filteredCatalog), 0)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.ByDappsteros)
	assert.Equal(t, len(filteredCatalog), 1)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.Official)
	assert.Equal(t, len(filteredCatalog), 0)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.Community)
	assert.Equal(t, len(filteredCatalog), 0)

	catalog["test2"] = &service.ComposeApp{
		Extensions: map[string]interface{}{
			common.ComposeExtensionNameXDappsterOS: map[string]interface{}{
				"author":    "test2",
				"developer": "test2",
			},
		},
	}

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, "test")
	assert.Equal(t, len(filteredCatalog), 0)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.ByDappsteros)
	assert.Equal(t, len(filteredCatalog), 1)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.Official)
	assert.Equal(t, len(filteredCatalog), 1)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.Community)
	assert.Equal(t, len(filteredCatalog), 0)

	catalog["test3"] = &service.ComposeApp{
		Extensions: map[string]interface{}{
			common.ComposeExtensionNameXDappsterOS: map[string]interface{}{
				"author":    "test3",
				"developer": "syncthing",
			},
		},
	}

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, "test")
	assert.Equal(t, len(filteredCatalog), 0)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.ByDappsteros)
	assert.Equal(t, len(filteredCatalog), 1)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.Official)
	assert.Equal(t, len(filteredCatalog), 1)

	filteredCatalog = v2.FilterCatalogByAuthorType(catalog, codegen.Community)
	assert.Equal(t, len(filteredCatalog), 1)
}

func TestFilterCatalogByAppStoreID(t *testing.T) {
	logger.LogInitConsoleOnly()

	catalog := map[string]*service.ComposeApp{}

	filteredCatalog := v2.FilterCatalogByAppStoreID(catalog, []string{"test"})
	assert.Equal(t, len(filteredCatalog), 0)

	catalog["test"] = &service.ComposeApp{
		Extensions: map[string]interface{}{
			common.ComposeExtensionNameXDappsterOS: map[string]interface{}{
				"main": "test",
			},
		},
		Services: []types.ServiceConfig{
			{
				Name: "test",
				Extensions: map[string]interface{}{
					common.ComposeExtensionNameXDappsterOS: map[string]interface{}{
						"app_store_id": "test",
					},
				},
			},
		},
	}

	filteredCatalog = v2.FilterCatalogByAppStoreID(catalog, []string{"test"})
	assert.Equal(t, len(filteredCatalog), 1)

	catalog["test2"] = &service.ComposeApp{
		Extensions: map[string]interface{}{
			common.ComposeExtensionNameXDappsterOS: map[string]interface{}{
				"main": "test2",
			},
		},
		Services: []types.ServiceConfig{
			{
				Name: "test2",
				Extensions: map[string]interface{}{
					common.ComposeExtensionNameXDappsterOS: map[string]interface{}{
						"app_store_id": "test2",
					},
				},
			},
		},
	}

	filteredCatalog = v2.FilterCatalogByAppStoreID(catalog, []string{"test"})
	assert.Equal(t, len(filteredCatalog), 1)

	filteredCatalog = v2.FilterCatalogByAppStoreID(catalog, []string{"test", "test2"})
	assert.Equal(t, len(filteredCatalog), 2)

	filteredCatalog = v2.FilterCatalogByAppStoreID(catalog, []string{"test", "test2", "test3"})
	assert.Equal(t, len(filteredCatalog), 2)

	filteredCatalog = v2.FilterCatalogByAppStoreID(catalog, []string{"test1", "test2", "test3"})
	assert.Equal(t, len(filteredCatalog), 1)
}

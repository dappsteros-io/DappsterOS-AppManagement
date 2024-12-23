package common

const (
	AppManagementServiceName = "app-management"
	AppManagementVersion     = "0.4.5"

	AppsDirectoryName = "Apps"

	ComposeAppAuthorDappsterOSTeam = "DappsterOS Team"

	ComposeExtensionNameXDappsterOS            = "x-dappsteros"
	ComposeExtensionPropertyNameStoreAppID     = "store_app_id"
	ComposeExtensionPropertyNameTitle          = "title"
	ComposeExtensionPropertyNameIsUncontrolled = "is_uncontrolled"

	ComposeYAMLFileName = "docker-compose.yml"

	ContainerLabelV1AppStoreID = "io.dappsteros.v1.app.store.id"

	DefaultCategoryFont = "grid"
	DefaultLanguage     = "en_us"
	DefaultPassword     = "dappsteros"
	DefaultPGID         = "1000"
	DefaultPUID         = "1000"
	DefaultUserName     = "admin"

	Localhost           = "127.0.0.1"
	MIMEApplicationYAML = "application/yaml"

	CategoryListFileName  = "category-list.json"
	RecommendListFileName = "recommend-list.json"
)

// the tags can add more. like "latest", "stable", "edge", "beta", "alpha"
var NeedCheckDigestTags = []string{"latest"}

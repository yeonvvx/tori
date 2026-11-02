package constants

import (
	"tori/internal/util"
	"time"
)

const (
	// Brand — swap these when rebranding the fork.
	AppName    = "Tori"
	AppRepoURL = "https://github.com/yeonix/tori"

	Version              = "3.10.4"
	VersionName          = "Saisei"
	GcTime               = time.Minute * 30
	ConfigFileName       = "config.toml"
	MalClientId          = "51cb4294feb400f3ddc66a30f9b9a00f"
	DiscordApplicationId = "1547509318121824338"
	AnilistApiUrl        = "https://graphql.anilist.co"
)

const (
	ToriRoomsApiUrl   = "https://tori.app/api/rooms"
	ToriRoomsApiWsUrl = "wss://tori.app/api/rooms"
	ToriRoomsVersion  = "1.0.0"
)

var DefaultExtensionMarketplaceURL = util.Decode("aHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL3llb25peC90b3JpLWV4dGVuc2lvbnMvcmVmcy9oZWFkcy9tYWluL21hcmtldHBsYWNlLmpzb24=")
var AnnouncementURL = util.Decode("aHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL3llb25peC9oaWJpa2UvcmVmcy9oZWFkcy9tYWluL3B1YmxpYy9hbm5vdW5jZW1lbnRzLmpzb24=")
var InternalMetadataURL = util.Decode("aHR0cHM6Ly9hbmltZS5jbGFwLmluZw==")

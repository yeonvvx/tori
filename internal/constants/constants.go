package constants

import (
	"tori/internal/util"
	"time"
)

const (
	// Brand — swap these when rebranding the fork.
	AppName    = "Tori"
	AppRepoURL = "https://github.com/yeonvvx/tori"

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

var DefaultExtensionMarketplaceURL = util.Decode("aHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL3llb252dngvdG9yaS1leHRlbnNpb25zL3JlZnMvaGVhZHMvbWFpbi9tYXJrZXRwbGFjZS5qc29u")
var AnnouncementURL = util.Decode("aHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL3llb252dngvaGliaWtlL3JlZnMvaGVhZHMvbWFpbi9wdWJsaWMvYW5ub3VuY2VtZW50cy5qc29u")
var InternalMetadataURL = util.Decode("aHR0cHM6Ly9hbmltZS5jbGFwLmluZw==")

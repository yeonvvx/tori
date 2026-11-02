package core

import (
	"tori/internal/extension"
	"tori/internal/extension_repo"
	manga_providers "tori/internal/manga/providers"

	"github.com/rs/zerolog"
)

func LoadCustomSourceExtensions(extensionRepository *extension_repo.Repository) {
	extensionRepository.LoadOnlyWrapper([]extension.Type{extension.TypeCustomSource}, func() {
		extensionRepository.ReloadExternalExtensions()
	})
}

func LoadExtensions(extensionRepository *extension_repo.Repository, logger *zerolog.Logger, config *Config) {
	// Load built-in extensions
	extensionRepository.ReloadBuiltInExtension(extension.Extension{
		ID:          manga_providers.LocalProvider,
		Name:        "Local",
		Version:     "",
		ManifestURI: "builtin",
		Language:    extension.LanguageGo,
		Type:        extension.TypeMangaProvider,
		Author:      "Tori",
		Lang:        "multi",
		Icon:        "https://raw.githubusercontent.com/yeonvvx/hibike/main/icons/local-manga.png",
	}, manga_providers.NewLocal(config.Manga.LocalDir, logger))

	// Load external extensions
	//extensionRepository.ReloadExternalExtensions()
	extensionRepository.LoadOnlyWrapper([]extension.Type{extension.TypeMangaProvider, extension.TypeOnlinestreamProvider, extension.TypeAnimeTorrentProvider, extension.TypePlugin}, func() {
		extensionRepository.ReloadExternalExtensions()
	})
}

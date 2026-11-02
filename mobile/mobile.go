package mobile

import (
	"embed"
	"os"
	"tori/internal/core"
	"tori/internal/handlers"
)

//go:embed all:web
var WebFS embed.FS

func StartServer(dataDir string, cacheDir string, port int) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
			}
		}()

		_ = os.Setenv("TORI_DATA_DIR", dataDir)
		_ = os.Setenv("TORI_WORKING_DIR", dataDir)

		flags := core.ToriFlags{
			DataDir:          dataDir,
			Port:             port,
			IsDesktopSidecar: false,
		}

		app := core.NewApp(&core.ConfigOptions{
			Flags: flags,
		}, nil)

		app.InitLogging(false)

		echoApp := core.NewEchoApp(app, &WebFS)

		handlers.InitRoutes(app, echoApp)

		core.RunEchoServer(app, echoApp)
	}()
}

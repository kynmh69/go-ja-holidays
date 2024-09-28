package template

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/kynmh69/go-ja-holidays/logging"
	"path/filepath"
)

func Render(templateGlob string) multitemplate.Renderer {
	logger := logging.GetLogger()
	logger.Debug("Glob:", templateGlob)
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templateGlob)
	if err != nil {
		logger.Panicln(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	r.AddFromFiles("view", layouts...)
	return r
}

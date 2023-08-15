package artifacts

import (
	"sw-gittycat-server/modules/application"

	"github.com/passon-engineering/sw-go-utility-lib/file"
)

func DeleteAllArtifacts(app *application.Application) error {
	ignoreFiles := map[string]bool{
		".gitignore": true,
		// add more files to ignore here
	}

	err := file.DeleteAllExceptIgnored(app.ServerPath+"/artifacts/", ignoreFiles)
	if err != nil {
		return err
	} else {
		return nil
	}
}

package storage

import (
	"errors"

	"github.com/centrifuge/go-centrifuge/centrifuge/bootstrapper"
	"github.com/centrifuge/go-centrifuge/centrifuge/config"
)

type Bootstrapper struct {
}

func (*Bootstrapper) Bootstrap(context map[string]interface{}) error {
	if _, ok := context[bootstrapper.BootstrappedConfig]; ok {
		levelDB := NewLevelDBStorage(config.Config.GetStoragePath())
		context[bootstrapper.BootstrappedLevelDb] = levelDB
		return nil
	}
	return errors.New("could not initialize levelDB")
}
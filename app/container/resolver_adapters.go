package container

import (
	"fmt"

	"github.com/storybuilder/storybuilder/app/config"
	"github.com/storybuilder/storybuilder/externals/adapters"
)

var resolvedAdapters Adapters

// resolveAdapters resolves all adapters.
func resolveAdapters(cfg *config.Config) Adapters {
	resolveDBAdapter(cfg.DBConfig)
	resolveDBTransactionAdapter()
	resolveLogAdapter(cfg.LogConfig)
	resolveValidatorAdapter()

	return resolvedAdapters
}

// resolveDBAdapter resolves the database adapter.
func resolveDBAdapter(cfg config.DBConfig) {
	db, err := adapters.NewMySQLAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	resolvedAdapters.DBAdapter = db
}

// resolveDBTransactionAdapter resolves the database transaction adapter.
func resolveDBTransactionAdapter() {
	tx := adapters.NewMySQLTxAdapter(resolvedAdapters.DBAdapter)

	resolvedAdapters.DBTxAdapter = tx
}

// resolveLogAdapter resolves the logging adapter.
func resolveLogAdapter(cfg config.LogConfig) {
	la, err := adapters.NewLogAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	resolvedAdapters.LogAdapter = la
}

// resolveValidatorAdapter resolves the validation adapter.
func resolveValidatorAdapter() {
	v, err := adapters.NewValidatorAdapter()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	resolvedAdapters.ValidatorAdapter = v
}

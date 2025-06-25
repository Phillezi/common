module github.com/Phillezi/common/scripts/glint

go 1.24.4

replace github.com/Phillezi/common/logging/zap => ../../logging/zap

replace github.com/Phillezi/common/config => ../../config

replace github.com/Phillezi/common/utils => ../../utils

require (
	github.com/Phillezi/common/config v0.0.0-00010101000000-000000000000
	github.com/Phillezi/common/logging/zap v0.0.0-00010101000000-000000000000
	github.com/Phillezi/common/utils v0.0.1
	github.com/fzipp/gocyclo v0.6.0
	github.com/gordonklaus/ineffassign v0.1.0
	github.com/spf13/cobra v1.9.1
	github.com/spf13/viper v1.20.1
	go.uber.org/zap v1.27.0
	golang.org/x/tools v0.30.0
	honnef.co/go/tools v0.6.1
)

require (
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/sagikazarmark/locafero v0.7.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.12.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20231108232855-2478ac86f678 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

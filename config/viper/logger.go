package viperconf

import lggr "github.com/Phillezi/common/config/logger"

var logger lggr.Logger = lggr.NoopLogger{}

// SetLogger allows consumers to inject a custom logger.
func SetLogger(l lggr.Logger) {
	if l != nil {
		logger = l
	} else {
		logger = lggr.NoopLogger{}
	}
}

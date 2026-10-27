package dm8

import "sync"

// Connector initialization may run while existing connections and background
// flushers read the process-wide diagnostic settings. Keep their existing
// last-configured behavior without serializing database operations.
var diagnosticConfigMu sync.RWMutex

type diagnosticConfig struct {
	LogLevel             int
	LogDir               string
	LogBufferSize        int
	LogFlushFreq         int
	LogFlushQueueSize    int
	StatEnable           bool
	StatDir              string
	StatFlushFreq        int
	StatHighFreqSqlCount int
	StatSlowSqlCount     int
	StatSqlMaxCount      int
	StatSqlRemoveMode    int
}

func readDiagnosticConfig() diagnosticConfig {
	diagnosticConfigMu.RLock()
	defer diagnosticConfigMu.RUnlock()
	return diagnosticConfig{
		LogLevel:             LogLevel,
		LogDir:               LogDir,
		LogBufferSize:        LogBufferSize,
		LogFlushFreq:         LogFlushFreq,
		LogFlushQueueSize:    LogFlushQueueSize,
		StatEnable:           StatEnable,
		StatDir:              StatDir,
		StatFlushFreq:        StatFlushFreq,
		StatHighFreqSqlCount: StatHighFreqSqlCount,
		StatSlowSqlCount:     StatSlowSqlCount,
		StatSqlMaxCount:      StatSqlMaxCount,
		StatSqlRemoveMode:    StatSqlRemoveMode,
	}
}

func (c *DmConnector) publishDiagnosticConfig() {
	diagnosticConfigMu.Lock()
	defer diagnosticConfigMu.Unlock()
	LogLevel = c.logLevel
	LogDir = c.logDir
	LogBufferSize = c.logBufferSize
	LogFlushFreq = c.logFlushFreq
	LogFlushQueueSize = c.logFlushQueueSize
	StatEnable = c.statEnable
	StatDir = c.statDir
	StatFlushFreq = c.statFlushFreq
	StatHighFreqSqlCount = c.statHighFreqSqlCount
	StatSlowSqlCount = c.statSlowSqlCount
	StatSqlMaxCount = c.statSqlMaxCount
}

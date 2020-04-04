package mixers

import (
	"sync"

	logging "github.com/codemodify/systemkit-logging"
)

// BufferConfig -
type BufferConfig struct {
	MaxLogEntries int
}

type bufferedLogger struct {
	logger          logging.Logger
	config          BufferConfig
	logEntries      []logging.LogEntry
	logEntriesMutex *sync.RWMutex
}

// NewBuffered - Buffers the log entries up to MAX-ENTRIES
func NewBuffered(logger logging.Logger, config BufferConfig) logging.Logger {
	return &bufferedLogger{
		logger:          logger,
		config:          config,
		logEntries:      []logging.LogEntry{},
		logEntriesMutex: &sync.RWMutex{},
	}
}

func (thisRef *bufferedLogger) Log(logEntry logging.LogEntry) logging.LogEntry {
	thisRef.logEntriesMutex.Lock()
	defer thisRef.logEntriesMutex.Unlock()

	thisRef.logEntries = append(thisRef.logEntries, logEntry)

	if len(thisRef.logEntries) > thisRef.config.MaxLogEntries {
		for _, logEntry := range thisRef.logEntries {
			thisRef.logger.Log(logEntry)
		}

		thisRef.logEntries = []logging.LogEntry{}
	}

	return logging.LogEntry{}
}

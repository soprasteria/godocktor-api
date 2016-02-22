package logs

import (
	"io"
	"log"
)

var (
	// Debug logger for events to inform the developer
	Debug *log.Logger
	// Info logger for events to inform the user
	Info *log.Logger
	// Warning logger for events that does not stop the program but should be fixed
	Warning *log.Logger
	// Error logger for events that stops the program
	Error *log.Logger
)

// InitLog inits the logs for Guts application
func InitLog(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer, flag int) {
	Debug = log.New(traceHandle, "[GO-DOCKTOR-API] [DEBUG] ", flag)
	Info = log.New(infoHandle, "[GO-DOCKTOR-API] [INFO] ", flag)
	Warning = log.New(warningHandle, "[GO-DOCKTOR-API] [WARN] ", flag)
	Error = log.New(errorHandle, "[GO-DOCKTOR-API] [ERROR] ", flag)
}

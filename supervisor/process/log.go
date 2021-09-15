package process

import "github.com/osgochina/dmicro/supervisor/proclog"

// 创建标准输出日志
func (that *Process) createStdoutLogger() proclog.Logger {
	logFile := that.GetStdoutLogfile()
	maxBytes := int64(that.entry.StdoutLogFileMaxBytes(50 * 1024 * 1024))
	backups := that.entry.StdoutLogFileBackups(10)

	props := make(map[string]string)
	syslogFacility := that.entry.Extend("syslog_facility", "")
	syslogTag := that.entry.Extend("syslog_tag", "")
	syslogPriority := that.entry.Extend("syslog_stdout_priority", "")

	if len(syslogFacility) > 0 {
		props["syslog_facility"] = syslogFacility
	}
	if len(syslogTag) > 0 {
		props["syslog_tag"] = syslogTag
	}
	if len(syslogPriority) > 0 {
		props["syslog_priority"] = syslogPriority
	}
	return proclog.NewLogger(that.GetName(), logFile, proclog.NewNullLocker(), maxBytes, backups, props)
}

// 创建标准错误日志
func (that *Process) createStderrLogger() proclog.Logger {
	logFile := that.GetStderrLogfile()
	maxBytes := int64(that.entry.StderrLogFileMaxBytes(50 * 1024 * 1024))
	backups := that.entry.StderrLogFileBackups(10)

	props := make(map[string]string)
	syslogFacility := that.entry.Extend("syslog_facility", "")
	syslogTag := that.entry.Extend("syslog_tag", "")
	syslogPriority := that.entry.Extend("syslog_stdout_priority", "")

	if len(syslogFacility) > 0 {
		props["syslog_facility"] = syslogFacility
	}
	if len(syslogTag) > 0 {
		props["syslog_tag"] = syslogTag
	}
	if len(syslogPriority) > 0 {
		props["syslog_priority"] = syslogPriority
	}

	return proclog.NewLogger(that.GetName(), logFile, proclog.NewNullLocker(), maxBytes, backups, props)
}

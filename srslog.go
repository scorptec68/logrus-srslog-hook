package logrus_srslog_hook


import (
	"fmt"
	"os"
	"github.com/RackSec/srslog"
	"github.com/sirupsen/logrus"
)

// SrslogHook to send logs via srslog.
type SrslogHook struct {
	Writer        *srslog.Writer
}

// Creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewSrslogHook("udp", "localhost:514", syslog.LOG_DEBUG, "")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewSrslogHook(network, raddr string, priority srslog.Priority, tag string) (*SrslogHook, error) {
	w, err := srslog.Dial(network, raddr, priority, tag)
	return &SrslogHook{w}, err
}

func (hook *SrslogHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	switch entry.Level {
	case logrus.PanicLevel:
		return hook.Writer.Crit(line)
	case logrus.FatalLevel:
		return hook.Writer.Crit(line)
	case logrus.ErrorLevel:
		return hook.Writer.Err(line)
	case logrus.WarnLevel:
		return hook.Writer.Warning(line)
	case logrus.InfoLevel:
		return hook.Writer.Info(line)
	case logrus.DebugLevel:
		return hook.Writer.Debug(line)
	default:
		return nil
	}
}

func (hook *SrslogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *SrslogHook) SetFormatter(f srslog.Formatter) {
	hook.Writer.SetFormatter(f)
}

package formatter

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat       = "%lvl% [ %time% ] %file% %func% => %msg% %fields%"
	defaultTimestampFormat = "2006-01-02 15.04.05"
)

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	// Timestamp format
	TimestampFormat string
	// Available standard keys: time, msg, lvl
	// Also can include custom fields but limited to strings.
	// All of fields need to be wrapped inside %% i.e %time% %msg%
	LogFormat string
}

// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	colorFormat := NewFormatter()

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	// method 1
	//strList := strings.Split(entry.Caller.File, "/")
	//fileName := fmt.Sprintf("%s/%s:%d", strList[len(strList)-2], strList[len(strList)-1], entry.Caller.Line)

	// method 2
	//fileLog := fmt.Sprintf("%s:%d", filepath.Base(entry.Caller.File), entry.Caller.Line)

	// method 3
	pwd, _ := os.Getwd()
	fileName := fmt.Sprintf("%s:%d", strings.ReplaceAll(entry.Caller.File, pwd+"/", ""), entry.Caller.Line)

	output = strings.Replace(output, "%file%", fileName, 1)

	output = strings.Replace(output, "%func%", fmt.Sprintf("%s()", filepath.Base(entry.Caller.Function)), 1)

	level := fmt.Sprintf("%-5s", strings.ToUpper(entry.Level.String()))
	output = strings.Replace(output, "%lvl%", colorFormat.Color(entry, level), 1)

	var fields string

	// Put keys in a string array and sort it.
	keys := make([]string, len(entry.Data))
	i := 0
	for k := range entry.Data {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	// Do the rest.
	for _, key := range keys {
		fields = fmt.Sprintf("%s %s=%v", fields, colorFormat.Color(entry, key), entry.Data[key])
	}

	output = strings.Replace(output, "%fields%", fields, 1)
	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
}

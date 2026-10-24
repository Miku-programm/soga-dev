package yaml

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	T "recovered.local/soga/net/udpnat"
)

// Original initial data0x3787AB0 is info=1, separate from logrus' DebugLevel.
var levelAt3787AB0 H2hbGW = 1
var eventsAt5096988 = make(chan any)
var observableAt5096990 = T.AkUZ2QiKCu3n(eventsAt5096988)

// Original BAvZOr3jHqy logger is logrus. Dependency layout/API is verified;
// the original exact patch version is not established.
var loggerAt5029DF8 = logrus.StandardLogger()

func init() { loggerAt5029DF8.SetOutput(os.Stdout); loggerAt5029DF8.SetLevel(logrus.DebugLevel) } //0x1b33380

// Dl2Gj6:0x1b33480. Publish the warning before applying the console threshold.
// The unbuffered send is part of the original backpressure behavior.
func Dl2Gj6(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	eventsAt5096988 <- GPDvVWMg{ZbnaAXx: 2, URokq7: message}
	uRTsMwuzLPTw(2, message)
}

// uRTsMwuzLPTw:0x1b33560 maps the four emitted levels to logrus Logln.
func uRTsMwuzLPTw(level H2hbGW, message string) {
	if level < levelAt3787AB0 {
		return
	}
	switch level {
	case 0:
		loggerAt5029DF8.Logln(logrus.DebugLevel, message)
	case 1:
		loggerAt5029DF8.Logln(logrus.InfoLevel, message)
	case 2:
		loggerAt5029DF8.Logln(logrus.WarnLevel, message)
	case 3:
		loggerAt5029DF8.Logln(logrus.ErrorLevel, message)
	}
}

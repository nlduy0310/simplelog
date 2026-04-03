package simplelog

import "fmt"

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type levelInfo struct {
	name     string
	severity int
}

var infoByLevel = map[Level]levelInfo{
	DEBUG:   {name: "DEBUG", severity: 0},
	INFO:    {name: "INFO", severity: 1},
	WARNING: {name: "WARNING", severity: 2},
	ERROR:   {name: "ERROR", severity: 3},
	FATAL:   {name: "FATAL", severity: 4},
}

func (l Level) Name() string {
	return mustGetLevelInfo(l).name
}

func (l Level) String() string {
	return l.Name()
}

func (l Level) Allow(other Level) bool {
	thisSeverity := mustGetLevelInfo(l).severity
	otherSeverity := mustGetLevelInfo(other).severity
	return thisSeverity <= otherSeverity
}

func assertValidLevel(l Level) {
	mustGetLevelInfo(l)
}

func mustGetLevelInfo(l Level) levelInfo {
	if i, err := getLevelInfo(l); err == nil {
		return i
	} else {
		panic(invalidLevel(l))
	}
}

func getLevelInfo(l Level) (levelInfo, error) {
	if i, ok := infoByLevel[l]; ok {
		return i, nil
	} else {
		return levelInfo{}, invalidLevel(l)
	}
}

func invalidLevel(l Level) error {
	return fmt.Errorf("invalid log level %v", l)
}

package log5ws

type Logger interface {
	Info(*LogInfo)
}

// public object
type LogInfo struct {
	Logger Logger
	ID     string
	Who    string
	What   string
	When   string
	Where  string
	Why    string
	Msg    string
}

// private object
type Log struct {
	logger Logger
	id     string
	who    string
	what   string
	when   string
	where  string
	why    string
	msg    string
}

func New(logger Logger) *Log {
	return &Log{
		logger: logger,
	}
}

func (l *Log) clone() *Log {
	return &Log{
		logger: l.logger,
		who:    l.who,
		what:   l.what,
		when:   l.when,
		where:  l.where,
		why:    l.why,
		id:     l.id,
	}
}

func (l *Log) ID(v string) *Log {
	l.id = v
	return l.clone()
}

func (l *Log) Who(v string) *Log {
	l.who = v
	return l.clone()
}

func (l *Log) What(v string) *Log {
	l.what = v
	return l.clone()
}

func (l *Log) When(v string) *Log {
	l.when = v
	return l.clone()
}

func (l *Log) Where(v string) *Log {
	l.where = v
	return l.clone()
}

func (l *Log) Why(v string) *Log {
	l.why = v
	return l.clone()
}

func (l *Log) Log(v string) *Log {
	l.why = v
	return l.clone()
}

func (l *Log) Info() {
	l.logger.Info(
		&LogInfo{
			Logger: l.logger,
			ID:     l.id,
			Who:    l.who,
			What:   l.what,
			When:   l.when,
			Where:  l.where,
			Why:    l.why,
			Msg:    l.msg,
		},
	)
}

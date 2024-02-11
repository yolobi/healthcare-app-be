package enums

const (
	DDMMYY = "2006-01-02"
)

type TimeDuration int

const (
	YEAR        TimeDuration = 0
	MONTH                    = 1
	DAY                      = 2
	HOUR                     = 3
	MINUTE                   = 4
	SECOND                   = 5
	MILLISECOND              = 6
)

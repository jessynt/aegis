package event_dao

const (
	quoteString       = "`"
	placeholderString = "?"
)

func quoted(s string) string { return quoteString + s + quoteString }

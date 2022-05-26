package validation

var reservedKeys = []string{
	"Country",
	"Province",
	"City",
	"Date",
	"Year",
	"Quarter",
	"Month",
	"DayOfMonth",
	"DayOfWeek",
	"HourOfDay",
}

func IsReservedName(key string) bool {
	for _, a := range reservedKeys {
		if a == key {
			return true
		}
	}
	return false
}

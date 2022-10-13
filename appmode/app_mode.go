package appmode

//go:generate go run github.com/dmarkham/enumer -type=AppMode -json
type AppMode int

const (
	Development AppMode = iota
	Production
)

func GetAppMode(appMode string) AppMode {
	mode, err := AppModeString(appMode)
	if err != nil {
		return Development
	}
	return mode
}

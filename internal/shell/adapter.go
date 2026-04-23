package shell

type Adapter interface {
	GenerateAlias(name, command string) string
	RCFilePath() string
	ManagedBlockMarkers() (start, end string)
}

package config

//LoadedData allows you to get access data that was in json config file
type LoadedData interface {
	FileName() string
	Get(string) string
	GetBool(string) bool
	GetFloat(string) float64
	GetInt(string) int64
	GetUint(string) uint64
}

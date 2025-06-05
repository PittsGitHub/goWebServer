package types

const ConfigFileName = "config.json"

type ConfigFile struct {
	Port string `json:"port"`
}

package example

import (
	"fmt"

	"github.com/1bro23/godotenvstruct.git"
)

type Config struct {
	HostCustomName string `env:"ENV_PREFIX_Config_Host"`
	Port           string //without custom name, default: {ENV_PREFIX}_{StructName}_{FieldName}
}

func main() {
	var cfg Config
	godotenvstruct.Bind("ENV_PREFIX_", &cfg)
	fmt.Println(cfg.HostCustomName, cfg.Port)
}

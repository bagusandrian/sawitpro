package config

type (
	Config struct {
		Database Database `yaml:"database"`
		Server   Server   `yaml:"server"`
	}
	Server struct {
		HTTP         HTTP   `yaml:"http"`
		JWTSecretKey string `yaml:"jwt_secret_key"`
	}
	HTTP struct {
		Address string `yaml:"address"`
		Timeout int    `yaml:"timeout"`
	}
	Database struct {
		Master DBConfig `yaml:"master"`
		Slave  DBConfig `yaml:"slave"`
	}
	DBConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DDname   string `yaml:"dbname"`
	}
	DBCredential struct {
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
		Title    string `json:"title"  yaml:"title"`
	}
)

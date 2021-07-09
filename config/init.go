package config

var ENV = &Env{}

func InitConfig() {
	ENV.Initenv()
}

package config

var (
	configContainer = make(map[string]string) //配置信息容器
)

type Config interface {
	Get(node string) Node
}

type Node interface {
	Int() (int, error)       //
	String() string          //字符串
	Float() (float64, error) //
}

func NewConfig() Config {
	return new(implementConfig)
}

func LoadConfig() error {
	return nil
}

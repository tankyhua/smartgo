package config

import (
	"strconv"
)

type implementNode struct {
	node string
}

func (this *implementNode) String() string {
	return configContainer[this.node]
}

func (this *implementNode) Int() (int, error) {
	return strconv.Atoi(configContainer[this.node])
}

func (this *implementNode) Float() (float64, error) {
	return strconv.ParseFloat(configContainer[this.node], 64)
}

type implementConfig struct {
}

func (this *implementConfig) Get(node string) Node {
	var nd = new(implementNode)
	nd.node = node
	return nd
}

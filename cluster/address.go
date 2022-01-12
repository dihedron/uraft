package cluster

import (
	"fmt"
	"strconv"
	"strings"
)

type Address struct {
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	Port int    `json:"port,omitempty" yaml:"port,omitempty"`
}

func (a Address) String() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func (a *Address) UnmarshalFlag(value string) error {
	tokens := strings.Split(value, ":")
	if len(tokens) != 2 {
		return fmt.Errorf("invalid format for address, expected '[<host>]:<port>', got '%s'", value)
	}
	host := strings.TrimSpace(tokens[0])
	port, err := strconv.Atoi(strings.TrimSpace(tokens[1]))
	if err != nil {
		return fmt.Errorf("invalid format for port: %w", err)
	}
	a.Host = host
	a.Port = port
	return nil
}

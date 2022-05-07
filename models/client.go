package models

type Client struct {
	Name  string   `yaml:"name"`
	Roles []string `yaml:"roles"`
}

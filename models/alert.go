package models

type RecvAlert struct {
	Kind       string `yaml:"kind"`
	Difference string `yaml:"difference"`
	Path       string `yaml:"path"`
	IssuedAt   int64  `yaml:"issued_at"`
}

type PrintableAlert struct {
	Kind       string `yaml:"kind"`
	Difference string `yaml:"difference"`
	Path       string `yaml:"path"`
	IssuedAt   string `yaml:"issued_at"`
}

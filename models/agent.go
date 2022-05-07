package models

type RecvAgent struct {
	Name              string   `yaml:"name"`
	HasBaseline       bool     `yaml:"has_baseline"`
	BaselineIsCurrent bool     `yaml:"baseline_is_current"`
	WatchedPaths      []string `yaml:"watched_paths"`
}

type Agent struct {
	Name         string   `yaml:"name"`
	WatchedPaths []string `yaml:"watched_paths"`
}

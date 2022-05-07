package models

type WatchedPathsUpdate struct {
	Name         string   `yaml:"name"`
	WatchedPaths []string `yaml:"watched_paths"`
}

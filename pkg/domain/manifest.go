package domain

type ColonyManifest struct {
	APIVersion   string             `yaml:"apiVersion"`
	Name         string             `yaml:"name"`
	Version      string             `yaml:"version"`
	Description  string             `yaml:"description"`
	Maintainers  []Maintainer       `yaml:"maintainers"`
	Dependencies []Dependency       `yaml:"dependencies"`
	Conditions   *Conditions        `yaml:"conditions,omitempty"`
}

type Maintainer struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email,omitempty"`
	URL   string `yaml:"url,omitempty"`
}

type Dependency struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Conditions struct {
	ColonyOSVersion string `yaml:"colonyOSVersion,omitempty"`
	Architecture    string `yaml:"architecture,omitempty"`
}

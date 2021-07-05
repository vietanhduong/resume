package home

type Duration struct {
	Start string `yaml:"start,omitempty"`
	End   string `yaml:"end,omitempty"`
}

type Achievement struct {
	Name     string    `yaml:"name"`
	Duration Duration `yaml:"duration"`
	Details  []string  `yaml:"details,flow,omitempty"`
}

type Section struct {
	Name         string         `yaml:"name"`
	Achievements []Achievement `yaml:"achievements,flow"`
}

type Metadata struct {
	Address  string   `yaml:"address"`
	Phone    string   `yaml:"phone"`
	Email    string   `yaml:"email"`
	Websites []string `yaml:"websites,flow,omitempty"`
}

type Resume struct {
	Name     string     `yaml:"name"`
	Metadata Metadata  `yaml:"metadata"`
	Spec     []Section `yaml:"spec"`
}

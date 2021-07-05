package home

// Duration block
type Duration struct {
	// required
	Start string `yaml:"start,omitempty"`
	// +optional
	End string `yaml:"end,omitempty"`
}

// Achievement block
type Achievement struct {
	// required
	Name string `yaml:"name"`
	// required
	Duration Duration `yaml:"duration"`
	// the achievement details
	// +optional
	Details []string `yaml:"details,flow,omitempty"`
}

// Section block
type Section struct {
	// required
	Name string `yaml:"name"`
	// required: at least 1 achievement
	Achievements []Achievement `yaml:"achievements,flow"`
}

// Metadata block
type Metadata struct {
	// required
	Address string `yaml:"address"`
	// required
	Phone string `yaml:"phone"`
	// required
	Email string `yaml:"email"`
	// +optional
	Websites []string `yaml:"websites,flow,omitempty"`
}

// Resume represent resume
type Resume struct {
	// required
	Name string `yaml:"name"`
	// required
	Metadata Metadata `yaml:"metadata"`
	// required: at least 1 section
	Spec []Section `yaml:"spec"`
}

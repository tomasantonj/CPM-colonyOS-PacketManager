package domain

// PackageService defines operations for managing package files on disk
type PackageService interface {
	// Initialize creates the package scaffolding
	Initialize(name string) error

	// LoadManifest reads the colony.yaml from a path
	LoadManifest(path string) (*ColonyManifest, error)

	// Pack creates a compressed artifact from the package directory
	Pack(path string, name string, version string) (string, error)

	// Unpack extracts a compressed artifact to a destination directory
	Unpack(artifactPath string, destPath string) error
}

// RegistryService defines operations for interacting with the remote registry
type RegistryService interface {
	Publish(artifactPath string) error
	Fetch(packageName, version string) (string, error)
	Search(query string) ([]string, error)
}

// TemplateEngine defines operations for rendering ColonyOS specs
type TemplateEngine interface {
	// Render takes any templates in the package and renders them using the values.yaml
	Render(packagePath string, values map[string]interface{}) ([]byte, error)
}

// Submitter defines the interface for submitting to ColonyOS
type Submitter interface {
	SubmitWorkflow(specJSON []byte) error
	RegisterFunction(specJSON []byte) error
}

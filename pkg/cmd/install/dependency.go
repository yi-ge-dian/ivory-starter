package install

// DependencyInstaller is an interface that defines the method for installing dependencies.
type DependencyInstaller interface {
	InstallDependencies() error
}

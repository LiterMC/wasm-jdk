package ir

type ClassLoader interface {
	LoadClass(name string) (Class, error)
	AvaliablePackages() ([]string)
	PackageLocation(name string) string
}

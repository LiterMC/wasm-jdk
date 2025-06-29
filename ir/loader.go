package ir

type ClassLoader interface {
	DefineClass(class Class)
	LoadClass(name string) (Class, error)
	LoadedClass(name string) Class
	AvaliablePackages() []string
	PackageLocation(name string) string
}

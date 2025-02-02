package vm

// VM Options
type Options struct {
	Loader      ClassLoader
	EntryClass  string
	EntryMethod string
	EntryArgs   []string
}

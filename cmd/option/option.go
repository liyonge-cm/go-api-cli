package option

type CommandOption interface {
	OPTION() string
	Help() string
	SetOptions(string) error
	Get() string
}

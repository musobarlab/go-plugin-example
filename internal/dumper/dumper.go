package dumper

type Dumper interface {
	Dump(string) error
}
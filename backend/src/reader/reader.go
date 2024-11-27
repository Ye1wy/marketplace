package reader

type Reader interface {
	Parse() (*Data, error)
}

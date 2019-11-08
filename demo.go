package demo_doc

import "io"

// Exported type has a function called `ExportedFunction` but it's not in docs
type Exported struct {
	io.Closer
	notExported
}

type notExported interface {
	ExportedFunction() error
}

// Example is in docs
func (Exported) Example() {

}

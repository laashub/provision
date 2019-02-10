package models

// Bundle holds the object's store's name
//
// swagger: model
type Bundled struct {
	// Bundle tracks the name of the store containing this object
	// read only: true
	Bundle string
}

// SetBundle sets the name of the content layer holding this object.
func (b *Bundled) SetBundle(name string) {
	b.Bundle = name
}

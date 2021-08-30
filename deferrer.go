package deferrer

// Deferrer supports deferring (resource cleanup) operations in an outer
// function. It basically defers defers. Usage:
//
//     outer := Deferrer{}
//     defer outer.Cleanup()
//
//     var f *os.File
//     func(){
//         f, _ := os.Open("foobar")
//         outer.Defer(func(){f.Close()})
//     }()
//
//     func(){
//         var b byte[256]
//	       _, _ = f.Read(b)
//     }()
type Deferrer []Deferred

// Deferred registers a deferred defer function: the registered defer function
// will only be called on triggering Deferrer.Cleanup, which should be done in
// the (outer) function where the Deferrer was created.
type Deferred func()

// Defer a function in an outer function.
func (d *Deferrer) Defer(f func()) {
	*d = append(*d, f)
}

// Cleanup calls all deferred clean-ups and must be deferred itself in the same
// function where the Deferrer has been created.
func (d *Deferrer) Cleanup() {
	for _, deferred := range *d {
		deferred()
	}
}

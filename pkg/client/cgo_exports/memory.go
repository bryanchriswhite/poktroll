package main

// #include <client.h>
import "C"
import "fmt"

const (
	NilRef  = GoRef(-1)
	CNilRef = C.go_ref(NilRef)
)

var (
	goMemoryMap  = map[GoRef]any{}
	nextGoMemRef = GoRef(0)
)

type GoRef int64

// main is a dummy function to satisfy the cgo requirements.
func main() {}

// TODO_IN_THIS_COMMIT: godoc...
func SetGoMem(value any) C.go_ref {
	nextGoMemRef++
	goMemoryMap[nextGoMemRef] = value
	return C.go_ref(nextGoMemRef)
}

// TODO_IN_THIS_COMMIT: godoc...
func GetGoMem[T any](ref C.go_ref) (T, error) {
	value, ok := goMemoryMap[GoRef(ref)]
	if !ok {
		return *new(T), fmt.Errorf("go memory reference not found: %d", ref)
	}

	valueT, ok := value.(T)
	if !ok {
		return valueT, fmt.Errorf("expected %T, got: %T", *new(T), value)
	}

	return valueT, nil
}

// TODO_IN_THIS_COMMIT: godoc...
//
//export FreeGoMem
func FreeGoMem(ref C.go_ref) {
	delete(goMemoryMap, GoRef(ref))
}

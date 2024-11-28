package main

// #include <client.h>
import "C"
import (
	"unsafe"

	"cosmossdk.io/depinject"
)

//export Supply
func Supply(goRef C.GoRef, cErr **C.char) C.GoRef {
	toSupply, err := GetGoMem[any](goRef)
	if err != nil {
		*cErr = C.CString(err.Error())
		return 0
	}

	return SetGoMem(depinject.Supply(toSupply))
}

//export SupplyMany
func SupplyMany(goRefs *C.GoRef, numGoRefs C.int, cErr **C.char) C.GoRef {
	refs := unsafe.Slice(goRefs, numGoRefs)

	//fmt.Printf(">>> refs: %+v\n", refs)
	//val, err := GetGoMem[any](refs[0])
	//if err != nil {
	//	*cErr = C.CString(err.Error())
	//	return 0
	//}
	//fmt.Printf(">>> ref: %+v\n", val)

	var toSupply []any
	for _, ref := range refs {
		valueToSupply, err := GetGoMem[any](ref)
		if err != nil {
			*cErr = C.CString(err.Error())
			//*cErr = C.CString(fmt.Sprintf("%+v", err))
			return 0
		}

		toSupply = append(toSupply, valueToSupply)
	}

	return SetGoMem(depinject.Supply(toSupply...))
}

//export Config
func Config(goRefs *C.GoRef, numGoRefs C.int, cErr **C.char) C.GoRef {
	refs := unsafe.Slice(goRefs, numGoRefs)

	var configs []depinject.Config
	for _, ref := range refs {
		cfg, err := GetGoMem[depinject.Config](ref)
		if err != nil {
			*cErr = C.CString(err.Error())
			return 0
		}

		configs = append(configs, cfg)
	}

	return SetGoMem(depinject.Configs(configs...))
}

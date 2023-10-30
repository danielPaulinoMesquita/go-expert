package main

func Compara[T Number](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

// this won't work because any cannot be comparable or defined
//func ComparaWithAnyParam[T any](a T, b T) bool {
//	if a == b {
//		return true
//	}
//
//	return false
//}

func ComparaWithComparableParam[T comparable](a T, b T) bool {
	//comparable verifies if is it the same
	if a == b {
		return true
	}

	return false
}

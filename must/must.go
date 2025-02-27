package must

func Succeed[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func BeTrue[T any](v T, cond bool) T {
	if !cond {
		panic("unexpected condition hit")
	}
	return v
}

package api

type Args struct {
	A int
	B int
	AAA *AAA
}

type AAA struct {
	X string
}
type Reply struct {
	C int
}

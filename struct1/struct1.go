package struct1

type Stores struct {
	StoresID  string `json:"StoresId"`
	StoreName string `json:"StoreName"`
	Catogiry  string `json:"catogiry"`
}

func Add(a int, b int) (c int) {
	c = a + b
	return
}

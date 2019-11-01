package mock

type Retiever struct {
	Contents string
}

func (r Retiever) Get(url string) string {
	return r.Contents
}


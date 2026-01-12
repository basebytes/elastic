package query

type Page struct {
	Start int `json:"from,omitempty"`
	Max   int `json:"size,omitempty"`
}

func (p *Page) From() int {
	if p.Start > maxCount {
		return -1
	}
	return p.Start
}

func (p *Page) Size() (max int) {
	if p.Max > maxSize {
		max = maxSize
	} else if p.Max <= 0 {
		max = defaultSize
	}
	if p.Start+max > maxCount {
		max = maxCount - p.Start
	}
	if max < 0 {
		max = 0
	}
	return
}

const (
	maxSize     = 100
	defaultSize = 10
	maxCount    = 10000
)

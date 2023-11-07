package abstractFactory

type IShort interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

type short struct {
	logo string
	size int
}

func (s *short) setLogo(logo string) {
	s.logo = logo
}

func (s *short) setSize(size int) {
	s.size = size
}

func (c *short) GetLogo() string {
	return c.logo
}

func (c *short) GetSize() int {
	return c.size
}

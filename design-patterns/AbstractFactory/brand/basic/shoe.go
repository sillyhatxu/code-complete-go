package basic

type Shoe struct {
	Logo string
	Size int
}

func (s *Shoe) GetLogo() string {
	return s.Logo
}

func (s *Shoe) SetLogo(logo string) {
	s.Logo = logo
}

func (s *Shoe) GetSize() int {
	return s.Size
}

func (s *Shoe) SetSize(size int) {
	s.Size = size
}

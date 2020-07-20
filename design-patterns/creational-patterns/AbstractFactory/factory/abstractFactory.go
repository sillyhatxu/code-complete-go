package factory

type ShortInterface interface {
	GetLogo() string
	SetLogo(logo string)
	GetSize() int
	SetSize(size int)
}

type ShoeInterface interface {
	GetLogo() string
	SetLogo(logo string)
	GetSize() int
	SetSize(size int)
}

type BrandFactory interface {
	MakeShoe() ShoeInterface
	MakeShort() ShortInterface
}

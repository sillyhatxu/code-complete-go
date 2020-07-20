package model

import "fmt"

type Banner struct {
	bannerString string
}

func (b *Banner) Banner(s string) {
	b.bannerString = s
	fmt.Printf("Banner.Banner (%v)\n", s)
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("Banner.ShowWithParen (%v)\n", b.bannerString)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("Banner.ShowWithAster (%v)\n", b.bannerString)
}

type PrintBanner struct {
	Banner
}

func (b *PrintBanner) PrintBanner(s string) {
	b.Banner.Banner(s)
}

func (b *PrintBanner) PrintWeek() {
	b.Banner.ShowWithParen()
}

func (b *PrintBanner) PrintStrong() {
	b.Banner.ShowWithAster()
}

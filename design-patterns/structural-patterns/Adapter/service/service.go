package service

type Bird interface {
	fly()
	makeSound()
}

type ToyDuck interface {
	squeak()
}

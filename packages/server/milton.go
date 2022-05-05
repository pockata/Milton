package milton



type FlowerPotService interface {
	Add(string, Unit) (FlowerPot, error)
	Remove(string) error
	All() ([]FlowerPot, error)
}

type FlowerPot interface {
	ID() string
	Name() string
	Update(name string) error
}

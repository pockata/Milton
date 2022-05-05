package milton

type UnitService interface {
	Pair(string, string) error
	Unpair(string) error
	All() ([]Unit, error)
}

type Unit interface {
	ID() string
	Name() string
	MDNS() string
}

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

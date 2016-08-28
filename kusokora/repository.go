package kusokora

type KusokoraRepository interface {
	GetAll() ([]Kusokora, error)
	Put(k Kusokora) error
}

type Kusokora struct {
	ID         int
	PictureURL string
	SourceURL  string
}

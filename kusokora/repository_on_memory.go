package kusokora

type KusokoraRepositoryOnMemory struct {
	Data []Kusokora
}

func NewKusokoraRepositoryOnMemory() *KusokoraRepositoryOnMemory {
	return &KusokoraRepositoryOnMemory{
		Data: []Kusokora{},
	}
}

func (kr *KusokoraRepositoryOnMemory) GetAll() ([]Kusokora, error) {
	return kr.Data, nil
}

func (kr *KusokoraRepositoryOnMemory) Put(k Kusokora) error {
	k.ID = kr.newID()
	kr.Data = append(kr.Data, k)
	return nil
}

func (kr *KusokoraRepositoryOnMemory) newID() int {
	if len(kr.Data) < 1 {
		return 1
	}
	ids := []int{}
	for _, k := range kr.Data {
		ids = append(ids, k.ID)
	}
	return max(ids) + 1
}

func max(a []int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

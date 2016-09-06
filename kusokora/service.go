package kusokora

type KusokoraService struct {
	kr KusokoraRepository
}

func (ks *KusokoraService) AddKusokora(picURL string) error {
	return ks.kr.Put(Kusokora{
		PictureURL: picURL,
	})
}

func NewKusokoraService(r KusokoraRepository) KusokoraService {
	return KusokoraService{
		kr: r,
	}
}

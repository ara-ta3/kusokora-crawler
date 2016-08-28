package kusokora

type KusokoraService struct {
	kr KusokoraRepository
}

func (ks *KusokoraService) AddKusokora(picURL, sourceURL string) error {
	return ks.kr.Put(Kusokora{
		PictureURL: picURL,
		SourceURL:  sourceURL,
	})
}

func NewKusokoraService(r KusokoraRepository) KusokoraService {
	return KusokoraService{
		kr: r,
	}
}

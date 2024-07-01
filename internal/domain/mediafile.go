package domain

type MediaFile struct {
	ID    int    `"db:id"`
	Name  string `"db:name"`
	Image []byte `"db:image"`
}

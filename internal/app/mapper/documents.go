package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllDocuments(f []domain.Document, bucketName string) (*responses.GetAllDocuments, error) {
	documents := make([]responses.Document, 0, len(f))
	for _, v := range f {
		documents = append(documents,
			responses.Document{
				ID:     v.ID,
				Name:   v.Name,
				Key:    bucketName + "/" + v.Key,
				ClubID: v.ClubID,
			})
	}

	return &responses.GetAllDocuments{Documents: documents}, nil
}

func MakeResponseDocument(v *domain.Document, bucketName string) (*responses.GetDocument, error) {
	return &responses.GetDocument{
		ID:     v.ID,
		Name:   v.Name,
		Key:    bucketName + "/" + v.Key,
		ClubID: v.ClubID,
	}, nil
}

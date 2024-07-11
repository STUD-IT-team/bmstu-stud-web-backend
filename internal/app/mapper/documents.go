package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllDocuments(d []domain.Document, bucketName string) (*responses.GetAllDocuments, error) {
	documents := make([]responses.Document, 0, len(d))
	for _, v := range d {
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

func MakeResponseDocumentsByClubID(d []domain.Document, bucketName string) (*responses.GetDocumentsByClubID, error) {
	documents := make([]responses.Document, 0, len(d))
	for _, v := range d {
		documents = append(documents,
			responses.Document{
				ID:     v.ID,
				Name:   v.Name,
				Key:    bucketName + "/" + v.Key,
				ClubID: v.ClubID,
			})
	}

	return &responses.GetDocumentsByClubID{Documents: documents}, nil
}

func MakeRequestPostDocument(v requests.PostDocument) *domain.Document {
	return &domain.Document{
		Name:   v.Name,
		Key:    v.Key,
		ClubID: v.ClubID,
	}
}

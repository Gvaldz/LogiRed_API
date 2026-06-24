package domain

import "logired/src/internal/documents/domain/entities"

type IDocument interface {
    CreateDocumentTx(tx interface{}, document entities.Document) error
}
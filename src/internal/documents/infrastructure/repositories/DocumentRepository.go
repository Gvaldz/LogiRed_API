package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"logired/src/internal/documents/domain"
	"logired/src/internal/documents/domain/entities"
)

type DocumentRepo struct {
	db *sql.DB
}

func NewDocumentRepo(db *sql.DB) domain.IDocument {
	return &DocumentRepo{
		db: db,
	}
}

func (repo *DocumentRepo) CreateDocumentTx(tx interface{}, document entities.Document) error {
	sqlTx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("error crítico: la transacción provista no es del tipo *sql.Tx")
	}


	query := `INSERT INTO documents (id_driver, id_document_type, url) VALUES ($1, $2, $3)`

	_, err := sqlTx.Exec(query, document.IdDriver, document.IdDocumentType, document.URL)
	if err != nil {
		return fmt.Errorf("error al guardar el documento tipo %d en base de datos: %w", document.IdDocumentType, err)
	}

	return nil
}
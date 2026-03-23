package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"logired/src/internal/proposals/domain/entities"
)

type ProposalRepo struct {
	db *sql.DB
}

func NewProposalRepo(db *sql.DB) *ProposalRepo {
	return &ProposalRepo{db: db}
}

func (r *ProposalRepo) CreateProposal(proposal entities.Proposal) error {
	query := `INSERT INTO proposals (price, comment, iddriver, idride, idproposalstatus) 
	          VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, proposal.Price, proposal.Comment, proposal.IdDriver, proposal.IdRide, proposal.IdStatus)
	if err != nil {
		return fmt.Errorf("error al crear propuesta: %w", err)
	}
	log.Println("[ProposalRepo] Propuesta creada correctamente")
	return nil
}

func (r *ProposalRepo) AcceptProposal(idProposal int32) error {
	query := `UPDATE proposals SET idproposalstatus = 1 WHERE idproposal = ?`
	result, err := r.db.Exec(query, idProposal)
	if err != nil {
		return fmt.Errorf("error al aceptar propuesta: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("propuesta no encontrada")
	}
	log.Println("[ProposalRepo] Propuesta aceptada correctamente")
	return nil
}

func (r *ProposalRepo) DeleteProposal(idProposal int32, idDriver int32) error {
	query := `DELETE FROM proposals WHERE idproposal = ? AND iddriver = ?`
	result, err := r.db.Exec(query, idProposal, idDriver)
	if err != nil {
		return fmt.Errorf("error al eliminar propuesta: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("propuesta no encontrada o no tienes permiso para eliminarla")
	}
	log.Println("[ProposalRepo] Propuesta eliminada correctamente")
	return nil
}
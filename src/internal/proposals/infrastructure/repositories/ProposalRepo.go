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
	query := `INSERT INTO proposals (price, comment, iddriver, idride, idproposalstatus, idcar) 
	          VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, proposal.Price, proposal.Comment, proposal.IdDriver, proposal.IdRide, proposal.IdStatus, proposal.IdCar)
	if err != nil {
		return fmt.Errorf("error al crear propuesta: %w", err)
	}
	log.Println("[ProposalRepo] Propuesta creada correctamente")
	return nil
}

func (r *ProposalRepo) AcceptProposal(idProposal int32, idStatus int32) error {
    query := "UPDATE proposals SET idproposalstatus = ? WHERE idproposal = ?"
    result, err := r.db.Exec(query, idStatus, idProposal)
    if err != nil {
        return fmt.Errorf("error al actualizar propuesta: %w", err)
    }
    rows, _ := result.RowsAffected()
    if rows == 0 {
        return fmt.Errorf("propuesta no encontrada")
    }
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

func (r *ProposalRepo) GetProposalById(idProposal int32) (entities.Proposal, error) {
	var p entities.Proposal
	query := "SELECT idproposal, price, comment, iddriver, idride, idproposalstatus, idcar FROM proposals WHERE idproposal = ?"
	err := r.db.QueryRow(query, idProposal).Scan(
		&p.IdProposal, &p.Price, &p.Comment, &p.IdDriver, &p.IdRide, &p.IdStatus, &p.IdCar,
	)
	if err != nil {
		return entities.Proposal{}, fmt.Errorf("propuesta no encontrada: %w", err)
	}
	return p, nil
}

func (r *ProposalRepo) GetProposalsByRideId(idRide int32) ([]entities.Proposal, error) {
    query := `
        SELECT idproposal, price, comment, iddriver, idride, idproposalstatus, idcar 
        FROM proposals 
        WHERE idride = ?
    `
    rows, err := r.db.Query(query, idRide)
    if err != nil {
        return nil, fmt.Errorf("error al obtener propuestas: %w", err)
    }
    defer rows.Close()

    var proposals []entities.Proposal
    for rows.Next() {
        var p entities.Proposal
        if err := rows.Scan(
            &p.IdProposal,
            &p.Price,
            &p.Comment,
            &p.IdDriver,
            &p.IdRide,
            &p.IdStatus,
			&p.IdCar,
        ); err != nil {
            return nil, fmt.Errorf("error al escanear propuesta: %w", err)
        }
        proposals = append(proposals, p)
    }
    return proposals, nil
}
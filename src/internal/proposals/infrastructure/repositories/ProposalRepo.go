package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"logired/src/internal/proposals/domain"
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
	          VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, proposal.Price, proposal.Comment, proposal.IdDriver, proposal.IdRide, proposal.IdStatus)
	if err != nil {
		return fmt.Errorf("error al crear propuesta: %w", err)
	}
	log.Println("[ProposalRepo] Propuesta creada correctamente")
	return nil
}

func (r *ProposalRepo) AcceptProposal(idProposal int32, idStatus int32) error {
    query := "UPDATE proposals SET idproposalstatus = $1 WHERE idproposal = $2"
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
	query := `DELETE FROM proposals WHERE idproposal = $1 AND iddriver = $2`
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
	query := "SELECT idproposal, price, comment, iddriver, idride, idproposalstatus FROM proposals WHERE idproposal = $1"
	err := r.db.QueryRow(query, idProposal).Scan(
		&p.IdProposal, &p.Price, &p.Comment, &p.IdDriver, &p.IdRide, &p.IdStatus,
	)
	if err != nil {
		return entities.Proposal{}, fmt.Errorf("propuesta no encontrada: %w", err)
	}
	return p, nil
}

func (r *ProposalRepo) GetProposalDetailById(idProposal int32) (domain.ProposalDetail, error) {
	query := `
		SELECT
			p.idproposal, p.price, p.comment, p.iddriver, p.idride, p.idproposalstatus,
			d.rating,
			(SELECT COUNT(*) FROM rides r WHERE r.iddriver = p.iddriver) AS trip_count,
			c.idcar, c.car_registration, c.brand, c.model, c.color, c.max_capacity,
			c.frontview_image, c.backview_image, c.leftview_image, c.rightview_image, c.plates_image, c.space_image
		FROM proposals p
		INNER JOIN drivers d ON p.iddriver = d.iduser
		LEFT JOIN cars c ON c.iduser = p.iddriver
		WHERE p.idproposal = $1
	`
	var detail domain.ProposalDetail
	err := r.db.QueryRow(query, idProposal).Scan(
		&detail.IdProposal, &detail.Price, &detail.Comment, &detail.IdDriver, &detail.IdRide, &detail.IdStatus,
		&detail.Rating, &detail.TripCount,
		&detail.IdCar, &detail.CarRegistration, &detail.Brand, &detail.Model, &detail.Color, &detail.MaxCapacity,
		&detail.FrontViewImage, &detail.BackViewImage, &detail.LeftViewImage, &detail.RightViewImage,
		&detail.PlatesImage, &detail.SpacesImage,
	)
	if err != nil {
		return domain.ProposalDetail{}, fmt.Errorf("propuesta no encontrada: %w", err)
	}
	return detail, nil
}

func (r *ProposalRepo) GetProposalsByRideId(idRide int32) ([]entities.Proposal, error) {
    query := `
        SELECT idproposal, price, comment, iddriver, idride, idproposalstatus 
        FROM proposals 
        WHERE idride = $1
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
        ); err != nil {
            return nil, fmt.Errorf("error al escanear propuesta: %w", err)
        }
        proposals = append(proposals, p)
    }
    return proposals, nil
}
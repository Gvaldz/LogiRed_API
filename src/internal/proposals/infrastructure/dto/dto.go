package dto

type CreateProposalRequest struct {
    Price   float32 `json:"price"`
    IdRide  int32   `json:"id_ride"`
    Comment string  `json:"comment"`
}

type ProposalResponse struct {
    IdProposal int32   `json:"idproposal"`
    Price      float32 `json:"price"`
    Comment    string  `json:"comment"`
    IdDriver   int32   `json:"iddriver"`
    IdRide     int32   `json:"idride"`
    IdStatus   int32   `json:"idproposalstatus"`
}

type AcceptProposalRequest struct {
    IdStatus int32 `json:"idstatus" binding:"required"`
}
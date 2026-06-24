package entities

type Document struct {
    IdDocument     int32  `json:"id_document"`
    IdDriver       int32  `json:"id_driver"`       
    IdDocumentType int    `json:"id_document_type"` 
    URL            string `json:"url"`
}
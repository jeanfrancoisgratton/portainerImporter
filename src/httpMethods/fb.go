// portainerImporter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/httpMethods/fb.go
// Original timestamp: 2024/04/05 23:04

package httpMethods

type ParentsInfo struct {
	Nom    string `json:"Nom"`
	Prenom string `json:"prenom,omitempty""`
	RAMQ   string `json:"ramq,omitempty"`
	NAS    string `json:"nas,omitempty"`
}

type RAMQInfo struct {
	Nom          string        `json:"nom"`
	Prenom       string        `json:"prenom,omitempty"`
	AssMaladie   string        `json:"ramq"`
	Parents      []ParentsInfo `json:"parents,omitempty"`
	DossierActif bool          `json:"dossierActif,omitempty"`
}

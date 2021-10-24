package model

// Service -- Representing service request
type Service struct {
	ID          uint64 `json:"_id"`
	RequestID   uint64 `json:"requestId"`
	Status      string `json:"status"`
	VesselName  string `json:"vesselName"`
	ServiceType string `json:"serviceType"`
	DataAgent   string `json:"dataAgent"`
	Cargo       string `json:"cargo"`
	ETD         string `json:"etd"`
	ETA         string `json:"eta"`
	UserID      uint64 `json:"userID"`
}

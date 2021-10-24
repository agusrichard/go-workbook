package repository

import (
	"fmt"
	"golang-restapi/config"
	"golang-restapi/model"
)

// CreateServiceRequest -- create service request
func CreateServiceRequest(service *model.Service, userID uint64) {
	sqlQuery := `
		INSERT INTO services (
			request_id, 
			status, 
			vessel_name, 
			service_type,
			data_agent,
			cargo,
			etd,
			eta,
			user_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	_, err := config.DB.Exec(sqlQuery,
		service.RequestID,
		service.Status,
		service.VesselName,
		service.ServiceType,
		service.DataAgent,
		service.Cargo,
		service.ETD,
		service.ETA,
		userID,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Success to create service request %v\n", *service)
}

// GetServices -- Get all services by a user
func GetServices(userID uint64) []model.Service {
	sqlQuery := `
		SELECT * FROM services
		WHERE user_id = $1;
		
	`
	rows, err := config.DB.Query(sqlQuery, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var result []model.Service
	for rows.Next() {
		var service = model.Service{}
		err = rows.Scan(
			&service.ID,
			&service.RequestID,
			&service.Status,
			&service.VesselName,
			&service.ServiceType,
			&service.DataAgent,
			&service.Cargo,
			&service.ETD,
			&service.ETA,
			&service.UserID,
		)
		if err != nil {
			panic(err)
		}
		result = append(result, service)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("GetServices result", result)
	return result
}

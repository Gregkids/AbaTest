package repository

import (
	"database/sql"
	"errors"

	"aba.technical.test/models"
)

type DeviceSQL struct {
	DB *sql.DB
}

func (q *DeviceSQL) GetAllDevice() ([]models.DeviceData, error) {
	ret := []models.DeviceData{}

	// Query Get All Devices
	query := `
	SELECT
		d.device_name
	FROM public.device d `

	rows, err := q.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		data := models.DeviceData{}
		err = rows.Scan(&data.DeviceName)

		if err != nil {
			return nil, err
		}

		ret = append(ret, data)
	}

	return ret, nil
}

func (q *DeviceSQL) GetOneDevice(reqID string) ([]models.DeviceData, error) {
	ret := []models.DeviceData{}

	// Query Get Name by Id
	query := `
	SELECT
		d.device_id,
		d.device_name,
		COALESCE(d.location, '') AS location,
		d.status
	FROM public.device d `

	query = query + " WHERE d.device_id=$1"
	data := models.DeviceData{}
	err := q.DB.QueryRow(query, reqID).Scan(&data.DeviceId, &data.DeviceName, &data.DeviceLocation, &data.DeviceStatus)

	if err == sql.ErrNoRows {
		return nil, errors.New("data not found")
	} else if err != nil {
		return nil, err
	}

	ret = append(ret, data)

	return ret, nil
}

package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

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
	FROM public.device d 
	`

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

func (q *DeviceSQL) InsertDevice(req *models.DeviceReq) error {
	ctx := context.Background()
	tx, err := q.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Query Insert Name
	query := `
	INSERT INTO public.device
		(device_id, device_name, location, status, updated_at)
	VALUES
		($1, $2, $3, $4, $5); 
	`

	_, err = tx.ExecContext(ctx, query, req.DeviceId, req.DeviceName, req.DeviceLocation, req.DeviceStatus, time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (q *DeviceSQL) UpdateDevice(reqID string, req *models.DeviceReq) error {
	ctx := context.Background()
	tx, err := q.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Query Insert Name
	query := `
	UPDATE public.device SET
		device_name=$2, location=$3, status=$4, updated_at=$5
	WHERE device_id=$1; 
	`
	_, err = tx.ExecContext(ctx, query, req.DeviceId, req.DeviceName, req.DeviceLocation, req.DeviceStatus, time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (q *DeviceSQL) DeleteDevice(reqID string) error {
	ctx := context.Background()
	tx, err := q.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Query Insert Name
	query := `
	DELETE FROM public.device 
		WHERE device_id=$1; 
	`
	_, err = tx.ExecContext(ctx, query, reqID)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	} else {
		return nil
	}
}

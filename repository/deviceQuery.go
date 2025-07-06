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
	var status bool

	// Query Get All Devices
	query := `
	SELECT
		d.device_id,
		d.device_name,
		COALESCE(d.location, '') AS location,
		d.status,
		d.updated_at
	FROM public.device d 
	`

	rows, err := q.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		data := models.DeviceData{}
		err = rows.Scan(&data.DeviceId, &data.DeviceName, &data.DeviceLocation, &status, &data.UpdatedAt)

		if status {
			data.DeviceStatus = "Online"
		} else {
			data.DeviceStatus = "Offline"
		}

		if err != nil {
			return nil, err
		}

		ret = append(ret, data)
	}

	return ret, nil
}

func (q *DeviceSQL) GetOneDevice(reqID string) ([]models.DeviceData, error) {
	ret := []models.DeviceData{}
	var status bool

	// Query Get One Device
	query := `
	SELECT
		d.device_id,
		d.device_name,
		COALESCE(d.location, '') AS location,
		d.status,
		d.updated_at
	FROM public.device d 
	WHERE d.device_id=$1
	`

	data := models.DeviceData{}
	err := q.DB.QueryRow(query, reqID).Scan(&data.DeviceId, &data.DeviceName, &data.DeviceLocation, &data.DeviceStatus, &data.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("data not found")
	} else if err != nil {
		return nil, err
	}

	if status {
		data.DeviceStatus = "Online"
	} else {
		data.DeviceStatus = "Offline"
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

	// Query Insert Device
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

	// Query Update Device
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

	// Query Delete Device
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

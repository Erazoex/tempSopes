package ramcontroller

import (
	dbhelper "backend/helpers/db.helper"
	"backend/models"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

func Post(ramModel models.RAM) error {
	conn, err := dbhelper.NewConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := dbhelper.CreateTableIfNotExists(conn, models.RamTable, models.ProcessTable, models.ChildTable); err != nil {
		return err
	}

	datetime := time.Now().Format("2006-01-02 15:04:05.000")
	if err := insertRAM(conn, ramModel, datetime); err != nil {
		return err
	}

	return nil
}

func Get() (*models.RAM, error) {
	cmd := exec.Command("cat", "/proc/ram_so1_1s2024")
	// fmt.Println(cmd)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var ramInfo models.RAM
	if err := json.Unmarshal(output, &ramInfo); err != nil {
		return nil, err
	}
	return &ramInfo, nil
}

func GetAll() ([]models.RAM, error) {
	conn, err := dbhelper.NewConnection()
	if err != nil {
		return make([]models.RAM, 0), err
	}
	defer conn.Close()

	if err := dbhelper.CreateTableIfNotExists(conn, models.RamTable, models.ProcessTable, models.ChildTable); err != nil {
		return make([]models.RAM, 0), err
	}

	return getRAM(conn)

}

func insertRAM(db *dbhelper.Db, r models.RAM, datetime string) error {
	query := fmt.Sprintf(`INSERT INTO %s 
	(total, used, percentage, free, fecha) 
		VALUES (%v, %v, %v, %v, '%v');`,
		models.RamTableName, r.Total, r.Used, r.Percentage, r.Free, datetime)
	// fmt.Println(query)
	_, err := dbhelper.Execute(db, query)
	if err != nil {
		return err
	}

	return nil
}

func getRAM(db *dbhelper.Db) ([]models.RAM, error) {
	query := fmt.Sprintf(`SELECT * FROM %s ORDER BY id DESC LIMIT 50;`, models.RamTableName)
	// fmt.Println(query)
	rows, err := dbhelper.Query(db, query)
	if err != nil {
		return make([]models.RAM, 0), err
	}
	defer rows.Close()

	var rams []models.RAM
	for rows.Next() {
		var id int
		var total, used, percentage, free int64
		var dateContainer []uint8
		if err := rows.Scan(&id, &total, &used, &percentage, &free, &dateContainer); err != nil {
			return make([]models.RAM, 0), err
		}
		datetimeStr := string(dateContainer)
		datetime, err := time.Parse("2006-01-02 15:04:05.999999999", datetimeStr)
		if err != nil {
			return make([]models.RAM, 0), err
		}
		rams = append(rams, models.RAM{
			Total:      total,
			Used:       used,
			Percentage: percentage,
			Free:       free,
			Fecha:      datetime,
		})
	}
	if err := rows.Err(); err != nil {
		return make([]models.RAM, 0), err
	}
	return rams, nil
}

package processcontroller

import (
	dbhelper "backend/helpers/db.helper"
	"backend/models"
	"fmt"
)

func New() (*models.ProcessState, error) {
	conn, err := dbhelper.NewConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if err := dbhelper.CreateTableIfNotExists(conn, models.ProcessStateTable); err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`INSERT INTO %s
	(state)
		VALUES ("%v");
	`, models.ProcessStateTableName, "New")
	// fmt.Println(query)
	result, err := dbhelper.Execute(conn, query)
	if err != nil {
		return nil, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &models.ProcessState{
		Pid:   lastId,
		State: "New",
	}, nil
}

func Ready(pid int64) (*models.ProcessState, error) {
	conn, err := dbhelper.NewConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if err := dbhelper.CreateTableIfNotExists(conn, models.ProcessStateTable); err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`UPDATE %s
		SET state = "%s"
		WHERE pid = %v AND state NOT IN ('Terminated');`, models.ProcessStateTableName, "Ready", pid)
	_, err = dbhelper.Execute(conn, query)
	if err != nil {
		return nil, err
	}

	query = fmt.Sprintf(`SELECT pid, state FROM %s WHERE pid=%v`, models.ProcessStateTableName, pid)
	// fmt.Println(query)
	var value models.ProcessState
	err = conn.Db.QueryRow(query).Scan(&value.Pid, &value.State)
	return &value, err
}

func Running(pid int64) (*models.ProcessState, error) {
	conn, err := dbhelper.NewConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if err := dbhelper.CreateTableIfNotExists(conn, models.ProcessStateTable); err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`UPDATE %s
		SET state = "%s"
		WHERE pid = %v AND state NOT IN ('Waiting', 'Terminated')`, models.ProcessStateTableName, "Running", pid)
	_, err = dbhelper.Execute(conn, query)
	if err != nil {
		return nil, err
	}

	query = fmt.Sprintf(`SELECT pid, state FROM %s WHERE pid=%v`, models.ProcessStateTableName, pid)
	// fmt.Println(query)
	var value models.ProcessState
	err = conn.Db.QueryRow(query).Scan(&value.Pid, &value.State)
	return &value, err
}

func Waiting(pid int64) (*models.ProcessState, error) {
	conn, err := dbhelper.NewConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if err := dbhelper.CreateTableIfNotExists(conn, models.ProcessStateTable); err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`UPDATE %s
		SET state = "%s"
		WHERE pid = %v AND state NOT IN ('Ready', 'Terminated')`, models.ProcessStateTableName, "Waiting", pid)
	_, err = dbhelper.Execute(conn, query)
	if err != nil {
		return nil, err
	}

	query = fmt.Sprintf(`SELECT pid, state FROM %s WHERE pid=%v`, models.ProcessStateTableName, pid)
	// fmt.Println(query)
	var value models.ProcessState
	err = conn.Db.QueryRow(query).Scan(&value.Pid, &value.State)
	return &value, err
}

func Terminated(pid int64) (*models.ProcessState, error) {
	conn, err := dbhelper.NewConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if err := dbhelper.CreateTableIfNotExists(conn, models.ProcessStateTable); err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`UPDATE %s
		SET state = "%s"
		WHERE pid = %v AND state NOT IN ('Ready', 'Waiting')`, models.ProcessStateTableName, "Terminated", pid)
	// fmt.Println(query)
	_, err = dbhelper.Execute(conn, query)
	if err != nil {
		return nil, err
	}

	query = fmt.Sprintf(`SELECT pid, state FROM %s WHERE pid=%v`, models.ProcessStateTableName, pid)
	// fmt.Println(query)
	var value models.ProcessState
	err = conn.Db.QueryRow(query).Scan(&value.Pid, &value.State)
	return &value, err
}

func GetAll() ([]models.ProcessState, error) {
	conn, err := dbhelper.NewConnection()
	if err != nil {
		return (make([]models.ProcessState, 0)), err
	}
	defer conn.Close()

	if err := dbhelper.CreateTableIfNotExists(conn, models.ProcessStateTable); err != nil {
		return (make([]models.ProcessState, 0)), err
	}
	return getProcesses(conn)
}

func getProcesses(db *dbhelper.Db) ([]models.ProcessState, error) {
	query := fmt.Sprintf(`SELECT * FROM  %s;`, models.ProcessStateTableName)
	// fmt.Println(query)
	rows, err := dbhelper.Query(db, query)
	if err != nil {
		return make([]models.ProcessState, 0), err
	}
	defer rows.Close()

	var estados []models.ProcessState
	for rows.Next() {
		var pid int64
		var state string
		if err := rows.Scan(&pid, &state); err != nil {
			return make([]models.ProcessState, 0), err
		}
		estados = append(estados, models.ProcessState{
			Pid:   pid,
			State: state,
		})
	}
	if err := rows.Err(); err != nil {
		return make([]models.ProcessState, 0), err
	}
	return estados, nil
}

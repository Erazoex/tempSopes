package cpucontroller

import (
	dbhelper "backend/helpers/db.helper"
	"backend/models"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

func Post(cpuModel models.CPU) error {
	conn, err := dbhelper.NewConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := dbhelper.CreateTableIfNotExists(conn, models.CpuTable, models.ProcessTable, models.ChildTable); err != nil {
		return err
	}

	datetime := time.Now().Format("2006-01-02 15:04:05.000")
	if err := insertCPU(conn, cpuModel, datetime); err != nil {
		return err
	}

	return nil
}

func Get() (*models.CPU, error) {
	cmd := exec.Command("cat", "/proc/cpu_so1_1s2024")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var cpuInfo models.CPU
	if err := json.Unmarshal(output, &cpuInfo); err != nil {
		return nil, err
	}
	return &cpuInfo, nil
}

func GetAll() ([]models.CPU, error) {
	conn, err := dbhelper.NewConnection()
	if err != nil {
		return make([]models.CPU, 0), err
	}
	defer conn.Close()

	if err := dbhelper.CreateTableIfNotExists(conn, models.CpuTable, models.ProcessTable, models.ChildTable); err != nil {
		return make([]models.CPU, 0), err
	}
	return getCPUS(conn)
}

func insertCPU(db *dbhelper.Db, c models.CPU, datetime string) error {
	query := fmt.Sprintf(`INSERT INTO %s 
	(cpu_total, cpu_percentage, running, sleeping, zombie, stopped, total, fecha) 
		VALUES (%v, %v, %v, %v, %v, %v, %v, '%v');`,
		models.CpuTableName, c.Cpu_total, c.Cpu_percentage, c.Running, c.Sleeping, c.Zombie, c.Stopped, c.Total, datetime)
	// fmt.Println(query)
	_, err := dbhelper.Execute(db, query)
	if err != nil {
		return err
	}
	var allChildren []models.Child
	processValues := ""
	for index, process := range c.Processes {
		if len(process.Children) > 0 {
			allChildren = append(allChildren, process.Children...)
		}
		if index == len(c.Processes)-1 {
			processValues += fmt.Sprintf("(%v, \"%v\",  %v, %v, '%v');\n", process.Pid, process.Name, process.User, process.State, datetime)
		} else {
			processValues += fmt.Sprintf("(%v, \"%v\",  %v, %v, '%v'),\n", process.Pid, process.Name, process.User, process.State, datetime)
		}
	}

	query = fmt.Sprintf(`INSERT INTO %s
	(pid, name, user, state, fecha)
		VALUES
			%s`, models.ProcessTableName, processValues)
	// fmt.Println(query)
	_, err = dbhelper.Execute(db, query)
	if err != nil {
		return err
	}

	childValues := ""
	for index, child := range allChildren {
		if index == len(allChildren)-1 {
			childValues += fmt.Sprintf("(%v, \"%v\",  %v, %v, '%v');\n", child.Pid, child.Name, child.State, child.PidParent, datetime)
		} else {
			childValues += fmt.Sprintf("(%v, \"%v\",  %v, %v, '%v'),\n", child.Pid, child.Name, child.State, child.PidParent, datetime)
		}
	}

	query = fmt.Sprintf(`INSERT INTO %s
	(pid, name,state, pidParent, fecha)
		VALUES
			%s`, models.ChildTableName, childValues)
	// fmt.Println(query)
	_, err = dbhelper.Execute(db, query)
	if err != nil {
		return err
	}
	return nil
}

func getCPUS(db *dbhelper.Db) ([]models.CPU, error) {
	query := fmt.Sprintf(`SELECT * FROM %s ORDER BY id DESC LIMIT 50;`, models.CpuTableName)
	// fmt.Println(query)
	rows, err := dbhelper.Query(db, query)
	if err != nil {
		return make([]models.CPU, 0), err
	}
	defer rows.Close()

	var cpus []models.CPU
	for rows.Next() {
		var id int
		var cpu_total, cpu_percentage, running, sleeping, zombie, stopped, total int64
		var dateContainer []uint8
		if err := rows.Scan(&id, &cpu_total, &cpu_percentage, &running, &sleeping, &zombie, &stopped, &total, &dateContainer); err != nil {
			return make([]models.CPU, 0), err
		}
		datetimeStr := string(dateContainer)
		datetime, err := time.Parse("2006-01-02 15:04:05.999999999", datetimeStr)
		if err != nil {
			return make([]models.CPU, 0), err
		}
		cpus = append(cpus, models.CPU{
			Cpu_total:      cpu_total,
			Cpu_percentage: cpu_percentage,
			Processes:      make([]models.Process, 0),
			Running:        running,
			Sleeping:       sleeping,
			Zombie:         zombie,
			Stopped:        stopped,
			Total:          total,
			Fecha:          datetime,
		})
	}
	if err := rows.Err(); err != nil {
		return make([]models.CPU, 0), err
	}
	return cpus, nil
}

package models

import (
	"fmt"
	"time"
)

type CPU struct {
	Cpu_total      int64     `json:"cpu_total"`
	Cpu_percentage int64     `json:"cpu_percentage"`
	Processes      []Process `json:"processes"`
	Running        int64     `json:"running"`
	Sleeping       int64     `json:"sleeping"`
	Zombie         int64     `json:"Zombie"`
	Stopped        int64     `json:"stopped"`
	Total          int64     `json:"total"`
	Fecha          time.Time `json:"fecha"`
}

func newCPUModel(cpu_total, cpu_percentage int64, processes []Process, running, sleeping, zombie, stopped, total int64) *CPU {
	return &CPU{
		Cpu_total:      cpu_total,
		Cpu_percentage: cpu_percentage,
		Processes:      processes,
		Running:        running,
		Sleeping:       sleeping,
		Zombie:         zombie,
		Stopped:        stopped,
		Total:          total,
	}
}

type Process struct {
	Pid      int64     `json:"pid"`
	Name     string    `json:"name"`
	User     int64     `json:"user"`
	State    int64     `json:"state"`
	Children []Child   `json:"child"`
	Fecha    time.Time `json:"fecha"`
}

func newProcessModel(pid int64, name string, user, state int64, children []Child) *Process {
	return &Process{
		Pid:      pid,
		Name:     name,
		User:     user,
		State:    state,
		Children: children,
	}
}

type Child struct {
	Pid       int64     `json:"pid"`
	Name      string    `json:"name"`
	State     int64     `json:"state"`
	PidParent int64     `json:"pidParent"`
	Fecha     time.Time `json:"fecha"`
}

func newChildModel(pid int64, name string, state, pidParent int64) *Child {
	return &Child{
		Pid:       pid,
		Name:      name,
		State:     state,
		PidParent: pidParent,
	}
}

var CpuTableName string = "cpu"
var CpuTable string = fmt.Sprintf(`
	%s (
		id INT AUTO_INCREMENT,
		cpu_total INT,
		cpu_percentage INT,
		running INT,
		sleeping INT,
		zombie INT,
		stopped INT,
		total INT,
		fecha DATETIME(6),
		PRIMARY KEY (id)
	);
`, CpuTableName)

var ProcessTableName string = "process"
var ProcessTable string = fmt.Sprintf(`
	%s (
		id INT AUTO_INCREMENT,
		pid INT,
		name VARCHAR(255),
		user INT,
		state INT,
		fecha DATETIME(6),
		PRIMARY KEY (id)
	);
`, ProcessTableName)

var ChildTableName string = "child"
var ChildTable string = fmt.Sprintf(`
	%s (
		id INT AUTO_INCREMENT,
		pid INT,
		name VARCHAR(255),
		state INT,
		pidParent INT,
		fecha DATETIME(6),
		PRIMARY KEY (id)
	);
`, ChildTableName)

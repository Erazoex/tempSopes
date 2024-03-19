package models

import (
	"fmt"
	"time"
)

type RAM struct {
	Total      int64     `json:"totalRam"`
	Used       int64     `json:"usedMemory"`
	Percentage int64     `json:"percentage"`
	Free       int64     `json:"freeMemory"`
	Fecha      time.Time `json:"fecha"`
}

func newRAMModel(total, used, percentage, free int64) *RAM {
	return &RAM{
		Total:      total,
		Used:       used,
		Percentage: percentage,
		Free:       free,
	}
}

var RamTableName string = "ram"
var RamTable string = fmt.Sprintf(`
	%s (
		id INT AUTO_INCREMENT,
		total INT,
		used INT,
		percentage INT,
		free INT,
		fecha DATETIME(6),
		PRIMARY KEY (id)
	);
`, RamTableName)

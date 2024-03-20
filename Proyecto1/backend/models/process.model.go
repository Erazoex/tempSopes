package models

import "fmt"

type ProcessState struct {
	Pid   int64  `json:"pid"`
	State string `json:"state"`
}

var ProcessStateTableName string = "processState"
var ProcessStateTable string = fmt.Sprintf(`
	%s (
		pid INT AUTO_INCREMENT,
		state VARCHAR(255),
		PRIMARY KEY (pid)
	);
`, ProcessStateTableName)

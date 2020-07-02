package models

import (
	"fmt"
	"time"

	"github.com/viswas163/UpstartProject/db"
)

// InsertRaw : Inserts the sensor into the sensors_raw_values table
func (s *Data) InsertRaw() error {
	var lastInsertID int
	err := db.GetInstance().QueryRow(`INSERT INTO sensors_raw_values(sensor_id, value, created_on)
     VALUES ((SELECT id FROM sensors WHERE name = $1), $2, $3) RETURNING id;`,
		s.Name, s.Value, time.Unix(s.Timestamp, 0)).Scan(&lastInsertID)
	panic(err)
	fmt.Println("Inserted to raw table in db! Last inserted id =", lastInsertID)
	return err
}

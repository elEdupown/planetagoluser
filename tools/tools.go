package tools

import (
	"time"
	"fmt"
)

func FechaMySQL() string{
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(),t.Hour(), t.Minute(), t.Second())
	return fecha
}
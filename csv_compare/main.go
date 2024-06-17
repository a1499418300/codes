package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func main() {
	mapV1 := make(map[string]int)
	listV1, err := parseCsv[RotateLoginPool]("rotate_login_pool514.csv")
	if err != nil {
		log.Fatalf("解析excel出错，err: %v", err)
	}
	for _, v := range listV1 {
		mapV1[v.RobotSerialNo] = v.Status
	}

	mapV2 := make(map[string]int)
	listV2, err := parseCsv[RotateLoginPool]("rotate_login_poo514v2l.csv")
	if err != nil {
		log.Fatalf("解析excel出错，err: %v", err)
	}
	for _, v := range listV2 {
		mapV2[v.RobotSerialNo] = v.Status
	}

	out := "out.sql"
	f, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开sql文件出错，err: %v", err)
	}
	defer f.Close()
	for robot, status := range mapV1 {
		if status2, ok := mapV2[robot]; ok {
			if status2 != status {
				_, e_ := f.Write([]byte(fmt.Sprintf("UPDATE rotate_login_pool SET status = %d WHERE robot_serial_no = '%s';\n", status2, robot)))
				if e_ != nil {
					log.Fatalf("写入sql出错，err: %v", err)
				}
			}
		}
	}
}

type RotateLoginPool struct {
	RobotSerialNo string `csv:"robot_serial_no"`
	Status        int    `csv:"status"`
}

func parseCsv[T any](path string) (res []*T, err error) {
	var file *os.File
	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	if err = gocsv.UnmarshalFile(file, &res); err != nil {
		return
	}
	return
}

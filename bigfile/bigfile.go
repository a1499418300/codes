package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// `D:\数据sql\rotate_login_record.sql`
func replaceAndSave(src string, old, new string) {
	dest := getDstFileName(src)
	srcf, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcf.Close()

	destf, err := os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer destf.Close()
	writer := bufio.NewWriter(destf)

	scanner := bufio.NewScanner(srcf)
	count := 0
	for scanner.Scan() {
		// 处理每一行
		fmt.Println(scanner.Text())
		line := scanner.Text()
		line = strings.Replace(line, old, new, -1)
		_, err := writer.WriteString(line)
		if err != nil {
			panic(err)
		}

		count++
		// if count > 1000 {
		// 	break
		// }
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func getDstFileName(path string) string {
	tmp := strings.Split(path, `\`)
	ori := tmp[len(tmp)-1]
	tmp2 := strings.Split(ori, ".")
	return strings.Join(tmp[:len(tmp)-1], `\`) + `\` + tmp2[0] + "_v2.sql"
}

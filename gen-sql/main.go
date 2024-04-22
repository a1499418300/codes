package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
)

func main() {
	genRobot("robot.csv")
	genChatroom("chatroom.csv")
	genChatroomRobot("chatroom_robot.csv")
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

func genSqlPath(path string) string {
	tmp := strings.Split(path, ".")
	if len(tmp) == 1 {
		return path + ".sql"
	}
	return tmp[0] + ".sql"
}

func genChatroom(path string) {
	type Record struct {
		ChatroomId       string `csv:"chatroom_id"`
		ChatroomSerialNo string `csv:"chatroom_serial_no"`
		MemberCount      string `csv:"member_count"`
		Link             string `csv:"link"`
		PrivateLink      string `csv:"private_link"`
		CreatorId        string `csv:"creator_id"`
	}
	list, err := parseCsv[Record](path)
	if err != nil {
		log.Fatalf("解析excel出错，err: %v", err)
	}

	sqlPath := genSqlPath(path)
	f, err := os.OpenFile(sqlPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开sql文件出错，err: %v", err)
	}
	defer f.Close()

	const sqlTpl = `INSERT IGNORE INTO collect_resource_chatroom (chatroom_id, chatroom_serial_no, chatroom_name, member_count, creator_id, link, private_link, refresh_member_flag, refresh_info_flag, refresh_history_flag,remark) VALUES ('%s', '%s', '', %s, '%s', '%s', '%s', 0, 0, 0, '%s');`
	for _, v := range list {
		_, e_ := f.Write([]byte(fmt.Sprintf(sqlTpl+"\n", v.ChatroomId, v.ChatroomSerialNo, v.MemberCount, v.CreatorId, v.Link, v.PrivateLink, "20240416手动导入")))
		if e_ != nil {
			log.Fatalf("写入sql出错，err: %v", err)
		}
	}
}

func genRobot(path string) {
	type Record struct {
		RobotId       string `csv:"robot_id"`
		RobotSerialNo string `csv:"robot_serial_no"`
		Account       string `csv:"account"`
		Type          string `csv:"type"`
	}
	list, err := parseCsv[Record](path)
	if err != nil {
		log.Fatalf("解析excel出错，err: %v", err)
	}

	sqlPath := genSqlPath(path)
	f, err := os.OpenFile(sqlPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开sql文件出错，err: %v", err)
	}
	defer f.Close()

	const sqlTpl = `INSERT IGNORE INTO kfpt_module_feature.collect_resource_robot (account, robot_id, robot_serial_no, id_prefix_number, id_len, account_weight_type, instock_source, is_init, remark, type) VALUES ('%s', '%s', '%s', '%s', %d, %d, '%s', 1, '%s', "%s");`
	for _, v := range list {
		_, e_ := f.Write([]byte(fmt.Sprintf(sqlTpl+"\n", v.Account, v.RobotId, v.RobotSerialNo, v.RobotId[:1], len(v.RobotId), checkTgAccountWeightType(v.RobotId), "平台号导入", "20240416手动导入", v.Type)))
		if e_ != nil {
			log.Fatalf("写入sql出错，err: %v", err)
		}
	}
}

func genChatroomRobot(path string) {
	type Record struct {
		ChatroomId       string `csv:"chatroom_id"`
		ChatroomSerialNo string `csv:"chatroom_serial_no"`
		RobotId          string `csv:"robot_id"`
		RobotSerialNo    string `csv:"robot_serial_no"`
	}
	list, err := parseCsv[Record](path)
	if err != nil {
		log.Fatalf("解析excel出错，err: %v", err)
	}

	sqlPath := genSqlPath(path)
	f, err := os.OpenFile(sqlPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开sql文件出错，err: %v", err)
	}
	defer f.Close()

	const sqlTpl = `INSERT IGNORE INTO kfpt_module_feature.collect_resource_chatroom_robot (robot_id, robot_serial_no, chatroom_id, chatroom_serial_no, status, remark) VALUES ('%s','%s','%s','%s',10,'%s');`
	for _, v := range list {
		_, e_ := f.Write([]byte(fmt.Sprintf(sqlTpl+"\n", v.RobotId, v.RobotSerialNo, v.ChatroomId, v.ChatroomSerialNo, "20240416手动导入")))
		if e_ != nil {
			log.Fatalf("写入sql出错，err: %v", err)
		}
	}
}

// checkTgAccountWeightType 检测账号权重类型
// 长度小于10，老号
// 大于等于10 且 第一位小于6，老号
// 账号权重类型: 1.新号 2.老号
func checkTgAccountWeightType(accountId string) int {
	if len(accountId) < 10 {
		return 2 // 老号
	} else {
		firstDigit, err := strconv.Atoi(string(accountId[0]))
		if err != nil {
			// 处理转换错误
			return 1
		}
		if firstDigit < 6 {
			return 2 // 老号
		}
		return 1 // 新号
	}
}

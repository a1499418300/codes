package main

import (
	"fmt"
	"sort"
)

type Room struct {
	Need  int
	No    string
	Codes []string
}

func allocatePhones(rooms []Room, phone []string) {
	need := rooms[0].Need
	if len(phone) < need*len(rooms) {
		// 数量不足，挨个分配
		need = 1
	}
	for i := 0; i < len(rooms); i++ {
		begin := need * i
		end := begin + need
		if len(phone)-1 < end {
			rooms[i].Codes = append(rooms[i].Codes, phone[begin:end]...)
			return
		}
		rooms[i].Codes = append(rooms[i].Codes, phone[begin:end]...)
		rooms[i].Need -= need
	}
	phone = phone[need*len(rooms):]
	for i, room := range rooms {
		if room.Need > 0 && len(phone) > 0 {
			allocatePhones(rooms[i:], phone)
			return
		}
	}
}

func main() {
	phoneNumbers := []string{"123", "456", "789", "101112", "131415", "161718", "192021", "222324"}
	rooms := []Room{
		{Need: 1, No: "room1"},
		{Need: 10, No: "room2"},
		{Need: 2, No: "room3"},
		{Need: 11, No: "room4"},
		{Need: 4, No: "room5"},
	}
	sort.SliceStable(rooms, func(i, j int) bool {
		return rooms[i].Need < rooms[j].Need
	})
	allocatePhones(rooms, phoneNumbers)
	for _, v := range rooms {
		fmt.Println(v)
	}
}

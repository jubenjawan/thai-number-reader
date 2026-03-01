package main

import (
	"fmt"
	"itone/thai-number-reader/handle"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// จงเขียนโปรแกรมที่รับตัวเลขจํานวนเต็มบวก (ตั้งแต่ 0 ถึง 99,999,999,999,999) แล้วแปลงตัวเลขนั้นให้เป็น
// คําอ่านภาษาไทย โดยแสดงผลลัพธ์ในรูปแบบข้อความ และเขียนโดยภาษา Go
func main() {
	var input string
	fmt.Print("กรุณาระบุตัวเลข: ")
	fmt.Scanln(&input)

	inputInt, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Printf("[ERROR] ไม่สามารถแปลงได้")
		return
	}

	p := message.NewPrinter(language.English)
	p.Printf("ตัวเลขที่ระบุ: %d\n", inputInt)
	fmt.Println("-------------------")

	res := handle.Process(int(inputInt))
	fmt.Printf("ผลลัพธ์: %s\n", res)
}

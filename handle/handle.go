package handle

import (
	"strconv"
	"strings"
)

var digits = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var units = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}
var unitsM = []string{"", "ล้าน", "ล้านล้าน"}

// สร้างฟังก์ชันย่อย
func ConvertToThai(num int) string {
	thaiWord := ""
	if num < 0 || num > 999999 {
		return thaiWord
	}
	unitPos := 0
	for num > 0 {
		digit := num % 10

		if digit != 0 {
			word := ""
			switch unitPos {
			case 0: // หลักหน่วย
				if digit == 1 && len(strconv.Itoa(num)) == 2 {
					word = "เอ็ด"
				} else {
					word = digits[digit]
				}
			case 1: // หลักสิบ
				switch digit {
				case 1:
					word = ""
				case 2:
					word = "ยี่"
				default:
					word = digits[digit]
				}
			default:
				word = digits[digit]

			}
			thaiWord = word + units[unitPos] + thaiWord
		}
		num /= 10
		unitPos++
	}

	return thaiWord
}

// สร้างฟังก์ชันหลัก:
func Process(num int) string {
	if num == 0 {
		return digits[num]
	}

	parts := []string{}
	groupIndex := 0

	for num > 0 {
		group := int(num % 1000000)
		if group != 0 { // ไม่ถึงแสน
			text := ConvertToThai(group)
			parts = append([]string{text + unitsM[groupIndex]}, parts...)
		}
		num /= 1000000
		groupIndex++
	}

	return strings.Join(parts, "")
}

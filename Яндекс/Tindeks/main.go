package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var candidates int
	var output, line string

	fmt.Scan(&candidates)

	for i := candidates; i > 0; i-- {
		fmt.Scanln(&line)
		mas := strings.Split(line, ",")
		fio_line := strings.Join([]string{mas[0], mas[1], mas[2]}, "")
		fio := strings.Split(fio_line, "")
		fio_len := int(len(fio))
		fio_byte := []byte(fio[0])
		alpha := int((fio_byte[0] - byte('A') + 1)) * 256

		bd_day, _ := strconv.Atoi(mas[3])
		bd_month, _ := strconv.Atoi(mas[4])
		bd_code := ((bd_day/10 + bd_day%10) + (bd_month/10 + bd_month%10)) * 64

		res := fio_len + alpha + bd_code
		output = fmt.Sprintf("%X ", res)
		fmt.Printf("%.3s ", output)
	}
}

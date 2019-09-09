package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func count(rdr io.Reader, cnts *[26]int32) {
	var b [1]byte
	for {
		_, _ = rdr.Read(b[:])
		if b[0] < 'a' || b[0] > 'z' {
			break
		}
		cnts[b[0]-'a']++
	}
}

func main() {
	var chars1Cnt [26]int32
	var chars2Cnt [26]int32

	rdr := bufio.NewReaderSize(os.Stdin, 1024)

	count(rdr, &chars1Cnt)
	count(rdr, &chars2Cnt)

	for i, cnt1 := range chars1Cnt {
		if cnt1 != chars2Cnt[i] {
			fmt.Println(`0`)
			return
		}
	}

	fmt.Println(`1`)
}

/*
Даны две строки, состоящие из строчных латинских букв. Требуется определить, являются ли эти строки анаграммами, т. е. отличаются ли они только порядком следования символов.

Формат ввода
Входной файл содержит две строки строчных латинских символов, каждая не длиннее 100 000 символов. Строки разделяются символом перевода строки.

Формат вывода
Выходной файл должен содержать единицу, если строки являются анаграммами, и ноль в противном случае.

Ввод:
qiu
iuq

Вывод:
1
*/

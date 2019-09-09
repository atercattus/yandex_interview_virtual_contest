package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	var maxLength, curLength int

	oneBytes := []byte(`1`)

	scan.Scan()
	n, _ := strconv.Atoi(scan.Text())

	for i := 0; i < n; i++ {
		scan.Scan()
		isOne := bytes.Equal(oneBytes, scan.Bytes())

		if isOne {
			curLength++
		} else {
			if curLength > maxLength {
				maxLength = curLength
			}
			curLength = 0
		}
	}

	if curLength > maxLength {
		maxLength = curLength
	}

	fmt.Println(maxLength)
}

/*
Требуется найти в бинарном векторе самую длинную последовательность единиц и вывести её длину.
Желательно получить решение, работающее за линейное время и при этом проходящее по входному массиву только один раз.

Формат ввода
Первая строка входного файла содержит одно число n, n ≤ 10000. Каждая из следующих n строк содержит ровно одно число — очередной элемент массива.

Формат вывода
Выходной файл должен содержать единственное число — длину самой длинной последовательности единиц во входном массиве.

Ввод:
5
1
0
1
0
1

Вывод:
1
*/

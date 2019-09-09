package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	scan.Scan()
	n, _ := strconv.Atoi(scan.Text())

	buf := make([]byte, 2*n)

	bracketChars := [...]byte{'(', ')'}

	maxNumber := 1 << uint(2*n)
loop:
	for number := 1; number < maxNumber; number++ {
		openCnt, closeCnt := 0, 0
		for bitNum := 0; bitNum < 2*n; bitNum++ {
			bit := (number >> uint(bitNum)) & 1

			if bit > 0 {
				closeCnt++
			} else {
				if openCnt++; openCnt > closeCnt {
					continue loop
				}
			}
			buf[2*n-bitNum-1] = bracketChars[bit]
		}
		if openCnt == closeCnt {
			fmt.Printf("%s\n", buf)
		}
	}
}

/*
Дано целое число n. Требуется вывести все правильные скобочные последовательности длины 2 ⋅ n, упорядоченные лексикографически (см. https://ru.wikipedia.org/wiki/Лексикографический_порядок).

В задаче используются только круглые скобки.

Желательно получить решение, которое работает за время, пропорциональное общему количеству правильных скобочных последовательностей в ответе, и при этом использует объём памяти, пропорциональный n.

Формат ввода
Единственная строка входного файла содержит целое число n, 0 ≤ n ≤ 11

Формат вывода
Выходной файл содержит сгенерированные правильные скобочные последовательности, упорядоченные лексикографически.

Ввод:
2

Вывод:
(())
()()
*/

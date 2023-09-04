package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var separator1 = "\" + "
var separator2 = "\" - "
var separator3 = "\" * "
var separator4 = "\" / "

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		index1 := strings.Index(text, separator1)
		index2 := strings.Index(text, separator2)
		index3 := strings.Index(text, separator3)
		index4 := strings.Index(text, separator4)

		var parts [3]string

		switch {
		case strings.Contains(text, separator1):
			if len(text[1:index1]) > 10 {
				fmt.Println("Превышено число символов первого операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}

			if len(text[index1+len(separator1):]) > 12 {
				fmt.Println("Превышено число символов второго операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}
			a := text[:index1] + "\""
			operator := separator1[2:3]
			b := text[index1+len(separator1):]
			parts = [3]string{a, operator, b}
		case strings.Contains(text, separator2):
			if len(text[1:index2]) > 10 {
				fmt.Println("Превышено число символов первого операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}

			if len(text[index2+len(separator2):]) > 12 {
				fmt.Println("Превышено число символов второго операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}
			a := text[:index2] + "\""
			operator := separator2[2:3]
			b := text[index2+len(separator2):]
			parts = [3]string{a, operator, b}
		case strings.Contains(text, separator3):
			if len(text[1:index3]) > 10 {
				fmt.Println("Превышено число символов первого операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}

			if len(text[index3+len(separator3):]) > 12 {
				fmt.Println("Превышено число символов второго операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}
			a := text[:index3] + "\""
			operator := separator3[2:3]
			b := text[index3+len(separator3):]
			parts = [3]string{a, operator, b}
		case strings.Contains(text, separator4):
			if len(text[1:index4]) > 10 {
				fmt.Println("Превышено число символов первого операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}

			if len(text[index4+len(separator4):]) > 12 {
				fmt.Println("Превышено число символов второго операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}
			a := text[:index4] + "\""
			operator := separator4[2:3]
			b := text[index4+len(separator4):]
			parts = [3]string{a, operator, b}
		default:
			fmt.Println("Неверный формат ввода.")
			os.Exit(1)
		}

		a := parts[0]
		operator := parts[1]
		b := parts[2]
		bToInt, _ := strconv.Atoi(b)

		if a[0] != '"' && a[len(a)-1] != '"' {
			fmt.Println("Первый и последний символы должны быть кавычками.")
			os.Exit(1)
		}

		switch operator {

		case "+":
			if b[0] != '"' && b[len(b)-1] != '"' {
				fmt.Println("Первый и последний символы должны быть кавычками.")
				os.Exit(1)
			}

			fmt.Printf("Результат:")
			fmt.Println(a[:len(a)-1] + b[1:])

		case "-":
			if b[0] != '"' && b[len(b)-1] != '"' {
				fmt.Println("Первый и последний символы должны быть кавычками.")
				os.Exit(1)
			}

			fmt.Printf("Результат:")
			aNoQuotes := a[1 : len(a)-1]
			bNoQuotes := b[1 : len(b)-1]
			if strings.Contains(aNoQuotes, bNoQuotes) {
				fmt.Printf("Результат: \"")
				fmt.Println(strings.Replace(aNoQuotes, bNoQuotes, "", -1) + "\"")
			} else {
				fmt.Println(a)
			}

		case "*":
			result := strings.Repeat(a[1:len(a)-1], bToInt)
			if b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9" ||
				b == "10" && len(result) < 40 {
				fmt.Printf("Результат: \"%s\"\n", result)
			} else if b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9" ||
				b == "10" && len(result) > 40 {
				fmt.Printf("Результат: \"")
				fmt.Println(result[0:41] + "..." + "\"")
			} else {
				fmt.Println("при умножении b должно быть числом от 1 до 10")
				os.Exit(1)
			}

		case "/":
			if b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9" ||
				b == "10" {
				aNoQuotes := a[1 : len(a)-1]
				fmt.Printf("Результат: \"")
				fmt.Printf(aNoQuotes[:len(aNoQuotes)/bToInt] + "\"\n")
			} else {
				fmt.Println("при делении b должно быть числом от 1 до 10")
				os.Exit(1)
			}
		}
	}
}

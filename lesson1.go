package main

import (
	"awesomeProject/packageTest"
	"fmt"
	"github.com/avkgimbarr/ModuleTest"
	"os"
	"strconv"
	"strings"
)

func main() {

	for {
		var text, _ = packageTest.Reader.ReadString('\n')
		text = strings.TrimSpace(text)

		var index1 = strings.Index(text, ModuleTest.Separator1)
		var index2 = strings.Index(text, ModuleTest.Separator2)
		var index3 = strings.Index(text, ModuleTest.Separator3)
		var index4 = strings.Index(text, ModuleTest.Separator4)

		switch {
		case strings.Contains(text, ModuleTest.Separator1):
			if len(text[1:index1]) > 10 {
				fmt.Println("Превышено число символов первого операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}

			if len(text[index1+len(ModuleTest.Separator1):]) > 12 {
				fmt.Println("Превышено число символов второго операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}
			a := text[:index1] + "\""
			operator := ModuleTest.Separator1[2:3]
			b := text[index1+len(ModuleTest.Separator1):]
			ModuleTest.Parts = [3]string{a, operator, b}

		case strings.Contains(text, ModuleTest.Separator2):
			if len(text[1:index2]) > 10 {
				fmt.Println("Превышено число символов первого операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}

			if len(text[index2+len(ModuleTest.Separator2):]) > 12 {
				fmt.Println("Превышено число символов второго операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}
			a := text[:index2] + "\""
			operator := ModuleTest.Separator2[2:3]
			b := text[index2+len(ModuleTest.Separator2):]
			ModuleTest.Parts = [3]string{a, operator, b}

		case strings.Contains(text, ModuleTest.Separator3):
			if len(text[1:index3]) > 10 {
				fmt.Println("Превышено число символов первого операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}

			if len(text[index3+len(ModuleTest.Separator3):]) > 12 {
				fmt.Println("Превышено число символов второго операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}
			a := text[:index3] + "\""
			operator := ModuleTest.Separator3[2:3]
			b := text[index3+len(ModuleTest.Separator3):]
			ModuleTest.Parts = [3]string{a, operator, b}

		case strings.Contains(text, ModuleTest.Separator4):
			if len(text[1:index4]) > 10 {
				fmt.Println("Превышено число символов первого операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}

			if len(text[index4+len(ModuleTest.Separator4):]) > 12 {
				fmt.Println("Превышено число символов второго операнда. Их может быть не больше 12 с учётом кавычек.")
				os.Exit(1)
			}
			a := text[:index4] + "\""
			operator := ModuleTest.Separator4[2:3]
			b := text[index4+len(ModuleTest.Separator4):]
			ModuleTest.Parts = [3]string{a, operator, b}
		default:
			fmt.Println("Неверный формат ввода.")
			os.Exit(1)
		}

		a := ModuleTest.Parts[0]
		operator := ModuleTest.Parts[1]
		b := ModuleTest.Parts[2]
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

			fmt.Printf("Результат: ")
			fmt.Println(a[:len(a)-1] + b[1:])

		case "-":
			if b[0] != '"' && b[len(b)-1] != '"' {
				fmt.Println("Первый и последний символы должны быть кавычками.")
				os.Exit(1)
			}

			fmt.Printf("Результат: ")
			aNoQuotes := a[1 : len(a)-1]
			bNoQuotes := b[1 : len(b)-1]
			if strings.Contains(aNoQuotes, bNoQuotes) {
				fmt.Println("\"" + strings.Replace(aNoQuotes, bNoQuotes, "", -1) + "\"")
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

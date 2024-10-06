package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
/// Declare variables to store the user's choice and message.
// Объявляем переменные для хранения выбора пользователя и сообщения.
	var toEncrypt bool // true — encrypt(шифровать), false — decrypt(дешифровать)
	var encoding string // selected cipher (ROT13 or Reverse)
	var message string // the message the user will enter

	/// Call the getInput() function, which collects the user's selection (encryption/decryption, cipher and message)
	toEncrypt, encoding, message = getInput()
	
	// В зависимости от выбранного шифра и операции (шифрование или дешифрование), программа выполняет соответствующую функцию.
	// Depending on the selected cipher and operation (encryption or decryption), the program performs the corresponding function.
	switch encoding {
	case "ROT13":
		if toEncrypt {
			// If encryption is selected, call the encrypt_rot13() function
			// Если выбрано шифрование, вызываем функцию encrypt_rot13()
			fmt.Println(encrypt_rot13(message))
		} else {
			// Если выбрано дешифрование, вызываем функцию decrypt_rot13()
			// If decryption is selected, call the decrypt_rot13() function
			fmt.Println(decrypt_rot13(message))
		}

	case "Reverse":
		if toEncrypt {
			// Шифрование с использованием "Reverse" (на данный момент не реализовано)
			// Encryption using "Reverse" (not implemented at the moment)
			fmt.Println(encrypt_reverse(message))
		}else {

			// Дешифрование с использованием "Reverse" (на данный момент не реализовано)
			// Decryption using "Reverse" (not implemented at the moment)
			fmt.Println(decrypt_reverse(message))
		}


	}

}

func getInput() (toEncrypt bool, encoding string, message string) {
	var operationInput string
	var cypherInput string
	toEncrypt = false
	encoding = ""

	// Ввод операции шифрования или дешифрования-Input of encryption or decryption operation
	for {
		// // Select operation: encrypt or decrypt
		// Выбор операции: шифровать или дешифровать
		fmt.Println("\nSelect operation (1/2):")
		fmt.Println("1. Encrypt.")
		fmt.Println("2. Decrypt.")
		fmt.Scan(&operationInput)

		if operationInput == "1" {
			toEncrypt = true
			break
		} else if operationInput == "2" {
			toEncrypt = false
			break
		} else {
			fmt.Println("Invalid input. Please select 1 or 2")
		}	}

	// Ввод выбора шифра
	// Enter cipher selection
	for {
		fmt.Println("\nSelect cypher (1/2):")
		fmt.Println("1. ROT13.")
		fmt.Println("2. Reverse.")
		fmt.Scan(&cypherInput)

		if cypherInput == "1" {
			encoding = "ROT13"
			break
		} else if cypherInput == "2" {
			encoding = "Reverse" // Use toROT13 to choose between ciphers // Используем toROT13 для выбора между шифрами
			break
		} else {
			fmt.Println("Invalid input. Please select 1 or 2")
		}
	}
	// enter message from custum
	// Ввод сообщения от пользователя
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nEnter the message:")
	message, _ = reader.ReadString('\n')
	// Обрезаем возможные пробелы в начале и конце
	// Trim possible spaces at the beginning and end
	message = message[:len(message)-1] // Remove the newline character // Убираем символ новой строки 
	


	return
}
// Encrypt the message with rot13
func encrypt_rot13(s string) string {
    result := ""
    var ch rune
    for i := 0; i < len(s); i++ {
        ch = rune(s[i])// rune(s[i]) Convert the symbol to the rune type/Преобразуем символ в тип rune

        if 'a' <= ch && ch <= 'z' { //Если символ — маленькая буква If the symbol is a lowercase letter 
            ch += 13
            if ch > 'z' {
                ch -= 26 // If after the shift we go beyond the alphabet, we go back  Если после сдвига вышли за пределы алфавита, возвращаемся назад
				
            }
        }
        // checks is it big letter and after shiftin is it bigger then 'Z'
		// Если символ — заглавная буква
        // it its bigger them
        if 'A' <= ch && ch <= 'Z' {
            ch += 13
            if ch > 'Z' {
                ch -= 26
            }
        }
        result += string(ch) // Add the encrypted character to the result  Добавляем зашифрованный символ в результат
    }
    return result

	} 
	// Decrypt the message with rot13
	// Дешифрование с использованием ROT13
func decrypt_rot13(s string) string {
    result := ""
    var ch rune
    for i := 0; i < len(s); i++ {
        ch = rune(s[i])

        if 'a' <= ch && ch <= 'z' {
            ch += 13
            if ch > 'z' {
                ch -= 26
            }
        }
        // checks is it big letter and after shiftin is it bigger then 'Z'
        // it its bigger them
        if 'A' <= ch && ch <= 'Z' {
            ch += 13
            if ch > 'Z' {
                ch -= 26
            }
        }
        result += string(ch)
    }
    return result
}
 func encrypt_reverse(s string) string {
	s = ""
	   return " Is not ready yer you ****** "
   }
   
   func decrypt_reverse(s string) string {
   
	   s = ""
	   return " Is not ready yer you ****** "
   }
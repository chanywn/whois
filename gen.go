package main

import (
	"os"
)

var fileSuccFile *os.File

func main() {
	fileSuccFile, _ = os.Create("./a.txt")

	words := []string{"z","x","c","v","b","n","m","l","k","j","h","g","f","d","s","a","q","w","e","r","t","y","u","i","o","p"}
	// nums := []string{"1","2","3","4","5","6","7","8","9","0"}
	
	// for _, word := range nums {
	// 	str1 := word + "cat\n";
	// 	fileSuccFile.WriteString(str1)
	// }

	for _, word := range words {
		for _, word1 := range words {
			str2 := word + word1 +"lang\n";
						fileSuccFile.WriteString(str2)
			// for _, word2 := range words {
			// 	for _, word3 := range words {
			// 		for _, word4 := range words {
			// 			str2 := word + word1 + word2 +  word3 + word4 +"\n";
			// 			fileSuccFile.WriteString(str2)
			// 		}
			// 	}
			// }
		}
	}

	// for _, num := range nums {
	// 	str3 := num + "cat\n";
	// 	fileSuccFile.WriteString(str3)
	// }
}
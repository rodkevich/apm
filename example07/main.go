package main

type MyString string

// func (s MyString) IsUpperCase() bool {
// 	for _, r := range s {
// 		// if  ! unicode.IsSpace(r) && !unicode.IsUpper(r){
// 		// 	return false
// 		// }
// 		if unicode.IsLower(r) {
// 			return false
// 		}
// 	}
// 	return true
// }

func (s MyString) IsUpperCase() bool {
	for _, letter := range s {
		if int(letter) >= 97 && int(letter) <= 122 {
			return false
		}
	}
	return true
}

func main() {
	var V MyString = "WHAT DOES THE FOX SAY"

	print(V.IsUpperCase())
}

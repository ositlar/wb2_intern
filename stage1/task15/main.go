package main

func main() {
	someFunc()
}

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/

//Данный фрагмент кода может привести к утечке памяти,
//так как переменная v создает огромную строку, а затем в переменную justString записывается только ее первые 100 символов.
//Остальная часть строки остается в памяти и не освобождается, что может привести к переполнению памяти.

func someFunc() {
	v := createHugeString(1 << 10)
	justString := make([]byte, 100)
	copy(justString, v[:100])
	v = ""
	//...
}

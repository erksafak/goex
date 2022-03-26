package main //DOSYAYA BYTE EKLEME

import (
	"fmt"
	//"io/ioutil"
	"log"
	"os"
)

var ek, son string

func main() {
	fmt.Print("Dosyaya yazılacak kelimeyi ekle:")
	fmt.Scanf("%v", &ek)
	ek = ek + "\n"
	file, err := os.OpenFile("demo1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte(ek)
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.", byteWritten)
	//ekmtn()
}

/* YAZMANIN KISA YOLU ANCAK ÖNCEKİ VERİYİ SİLİP ÜSTÜNE YAZIYOR

func ekmtn() {
	err := ioutil.WriteFile("demo1.txt", []byte("MERHABA"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
*/

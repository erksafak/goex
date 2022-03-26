package main //DOSYA OLUŞTURMA / DOSYA BİLGİSİ EDİNME / DOSYA TAŞIMA VE ADLANDIRMA

import (
	//"fmt"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	newFile  *os.File
	fileInfo *os.FileInfo
	err      error
)

func main() {
	//
	newFile, err = os.Create("demo1.txt") // Dosya oluşturursun.     ( =  kullanıldı)
	if err != nil {
		log.Fatal(err)
	}
	//
	fileInfo, err := os.Stat("demo.txt") //Dosya adı bilgisine ve dosya bilgilerine ulaşırsın.    ( :=  kullanıldı)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size())
	fmt.Println("Perm.:", fileInfo.Mode())
	fmt.Println("Last Mod.:", fileInfo.ModTime())
	fmt.Println("Is Dic.:", fileInfo.IsDir())
	//
	time.Sleep(2 * time.Second)
	taşıveadlandır()
	fmt.Println("Dosya taşıma başarılı")
	fmt.Println("---------------------------------------------------------------------------------------------------------")
}

func taşıveadlandır() {
	yer := "demo2.txt" //dosya isimlerini kontrol et, yoksa hata verir
	yeniyer := "./moved/demo3.txt"
	err := os.Rename(yer, yeniyer) //dosyayı taşıdık, aynı zamanda isim değiştirebilirsin
	fmt.Println("---------------------------------------------------------------------------------------------------------")
	if err != nil {
		log.Fatal(err)
	}
}

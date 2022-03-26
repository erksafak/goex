package main

import (
	"fmt"
	"time"
)

func main() {
	var ad string
	p := fmt.Println
	p("----------------------------------------------------------")
	fmt.Print("isim:")
	fmt.Scanf("%s", &ad)
	o := time.Date(1987, 4, 23, 0, 0, 0, 0, time.UTC)
	e := time.Date(1985, 12, 24, 0, 0, 0, 0, time.UTC)
	t := time.Date(2015, 12, 9, 0, 0, 0, 0, time.UTC)
	z := time.Date(2019, 3, 1, 0, 0, 0, 0, time.UTC)
	if ad == "osman" {
		fmt.Printf("osman'nın doğum günü: %v\n", o)
		now := time.Now()
		süre := now.Sub(o)
		//süre:= time.Now().Sub(o)
		p("Doğumdan itibaren geçen süre:", süre)
		p(o.Month())
	} else if ad == "elmas" {
		fmt.Printf("elmas'ın doğum günü: %v\n", e)
		süre := time.Now().Sub(e)
		p("Doğumdan itibaren geçen süre:", süre)
		p(e.Month())
	} else if ad == "tarık" {
		fmt.Printf("tarık'ın doğum günü: %v\n", t)
		süre := time.Now().Sub(t)
		p("Doğumdan itibaren geçen süre:", süre)
		p(t.Month())
	} else if ad == "zeynep" {
		fmt.Printf("zeynep'in doğum günü: %v\n", z)
		süre := time.Now().Sub(z)
		p("Doğumdan itibaren geçen süre:", süre)
		p(z.Month())
	}
	p("----------------------------------------------------------")
}

/*fmt.Printf("time:%v\n",time.Now())
	t:=time.Now()
	s:=time.Date(1987,4,23,9,30,26,64,time.UTC)
	fmt.Printf("saniye=%d\n",t.Second())
    time.Sleep(time.Second)
	v:=time.Now()
	fmt.Printf("saniye=%v\n",v.Second())
	fmt.Printf("ay=%s\n%d.ay\nh.gün=%s\na.gün=%d\n",v.Month(),v.Month(),v.Weekday(),v.Day())
	tomorrow:=s.AddDate(-23,0,3)
	fmt.Println(tomorrow)
	}
*/

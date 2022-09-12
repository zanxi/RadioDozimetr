package main

import (
	"fmt"
	"math"
	"time"
)


type RadioDozimetr struct {
	Ntime   int     //Время накопления дозы пульта
	Value   float64 //Текущеее значение дозы пульта
	Pogr    float64 //Погрешность в % текущего значения
	SumDoza float64 //Накопленная доза за время Ntime
}

func (d *RadioDozimetr) GetTime() (hour int, min int, sec int) {
	hour = d.Ntime / 3600
	min = (d.Ntime - hour*3600) / 60
	sec = d.Ntime - (hour*3600 + min*60)
	return hour, min, sec
}

func getBit(xd byte, n uint) int {
	m := 1 << n
	if xd&byte(m) != 0 {
		return 1
	}
	return 0
}

func xD(xd byte) (izB byte, izC float64, izD string) {
	izB = byte((getBit(xd, 7) << 1) | (getBit(xd, 6)))
	c := (getBit(xd, 5) << 1) | (getBit(xd, 4))
	d := (getBit(xd, 3) << 3) | (getBit(xd, 2) << 2) | (getBit(xd, 1) << 1) | (getBit(xd, 0))
	switch d {
	case 0:
		izD = "%"
	case 1:
		izD = "%"
	case 2:
		izD = "мкЗв"
	case 3:
		izD = "мкЗв/ч"
	case 4:
		izD = "1(с*см^2)"
	default:
		izD = "********"
	}
	switch c {
	case 0:
		izC = 1.0
	case 1:
		izC = 1e3
	case 2:
		izC = 1e6
	default:
		izC = 1.0

	}
	return izB, izC, izD
}

func getFloat(buf []byte, pos int) float64 {
	result := 0.0
	s := 1
	if buf[pos] > 127 {
		s = -1
	}
	p := int16(buf[pos])
	if p > 127 {
		p -= 128
	}
	p -= 63

	m := (float64(uint16(buf[pos+1])*256+uint16(buf[pos+2])) / 65536) + 1
	result = float64(s) * m * math.Pow(2.0, float64(p))
	return result
}


func crcCalc(buf []byte) bool {
	var crc uint16
	crc = 0xffff
	for i := 0; i < len(buf)-2; i++ {
		crcL := (crc & 0xff) ^ uint16(buf[i])
		crc = (crc & 0xff00) | (crcL & 0xff)
		for j := 0; j < 8; j++ {
			if crc&1 > 0 {
				crc ^= 0xa001
			}

		}
	}
	fmt.Println(crc&0xff, (crc&0xff00)>>8)
	if (crc&0xff != uint16(buf[len(buf)-2])) || ((crc&0xff00)>>8 != uint16(buf[len(buf)-1])) {
		return false
	}
	return true
}


func worker(done chan bool){
   fmt.Print("working...")
   fmt.Print(time.Second)

   time.Sleep(time.Second)
   fmt.Println("done")

   done <- true
}

func main(){
  
  defer fmt.Println("Bye !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

   fmt.Println("Hi World!!!")

   var a =10
   var b = 5
   var res int
   res = a+b

   var web string = "ittttt"

   fmt.Println("res = ", res)

   fmt.Println("string = ", web)

   var num float64 = 4.675
   fmt.Println("%T : ", num)

   var age = 20

   if age<15{
     fmt.Println("Children")
   } else if age==20{
     fmt.Println("SchoolChildren")
   }


  switch age {
  case 5: fmt.Println("Case --- Child")
  case 10: fmt.Println("Case --- School")
  case 20: fmt.Println("Case --- Student")
  }

  var i = 10
  for i>0 {
    fmt.Println("i = ", i)
    i--
  }

  for i:=0; i<=5; i++{
    fmt.Println("iii = ", i);
  }

  var arr[3] int
  arr[0] = 45
  arr[1] = 97
  arr[2] = 76

  fmt.Println("arr[1] = ", arr[1])

  nums := [3]float64 {4.23, 5.23, 98.1}

  for i, value := range nums{
    fmt.Println("value = ",value, "; i=",i)
  }

  websites := make(map[string] float64)
  websites["qwerty"] = 54.54645
  websites["adin"] = 423.545
  fmt.Println("float = ",websites["adin"])

  var r int
  r= summ(a,b)
  fmt.Println("r = ", r)

  done:=make(chan bool,1)
  go worker(done)
  <- done



}

func summ(a_ int, b_ int) int{
  var res int
  res = a_ + b_
  return res
}

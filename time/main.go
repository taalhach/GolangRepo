package main

import (
	"time"
	"fmt"
	"math"
	"strconv"
)

func main() {
	t:=1542458820
	fmt.Println(time.Unix(int64(t),0))
	//t := time.Now()
	//x := t.Unix()
	//y := t.UnixNano()
	//c := int64(math.Mod(float64(x), 60))
	//fmt.Println("c: ", c)
	//fmt.Println(x, "dif: ", x-c)
	//d := time.Unix(x, y)
	//fmt.Println(d)
	//d = time.Unix(x, 0)
	//year, _ := d.ISOWeek()
	////date := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
	////fmt.Println("da",date)
	////fmt.Println(date.ISOWeek())
	//fmt.Println(time.Now().Weekday() == time.Monday)
	////fmt.Println(fweekDay(year, m, time.UTC))
	//yer,week:=t.ISOWeek()
	//fmt.Println("wow",t)
	//date:=time.Date(year,time.October,0,0,0,0,0,time.UTC)
	//yer,week=date.ISOWeek()
	//log.Println(yer,week,date.Month(),date,date.Day())
	//idontknow(yer,date.Month(),week)
	////fmt.Println(firstDayOfISOWeek(yer,week,time.UTC))
	//month()
}

func month()  {
	mnth,_:=strconv.ParseInt(fmt.Sprintf("%d",time.Now().Month()),10,64)
	c:=(math.Ceil(float64(mnth)/float64(5))*float64(5))-float64(5-1)
	fmt.Println(mnth,time.Now().Month(),c)

}
func idontknow(year int, month time.Month, week int)  {
	date:=time.Date(year,month,0,0,0,0,0,time.UTC)
	_,w:=date.ISOWeek()
	fmt.Println(w,week)
	for w<week{
		date=date.AddDate(0,0,1)
		_,w=date.ISOWeek()

	}
	fmt.Println(date,"the fuck")

}
//func firstDayOfISOWeek(year int, week int, timezone *time.Location) time.Time {
//	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
//	isoYear, isoWeek := date.ISOWeek()
//	for date.Weekday() != time.Monday { // iterate back to Monday
//		date = date.AddDate(0, 0, -1)
//		isoYear, isoWeek = date.ISOWeek()
//	}
//	for isoYear < year { // iterate forward to the first day of the first week
//		date = date.AddDate(0, 0, 1)
//		isoYear, isoWeek = date.ISOWeek()
//	}
//	for isoWeek < week { // iterate forward to the first day of the given week
//		date = date.AddDate(0, 0, 1)
//		isoYear, isoWeek = date.ISOWeek()
//	}
//	return date
//}

//func fweekDay(year, week int, timezone *time.Location) time.Time {
//	fmt.Println("s ", year, week)
//	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
//	fmt.Println("date: ",date)
//	iyear, iweek := date.ISOWeek()
//	fmt.Println("O loop", date)
//	for date.Weekday() != time.Monday {
//		date = date.AddDate(0, 0, -1)
//		iyear, iweek = date.ISOWeek()
//		fmt.Print(date.Weekday(), " ")
//	}
//	fmt.Println("D loop", date)
//	j := 0
//	for iyear < year {
//		date = date.AddDate(0, 0, 1)
//		iyear, iweek = date.ISOWeek()
//		fmt.Print(j, " ")
//		j++
//	}
//	fmt.Println("l loop", date, iweek, week)
//	i := 0
//	for iweek < week {
//		date = date.AddDate(0, 0, 1)
//		iyear, iweek = date.ISOWeek()
//		fmt.Print(i, " ")
//		i++
//	}
//	fmt.Println(date)
//	return date
//
//}

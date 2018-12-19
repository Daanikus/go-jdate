// Copyright 2014 mohammadreza hasanzadeh.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"
)

type Date struct {
	year, month, day int
}

var g_days_in_month = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var j_days_in_month = []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

func (p *Date) AdToJalali(year int, month int, day int) string {

	p.day = day
	p.month = month
	p.year = year

	gy := p.year - 1600
	gm := p.month - 1
	gd := p.day - 1

	g_day_no := 365*gy + (gy+3)/4 - (gy+99)/100 + (gy+399)/400

	for i := 0; i < p.month-1; i++ {
		g_day_no += g_days_in_month[i]

	}
	if gm > 1 && ((gy%4 == 0 && gy%100 != 0) || (gy%400 == 0)) {
		g_day_no += 1
	}
	// leap and after Feb
	g_day_no += gd

	j_day_no := g_day_no - 79

	j_np := j_day_no / 12053
	j_day_no %= 12053
	jy := 979 + 33*j_np + 4*int(j_day_no/1461)

	j_day_no %= 1461

	if j_day_no >= 366 {
		jy += (j_day_no - 1) / 365
		j_day_no = (j_day_no - 1) % 365
	}

	var ix int
	for i := 0; i < 11; i++ {
		if !(j_day_no >= j_days_in_month[i]) {
			i -= 1
			break
		}
		j_day_no -= j_days_in_month[i]
		ix = i
	}

	jm := ix + 2
	jd := j_day_no + 1

	return fmt.Sprintf("[%d] | [%d] | [%d]", jy, jm, jd)

}

func (p *Date) JalaliToAd(year int, month int, day int) string {
	p.day = day
	p.month = month
	p.year = year

	jy := p.year - 979
	jm := p.month - 1
	jd := p.day - 1

	j_day_no := 365*jy + int(jy/33)*8 + (jy%33+3)/4

	for i := 0; i < jm; i++ {
		j_day_no += j_days_in_month[i]

	}

	j_day_no += jd

	g_day_no := j_day_no + 79

	gy := 1600 + 400*int(g_day_no/146097) // 146097 = 365*400 + 400/4 - 400/100 + 400/400
	g_day_no = g_day_no % 146097

	if g_day_no >= 36525 { // 36525 = 365*100 + 100/4
		g_day_no -= 1
		gy += 100 * int(g_day_no/36524) //  36524 = 365*100 + 100/4 - 100/100
		g_day_no = g_day_no % 36524

		if g_day_no >= 365 {
			g_day_no += 1
		} else {
			_ = 0

		}
	}

	gy += 4 * int(g_day_no/1461) // 1461 = 365*4 + 4/4
	g_day_no %= 1461

	if g_day_no >= 366 {
		_ = 0
		g_day_no -= 1
		gy += g_day_no / 365
		g_day_no = g_day_no % 365
	}
	i := 0
	for g_day_no >= g_days_in_month[i] {

		if !(g_day_no >= g_days_in_month[i]) {
			i -= 1
			break
		}
		g_day_no -= g_days_in_month[i]

		i += 1
	}

	gmonth := i + 1
	gday := g_day_no + 1
	gyear := gy

	return fmt.Sprintf("[%d] | [%d] | [%d]", gyear, gmonth, gday)

}



/*
you can get current date in jalali by this function .
rename package name to calender and use it like fmt.print(calender.JdateNOW())
to print out current date in jalali.
*/
func JdateNow() string {
	currentTime := time.Now()
	val := &Date{
		year:  currentTime.Year(),
		month: int(currentTime.Month()),
		day:   currentTime.Day(),
	}

	return val.AdToJalali(currentTime.Year(),int(currentTime.Month()),currentTime.Day())
}

func main() {

	currentTime := time.Now()
	val := &Date{
		year:  currentTime.Year(),
		month: int(currentTime.Month()),
		day:   currentTime.Day(),
	}

	fmt.Println(val.JalaliToAd(1397, 9, 27))  // will print [2018] | [12] | [18]
	fmt.Println(val.AdToJalali(2018, 12, 20)) //will print [1397] | [9]  | [29]

	fmt.Println(JdateNow()) // will print current day in jalali date

}

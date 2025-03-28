package main

import (
	"fmt"
	"strings"
)

type PersianDate struct {
	FORMAT string

	latinNumbers       []string
	persianNumbers     []string
	persianMonths      []string
	persianShortMonths []string
	persianDays        []string
	persianShortDays   []string
	persianSeasons     []string
}

// NewPersianDate creates a new PersianDate object
func NewPersianDate(format string) *PersianDate {
	persianNumbers := []string{"۰", "۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹"}

	latinNumbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	persianMonths := []string{"فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور", "مهر", "آبان", "آذر", "دی", "بهمن", "اسفند"}

	persianShortMonths := []string{"فر", "ار", "خر", "تی‍", "مر", "شه‍", "مه‍", "آب‍", "آذ", "دی", "به‍", "اس‍"}

	persianDays := []string{"یکشنبه", "دوشنبه", "سه شنبه", "چهارشنبه", "پنج شنبه", "جمعه", "شنبه"}

	persianShortDays := []string{"ی", "د", "س", "چ", "پ", "ج", "ش"}

	persianSeasons := []string{"بهار", "تابستان", "پاییز", "زمستان"}

	return &PersianDate{FORMAT: format, persianNumbers: persianNumbers, latinNumbers: latinNumbers, persianMonths: persianMonths, persianShortMonths: persianShortMonths, persianDays: persianDays, persianShortDays: persianShortDays, persianSeasons: persianSeasons}
}

func (p *PersianDate) Jalali() string {
	return p.FORMAT
}

// Detect wheter if given persian year is leap year (kabiseh) or not
func (p *PersianDate) IsLeapYear(year int) bool {
	yearsInCycle := year % 33
	remaineders := []int{1, 5, 9, 13, 17, 21, 25, 29, 0}
	for _, remaineder := range remaineders {
		if yearsInCycle == remaineder {
			return true
		}
	}
	return false
}

// Converts Latin number like 1402 to Persian numbers like ۱۴۰۲
func (p *PersianDate) ToPersianNumbers(text string) string {
	for i, value := range p.latinNumbers {
		text = strings.ReplaceAll(text, value, p.persianNumbers[i])
	}
	return text
}

// Converts Persian numbers to Latin numbers like ۱۴۰۲ to 1402
func (p *PersianDate) ToLatinNumbers(text string) string {
	for i, value := range p.persianNumbers {
		text = strings.ReplaceAll(text, value, p.latinNumbers[i])
	}
	return text
}
func main() {
	persianDate := NewPersianDate("%d/%m/%Y")

	fmt.Println("1402", persianDate.IsLeapYear(1402))
	fmt.Println("1403", persianDate.IsLeapYear(1403))
	fmt.Println("1404", persianDate.IsLeapYear(1404))
	fmt.Println("1405", persianDate.IsLeapYear(1405))
	fmt.Println("1406", persianDate.IsLeapYear(1406))
	fmt.Println("1407", persianDate.IsLeapYear(1407))
	fmt.Println("1408", persianDate.IsLeapYear(1408))
	fmt.Println("1409", persianDate.IsLeapYear(1409))
	fmt.Println("1410", persianDate.IsLeapYear(1410))

	fmt.Println(persianDate.ToPersianNumbers("1402/01/01 salam"))
	fmt.Println(persianDate.ToLatinNumbers("۱۴۰۲/۰۱/۰۱ سلام"))
}

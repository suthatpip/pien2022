package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"time"

	"github.com/aquilax/truncate"
	"github.com/nleeper/goment"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var month = []string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"}
var week = []string{"อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัสบดี", "ศุกร์", "เสาร์"}

func RandSeq(n int) string {
	rand.Seed(time.Now().UnixNano() + RandInt(1, 1000))
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// func ParseMilliTimestamp(tm int64) time.Time {
// 	sec := tm / 1000
// 	msec := tm % 1000
// 	return time.Unix(sec, msec*int64(time.Millisecond))
// }

func RandInt(min int, max int) int64 {
	return int64(min + rand.Intn(max-min))
}

// func CreateShortCode(n int) string {
// 	return RandSeq(n)
// }

// func StringToBytes(s string) []byte {
// 	return *(*[]byte)(unsafe.Pointer(
// 		&struct {
// 			string
// 			Cap int
// 		}{s, len(s)},
// 	))
// }

// // BytesToString converts byte slice to string without a memory allocation.
// func BytesToString(b []byte) string {
// 	return *(*string)(unsafe.Pointer(&b))
// }

// func ErrorToString(err error) string {
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return ""
// }

func ToString(str interface{}) string {

	switch reflect.ValueOf(str).Kind() {
	case reflect.Slice:
		e, err := json.Marshal(str)
		if err != nil {
			return ""
		}
		return string(e)
	case reflect.Map:
		e, err := json.Marshal(str)
		if err != nil {
			return ""
		}
		return string(e)
	default:
		return fmt.Sprintf("%+v", str)
	}
}

func TruncateText(s string) string {
	truncated := truncate.Truncate(s, 200, "...", truncate.PositionEnd)
	return truncated
}

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func SigleLine(str string) string {
	re := regexp.MustCompile(`\r?\n|  |\t`)
	newstr := re.ReplaceAllString(str, "")
	return newstr
}

// func ToMMDDYYYY(myDateString string) (string, error) {
// 	myDate, err := time.Parse("02-01-2006", myDateString)
// 	if err != nil {
// 		return "", err
// 	}
// 	return myDate.Format("01-02-2006"), nil
// }

// func ToDatex(myDateString string) time.Time {
// 	myDate, _ := time.Parse("02/01/2006", myDateString)
// 	d := myDate.AddDate(-543, 0, 0)
// 	return d
// }

func DateTH(date string) string {
	d, _ := goment.New(date, "DD/MM/YYYY")
	return fmt.Sprintf("%v %v %v", d.Date(), month[d.Month()-1], d.Year()+543)
}

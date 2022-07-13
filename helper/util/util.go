package util

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/aquilax/truncate"
	"github.com/google/uuid"
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

var (
	suffixes [5]string
)

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func HumanFileSize(size float64) string {

	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024)
	getSize := round(math.Pow(1024, base-math.Floor(base)), .5, 2)

	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}

func GetUUID() string {
	id := uuid.New()
	return id.String()
}

func GetUniqNumber() int64 {
	uniqueNumber := (time.Now().UnixNano() / (1 << 11)) + (1 >> 9) + RandInt(1, 1000)
	return uniqueNumber
}

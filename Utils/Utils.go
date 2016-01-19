package Utils
import (
	"strconv"
	"crypto/md5"
	"encoding/hex"
)

func SliceString2SliceInterface(strInputs []string) []interface{} {
	vals := make([]interface{}, len(strInputs))
	for i, v := range strInputs {
		vals[i] = v
	}

	return vals
}
func SliceFloat2SliceInterface(strInputs []float64) []interface{} {
	vals := make([]interface{}, len(strInputs))
	for i, v := range strInputs {
		vals[i] = v
	}

	return vals
}
func SliceBool2SliceInterface(strInputs []bool) []interface{} {
	vals := make([]interface{}, len(strInputs))
	for i, v := range strInputs {
		vals[i] = v
	}

	return vals
}

func Int2String(intVal int) string {
	return strconv.Itoa(intVal)
}
func Int642String(intVal int64) string {
	return strconv.FormatInt(intVal,10)
}
func Float2String(intVal float64) string {
	return strconv.FormatFloat(intVal,'f',4,64)
}

func String2IntOrZero(strVal string) int64  {
	val,err:= strconv.ParseInt(strVal,0,64)

	if err!=nil {
		return 0
	} else {
		return val
	}
}
func String2FloatOrZero(strVal string) float64  {
	val,err:= strconv.ParseFloat(strVal,64)

	if err!=nil {
		return 0
	} else {
		return val
	}
}
func String2BoolOrFalse(strVal string) bool  {
	val,err:= strconv.ParseBool(strVal)

	if err!=nil {
		return false
	} else {
		return val
	}
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
package util

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"strconv"
)

func ToFloat64(v interface{}) float64 {
	if v == nil {
		return 0.0
	}

	switch v.(type) {
	case float64:
		return v.(float64)
	case string:
		vStr := v.(string)
		vF, _ := strconv.ParseFloat(vStr, 64)
		return vF
	default:
		panic("to float64 error.")
	}
}

func ToInt(v interface{}) int {
	if v == nil {
		return 0
	}

	switch v.(type) {
	case string:
		vStr := v.(string)
		vInt, _ := strconv.Atoi(vStr)
		return vInt
	case int:
		return v.(int)
	case float64:
		vF := v.(float64)
		return int(vF)
	default:
		panic("to int error.")
	}
}

func ToUint64(v interface{}) uint64 {
	if v == nil {
		return 0
	}

	switch v.(type) {
	case int:
		return uint64(v.(int))
	case float64:
		return uint64((v.(float64)))
	case string:
		uV, _ := strconv.ParseUint(v.(string), 10, 64)
		return uV
	default:
		panic("to uint64 error.")
	}
}

func ToBytes(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//将interface{}转换为某个指定的结构体
func ToStruct(vFrom interface{}, vTo interface{}) error {
	bytes, err := json.Marshal(vFrom)
	if err != nil {
		return errors.New("data err, please check vFrom")
	}
	err = json.Unmarshal(bytes, vTo)
	if err != nil {
		return errors.New("data err, please check vTo")
	}
	return nil
}

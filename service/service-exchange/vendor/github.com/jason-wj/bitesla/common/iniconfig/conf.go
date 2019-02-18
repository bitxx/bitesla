package iniconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//MarshalFile 将二进制编码的配置信息，保存在文件中
func MarshalFile(data interface{}, filename string) (err error) {
	result, err := Marshal(data)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(filename, result, 0755)
	return
}

//UnMarshalFile 从文件读取配置文件，然后返回给结构体
func UnMarshalFile(filename string, result interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return UnMarshal(data, result)
}

//Marshal 将结构体转为配置文件的二进制编码
func Marshal(data interface{}) (result []byte, err error) {
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() != reflect.Struct {
		err = errors.New("please pass address")
		return
	}

	var confStr []string
	valueInfo := reflect.ValueOf(data) //结构体中每个属性对应值的集合

	for i := 0; i < typeInfo.NumField(); i++ {
		sectionField := typeInfo.Field(i)
		sectionVal := valueInfo.Field(i)
		fieldType := sectionField.Type
		if fieldType.Kind() != reflect.Struct {
			continue
		}
		tagVal := sectionField.Tag.Get("ini")
		if len(tagVal) <= 0 {
			tagVal = sectionField.Name
		}

		section := fmt.Sprintf("\n[%s]\n", tagVal)
		confStr = append(confStr, section)

		for j := 0; j < fieldType.NumField(); j++ {
			keyField := fieldType.Field(j)
			fieldTagVal := keyField.Tag.Get("ini")
			if len(fieldTagVal) == 0 {
				fieldTagVal = keyField.Name
			}
			valField := sectionVal.Field(j)
			item := fmt.Sprintf("%s=%v\n", fieldTagVal, valField.Interface())
			confStr = append(confStr, item)
		}

	}

	for _, val := range confStr {
		byteValue := []byte(val)
		result = append(result, byteValue...)
	}

	return
}

//UnMarshal 将配置二进制编码转为结构体
func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	//var m map[string]string

	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass address")
		return
	}

	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}

	var lastSectionFieldName string

	for index, line := range lineArr {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if line[0] == ';' || line[0] == '#' {
			continue
		}
		if line[0] == '[' {
			//lastSectionFieldName 当前的属性名
			lastSectionFieldName, err = parseSection(line, typeStruct)
			if err != nil {
				err = fmt.Errorf("%v,lineno:%d", err, index+1)
				return
			}
			continue
		}

		err = parseItem(lastSectionFieldName, line, result)
		if err != nil {
			err = fmt.Errorf("%v lineno:%d", err, index+1)
			return
		}
	}
	return
}

//parseItem lastFieldName为属性名
func parseItem(lastFieldName, line string, result interface{}) (err error) {

	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("syntax error,line:%s", line)
		return
	}

	key := strings.TrimSpace(line[0:index])  //key不可为空
	val := strings.TrimSpace(line[index+1:]) //value可以为空

	if len(key) <= 0 {
		err = fmt.Errorf("syntax error,line:%s", line)
		return
	}

	resultValue := reflect.ValueOf(result)
	sectionValue := resultValue.Elem().FieldByName(lastFieldName) //属性名对应的属性
	sectionType := sectionValue.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field:%s must be struct", lastFieldName)
		return
	}

	keyFieldName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		tagValue := field.Tag.Get("ini")
		if tagValue == key {
			keyFieldName = field.Name
			break
		}
	}
	if len(keyFieldName) == 0 {
		return
	}
	fieldValue := sectionValue.FieldByName(keyFieldName)
	if fieldValue == reflect.ValueOf(nil) { //表示fieldValue是一个空属性，无意义
		err = fmt.Errorf("field:%s not support nil", keyFieldName)
		return
	}

	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(val)
	case reflect.Slice:
		split := strings.Split(val, ",")
		of := reflect.ValueOf(split)
		fieldValue.Set(of)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, errRet := strconv.ParseInt(val, 10, 64)
		if err != nil {
			err = errRet
			return
		}
		fieldValue.SetInt(intVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		intVal, errRet := strconv.ParseUint(val, 10, 64)
		if err != nil {
			err = errRet
			return
		}
		fieldValue.SetUint(intVal)
	case reflect.Float32, reflect.Float64:
		floatValue, errRet := strconv.ParseFloat(val, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetFloat(floatValue)
	case reflect.Bool:
		boolValue, errRet := strconv.ParseBool(val)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetBool(boolValue)
	default:
		err = fmt.Errorf("unsupport type:%v", fieldValue.Type().Kind())
	}

	return
}

func aaa(v interface{}) interface{} {
	sl := reflect.Indirect(reflect.ValueOf(v))
	typeOfT := sl.Type().Elem()

	ptr := reflect.New(typeOfT).Interface()
	s := reflect.ValueOf(ptr).Elem()
	return reflect.Append(sl, s)
}

func parseSection(line string, typeStruct reflect.Type) (fieldName string, err error) {
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error,invalid section:%s", line)
		return
	}

	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error,invalid section:%s", line)
		return
	}

	if line[0] == '[' && line[len(line)-1] == ']' {
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error,invalid section:%s", line)
			return
		}

		for i := 0; i < typeStruct.NumField(); i++ {
			field := typeStruct.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == sectionName {
				fieldName = field.Name //属性名，不是属性本身
				break
			}
		}
	}
	return
}

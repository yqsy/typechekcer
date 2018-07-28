package typechekcer

import (
	"reflect"
	"github.com/pkg/errors"
	"strings"
	"fmt"
)

// 1. 检查是否有key (多级)
// 2. 检查对应的value的类型
// 当检查slice时,可指定检查slice每个成员的类型
func CheckMapValue(map_ interface{}, key string, type_ reflect.Kind, sliceType reflect.Kind) error {
	if key == "" {
		return errors.New("error input")
	}

	keys := strings.Split(key, ".")

	var curObj = map_

	// 遍历找到最后一个key的value
	for i := 0; i < len(keys); i++ {
		if curObjMap, ok := curObj.(map[string]interface{}); !ok {
			return errors.New("obj invalid")
		} else {
			if curObj, ok = curObjMap[keys[i]]; !ok {
				return errors.New(fmt.Sprintf("key %v not exist", keys[i]))
			}
		}
	}

	if reflect.TypeOf(curObj).Kind() != type_ {
		return errors.New("type not match")
	}

	if type_ == reflect.Slice && sliceType != reflect.Invalid {
		err := CheckSliceWholeValue(curObj, sliceType)
		if err != nil {
			return err
		}
	}

	return nil
}

// 检查slice所有成员类型
func CheckSliceWholeValue(slice_ interface{}, type_ reflect.Kind) error {
	if curSlice, ok := slice_.([]interface{}); !ok {
		return errors.New("slice invalid")
	} else {

		// 判断slice内所有元素的类型
		for i := 0; i < len(curSlice); i++ {
			if reflect.TypeOf(curSlice[i]).Kind() != type_ {
				return errors.New("type not match")
			}
		}

		return nil
	}
}

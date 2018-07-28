package typechekcer

import (
	"testing"
	"io/ioutil"
	"github.com/yqsy/recipes/dht/bencode"
	"reflect"
)

func TestSimpleCheckMapValue(t *testing.T) {
	// 解析torrent bencode的吧

	torrentByte, err := ioutil.ReadFile("./1.torrent")
	if err != nil {
		t.Fatal(err)
	}

	objInterface, err := bencode.Decode(string(torrentByte))
	if err != nil {
		t.Fatal(err)
	}

	if err = CheckMapValue(objInterface, "info.name", reflect.String, reflect.Invalid); err != nil {
		t.Fatal(err)
	}

	if err = CheckMapValue(objInterface, "info.piece length", reflect.Int, reflect.Invalid); err != nil {
		t.Fatal(err)
	}

	if err = CheckMapValue(objInterface, "announce", reflect.String, reflect.Invalid); err != nil {
		t.Fatal(err)
	}

	if err = CheckMapValue(objInterface, "info.files", reflect.Slice, reflect.Map); err != nil {
		t.Fatal(err)
	}

	// key 不存在
	if err = CheckMapValue(objInterface, "info.filess", reflect.Slice, reflect.Invalid); err == nil {
		t.Fatal(err)
	}

	// key == ""
	if err = CheckMapValue(objInterface, "", reflect.Slice, reflect.Invalid); err == nil {
		t.Fatal(err)
	}

	// key == info.
	if err = CheckMapValue(objInterface, "info.", reflect.Map, reflect.Invalid); err == nil {
		t.Fatal(err)
	}
}

func TestSimpleSliceWholeValue(t *testing.T) {
	torrentByte, err := ioutil.ReadFile("./1.torrent")
	if err != nil {
		t.Fatal(err)
	}

	objInterface, err := bencode.Decode(string(torrentByte))
	if err != nil {
		t.Fatal(err)
	}

	if err = CheckMapValue(objInterface, "info.files", reflect.Slice, reflect.Map); err != nil {
		t.Fatal(err)
	}

	files := objInterface.(map[string]interface{})["info"].(map[string]interface{})["files"].([]interface{})

	for i := 0; i < len(files); i++ {
		err := CheckMapValue(files[i], "length", reflect.Int, reflect.Invalid)
		if err != nil {
			t.Fatal("err")
		}

		err = CheckMapValue(files[i], "path", reflect.Slice, reflect.String)
		if err != nil {
			t.Fatal("err")
		}

		paths := files[i].(map[string]interface{})["path"].([]interface{})

		_ = paths
	}
}

<!-- TOC -->

- [1. 说明](#1-说明)

<!-- /TOC -->

<a id="markdown-1-说明" name="1-说明"></a>
# 1. 说明

在go这样的静态语言中处理动态类型(interface{})需要时时刻刻做检查,检查的代码太多了,所以就需要一个站在高层次的包装了


例如:
```json
{
    "info": {
        "files": [
            {
                "length": 2342023084,
                "path": [
                    "The.Avengers.2012.复仇者联盟.双语字幕.HR-HDTV.AC3.1024X576.x264-人人影视制作.mkv"
                ]
            },
        ]
    }
}
```

我的bencode库解析是把二进制数据转换成go内置的数据结构map,slice,int,string,

有两个用到interface的地方
* map[string]interface{}
* []interface{}

直接获取int,string还好办,最怕就是获取[key]slice的了,需要做很多的检查

譬如,获取`path`.用了这个库只需要做如下的检查即可
```bash
# 一步检查slice,并保证slice里面每个元素都是map
if err = CheckMapValue(objInterface, "info.files", reflect.Slice, reflect.Map); err != nil {
		t.Fatal(err)
	}
    
# 一步解析到slice
files := objInterface.(map[string]interface{})["info"].(map[string]interface{})["files"].([]interface{})


# 再次一步检查slice,并保证slice里面每个元素都是string
err = CheckMapValue(files[i], "path", reflect.Slice, reflect.String)
		if err != nil {
			t.Fatal("err")
		}

# 一步解析到slice
paths := files[i].(map[string]interface{})["path"].([]interface{})

```

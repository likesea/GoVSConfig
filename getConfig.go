package main

import (
	"encoding/xml"
	"os"
	"strings"
)

func GetConfig(path string) map[string]map[string]string {
	var t xml.Token
	var err error
	var sta stack
	var envSettingMap = make(map[string]map[string]string)
	isNewEnvSetting := true
	fi, err := os.Open(path)
	decoder := xml.NewDecoder(fi)
	var innerSettingMap map[string]string
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {

		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			name := token.Name.Local
			if isNewEnvSetting && sta.Len() > 0 {
				isNewEnvSetting = false
				innerSettingMap = make(map[string]string)
				sta.Put(name)
			} else {
				sta.Put(name)
			}
		// 处理元素结束（标签）
		case xml.EndElement:
			if sta.Len() > 2 {
				v := sta.Pop()
				k := sta.Pop()
				innerSettingMap[k] = v
			} else if sta.Len() == 2 {
				k := sta.Pop()
				envSettingMap[k] = innerSettingMap
				isNewEnvSetting = true
			}
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content := string([]byte(token))
			if strings.TrimSpace(content) != "" && sta.Len() > 0 {
				//fmt.Println("___________________________________")
				//fmt.Println(content)
				sta.Put(content)
			}
		}
	}
	return envSettingMap
}

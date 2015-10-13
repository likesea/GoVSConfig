package main

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

//os.Args[0] 程序执行路径， os.Args[1] 必传，项目路径，os.Args[2]可选，当前环境 os.Args[3] 可选，是配置文件路径，默认项目路径
//下的ConfigProfile.xml文件， os.Args[4] 可选，配置文件模板扩展名，默认.tpl
func main() {
	//var absPaths = GetTmpFile("F:\\GoLearn\\src\\GoConfig", ".tpl")
	var s = os.Args
	var prjPath string
	if len(s) < 2 {
		panic(errors.New("project path is a must"))
	}
	jsonDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	//prjPath = "C:\\Users\\yanbi_000\\Documents\\Visual Studio 2013\\Projects\\WebAppConfigTest\\WebAppConfigTest"
	//prjPath = "F:\\GoLearn\\src\\GoConfig"
	//prjPath = "F:\\WebAppConfigTest\\WebAppConfigTest"
	prjPath = s[1]
	// "$(ProjectDir)" 传进来会有一个多余的引号，不知何故，C:\Users\yanbi_000\Documents\Visual Studio 2013\Projects\WebAppConfigTest\WebAppConfigTest"
	prjPath = strings.Replace(prjPath, "\"", "", 1)
	var configPath = prjPath + "\\ConfigProfile.xml"
	if len(s) == 3 {
		configPath = os.Args[2]
	}
	var currentEnv = GetEnvironment(jsonDir)
	//fmt.Println(currentEnv)
	var configMaps = GetConfig(configPath)
	//fmt.Println(configMaps)
	configMap := configMaps[currentEnv]
	//fmt.Println(configMap)
	generateConfigFile(prjPath, ".tpl", configMap)
}

func generateConfigFile(path string, ext string, maps map[string]string) {
	var absPaths = GetTmpFile(path, ext)
	//var outPutContent []string
	for _, v := range absPaths {
		newFilePath := strings.Replace(v, ext, ".config", 1)
		//fmt.Println(newFilePath)
		//创建新文件
		createdfile, err := os.Create(newFilePath)
		if err != nil {
			panic(err)
		}
		//打开文件
		file, err := os.Open(v)
		if err != nil {
			panic(err)
		}
		//写缓存
		w := bufio.NewWriter(createdfile)
		defer file.Close()
		defer createdfile.Close()
		//行读
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			//find the key in every line
			lineContent := scanner.Text()
			begin := strings.Index(lineContent, "{")
			if begin != -1 {
				end := strings.Index(lineContent, "}")
				key := lineContent[begin+2 : end]
				lineContent = strings.Replace(lineContent, lineContent[begin:end+1], maps[key], 1)
			}
			w.WriteString(lineContent + "\n")
		}
		w.Flush()
	}
}

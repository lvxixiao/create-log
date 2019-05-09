package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// 浏览器ua
var userAgents = []string{
	`Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_2 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5 Quark/2.4.2.986`,                                                                                        // 夸克
	`Mozilla/5.0 (Linux; Android 8.0; MI 6 Build/OPR1.170623.027; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/48.0.2564.116 Mobile Safari/537.36 T7/10.3 SearchCraft/2.6.3 (Baidu; P1 8.0.0)`,                                                   // 简单搜索
	`Mozilla/5.0 (Linux; Android 6.0; NEM-AL10 Build/HONORNEM-AL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.132 MQQBrowser/6.2 TBS/043906 Mobile Safari/537.36 MicroMessenger/6.6.1.1220(0x26060133) NetType/WIFI Language/zh_CN`, // 微信
	`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0)`,                                                                                                    // ie10
	`Mozilla/5.0 (Symbian/3; Series60/5.2 NokiaN8-00/012.002; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.0 Mobile Safari/533.4 3gpp-gba`,                                                                    // 塞班
	`netdisk;5.5.1;PC;PC-Windows;6.2.9200;WindowsBaiduYunGuanJia`, // 百度云
}

type resouce struct {
	url    string
	target string
	start  int
	end    int
}

func ruleResouce() []resouce {
	var res []resouce
	r1 := resouce{
		url:    "http://localhost:8888",
		target: "",
		start:  0,
		end:    0,
	}
	r2 := resouce{
		url:    "http://localhost:8888/list/${id}",
		target: "$id}",
		start:  1,
		end:    21,
	}
	r3 := resouce{
		url:    "http://localhost:8888/movie/${id}",
		target: "${id}",
		start:  1,
		end:    12924,
	}
	res = append(append(append(res, r1), r2), r3)
	return res
}

// buildURL 创建URL集合
func buildURL(res []resouce) []string {
	var list []string
	for _, resItem := range res {
		if len(resItem.url) == 0 {
			list = append(list, resItem.url)
		} else {
			for i := resItem.start; i <= resItem.end; i++ {
				urlStr := strings.Replace(resItem.url, resItem.target, strconv.Itoa(i), -1)
				list = append(list, urlStr)
			}
		}
	}
	return list
}

// makeLog 创建日志
func makeLog(current, refer, ua string) string {
	v := url.Values{}
	time := time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	v.Set("time", time)
	v.Set("url", current)
	v.Set("refer", refer)
	v.Set("ua", ua)
	// 模拟生成uid
	hasher := md5.New()
	hasher.Write([]byte(refer + ua))
	uid := hex.EncodeToString(hasher.Sum(nil))
	v.Set("uid", uid)
	paramsStr := v.Encode()

	logoTemplate := `127.0.0.1 - - ${time} "OPTIONS /dig?${paramsStr} HTTP/1.1" 200 43 "-" "${ua}" "-"`
	log := strings.Replace(logoTemplate, "${time}", time, -1)
	log = strings.Replace(log, "${paramsStr}", paramsStr, -1)
	log = strings.Replace(log, "${ua}", "ua", -1)

	return log
}

//randInt 获取随机数
func randInt(min, max int) int {
	if min > max {
		return max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min // Intn(int) [0, n)
}

func main() {

	var total = flag.Int("total", 100, "create log number")
	var filePath = flag.String("filePath", "./log.txt", "file path")
	flag.Parse() // 获得命令行参数
	res := ruleResouce()
	list := buildURL(res)
	var logStr string
	for i := 0; i <= *total; i++ {
		current := list[randInt(0, len(list))]
		refer := list[randInt(0, len(list))]
		ua := userAgents[randInt(0, len(userAgents))]
		logStr = logStr + makeLog(current, refer, ua) + "\n"
	}
	fd, err := os.OpenFile(*filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	defer fd.Close()
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		fmt.Println("writing now")
		fd.Write([]byte(logStr))
		fmt.Println("done")
	}
}

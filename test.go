package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)
// 超时时间：5秒
var client = &http.Client{Timeout: 5 * time.Second}

func Get(url string) string {
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

func main() {
	//s := "嘉兴平湖天籁琴行,嘉善和乐琴行,厦门赤澄文化传播有限公司,南京艺歌乐器有限公司,扬中市艺洲琴行有限公司,厦门市湖里区忆星艺琴行,鄂尔多斯市四季之音文化传媒有限公司,苏州贝特琴行,苏州工业园区索妮娅艺术培训有限公司,厦门市湖里区睿音曼琴行,河南晨露绿化工程有限公司,慈溪百乐艺术特长培训有限公司,杭州悦琴轩文化艺术有限公司,龙海市石码艾乐琴行,慈溪市知音琴社,北塘区星籁琴行,青岛森雅琴行有限公司,厦门音煌文化传播有限公司,厦门思明区贝雅尔教育咨询服务部,青岛爱与乐文化艺术交流中心,胶州市艺卓的琴行工作室,杭州朵儿艺术培训有限公司,海宁天乐艺术培训学校,常州金坛繁星音乐文化传播有限公司,李沧区鲁巴多教育咨询中心,杭州下城莘瀚文化艺术有限公司,上海尧音艺术中心,江苏美合天音琴行有限公司,杭州蓝音文化创意有限公司,青岛韵声文化传播有限公司"
	s :="天籁琴行,和乐琴行,赤澄琴行,艺歌琴行,艺洲琴行,星艺琴行,四季之音,贝特琴行,索妮娅艺术培训,睿音钢琴艺术中心,天韵钢琴艺术中心,百乐琴行,柒音琴行,艾樂琴行,慈溪知音琴社,星籁琴行,森雅琴行,音煌琴行,贝希钢琴,青岛爱与乐艺术学校,艺卓琴行,朵儿音乐,天乐琴行,珠江钢琴旗舰店,青岛博帆教育,莘瀚艺术,尧音艺术中心,天音琴行,蓝音音乐,韵声琴行"
	res := strings.Split(s, ",")
	for _,v:=range res{
		fmt.Println(v)
		res:= Get("https://restapi.amap.com/v3/place/text?key=a1f77298bbb820c9a75622b466115b12&keywords=" + v)
		fmt.Println(res)
	}

}

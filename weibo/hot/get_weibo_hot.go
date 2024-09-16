package hot

import (
	"data-collector/weibo/hot/dao/do"
	"data-collector/weibo/hot/vo"
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func submit() {

	resp, err := http.Get("https://weibo.com/ajax/side/hotSearch")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Bo  dy failed,err:", err)
		return
	}
	var v interface{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		return
	}
	hotInfo := v.(map[string]interface{})["data"].(map[string]interface{})
	mainText := vo.Resp{}
	err = mapstructure.Decode(hotInfo, &mainText)
	if err != nil {
		log.Fatal(err)
	}

	dbParam := do.DBConn{
		Host:   "127.0.0.1",
		Port:   5432,
		User:   "postgres",
		Pass:   "hcq10086",
		DBName: "postgres",
	}
	conn, _ := do.New(dbParam)
	t := time.Now()
	p, _ := strconv.Atoi(t.Format("20060102"))

	aimData := mainText.RealTime
	for i, _ := range aimData {
		aimData[i].Period = p
		aimData[i].CreateTime = t
	}

	conn.Table("weibo_hot").CreateInBatches(aimData, 1000)
}

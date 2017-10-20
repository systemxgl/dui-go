package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mahonia"
	"math/big"
	"net/http"
	"sort"
	"strings"
	"time"
)

/*
* 发送post请求
 */
func sendPost(url string, data string) string {
	var jsonStr = []byte(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

/*对象转成json字符串*/
func toJson(data interface{}) string {
	b, err := json.Marshal(data)

	if err != nil {
		return err.Error()
	}
	return string(b)
}

/*获取时间戳*/
func getTimestamp() string {
	t := time.Now()
	timestamp := fmt.Sprintf("%d", t.Unix())
	return timestamp
}

/*获取随机数*/
func getNonce() string {
	rand := RandInt64(100000000, 999999999)
	return fmt.Sprintf("%d", rand)
}

/*
* 生成区间随机数
 */
func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		RandInt64(min, max)
	}
	return i.Int64()
}

/*创建通用请求参数*/
func createParams() string {
	nonce := getNonce()
	timestamp := getTimestamp()
	signstr := signString(APPSECRET, nonce, timestamp)
	params := fmt.Sprintf("?appid=%s&nonce=%s&timestamp=%s&signature=%s", APPID, nonce, timestamp, signstr)
	return params
}

/*获取url*/
func getUrl(action string) string {
	return BAESURL + action + createParams()
}

/*字符串转base64*/
func stringToBase64(data string) string {
	input := []byte(mahonia.NewEncoder("gbk").ConvertString(data))
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

/*
*加密字符串
 */
func signString(appsecret string, nonce string, timestamp string) string {
	var arrTmp = []string{appsecret, nonce, timestamp}
	sort.Strings(arrTmp)
	arrStr := strings.Join(arrTmp, "")
	h := sha1.New()
	h.Write([]byte(arrStr))
	bs := h.Sum(nil)
	result := fmt.Sprintf("%x", bs)
	return result
}

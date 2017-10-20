// pro2 project main.go
package main

import (
	"fmt"
)

func main() {
	var uuid = "设备编号"
	/*
		* 用户绑定设备
		返回数据格式 {"OpenUserId":160715,"Code":200,"Message":"成功"}
	*/
	result := userBind(uuid, "321")
	fmt.Println(result)
	/*
		* 获取设备状态
		返回数据格式 {"State":0,"Code":200,"Message":"成功"}
	*/
	result = getDeviceState(uuid)
	fmt.Printf(result)

	/*
		* 提交打印内容
		返回数据格式 {"TaskId":10,"Code":200,"Message":"成功"}
	*/
	//要打印的内容
	var content = "欢迎使用对对机\n测试换行"
	var jsonContent = "[{\"Alignment\":0,\"BaseText\":\"" + stringToBase64(content) + "\",\"Bold\":0,\"FontSize\":0,\"PrintType\":0}]"
	result = printContent(uuid, jsonContent, "0") //0改成用户设备绑定返回的OpenUserId
	//	/*
	//		* 获取任务状态
	//		返回数据格式 {"State":0,"Code":200,"Message":"成功"}
	//	*/
	result = getPrintTaskState(0) //0改成任务编号
}
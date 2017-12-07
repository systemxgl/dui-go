package main

/*用户绑定设备*/
func userBind(uuid string, userId string) string {
	var url = getUrl("/home/userbind")
	entity := UserBindEntity{
		Uuid:   uuid,
		UserId: userId,
	}
	data := toJson(entity)
	return sendPost(url, data)
}

/*获取设备状态*/
func getDeviceState(uuid string) string {
	var url = getUrl("/home/getdevicestate")
	entity := DeviceStateEntity{
		Uuid: uuid,
	}
	data := toJson(entity)
	return sendPost(url, data)
}

/*提交打印内容*/
func printContent(uuid string, content string, openUserId string) string {
	var url = getUrl("/home/printcontent2")
	entity := PrintContentEntity{
		Uuid:         uuid,
		PrintContent: content,
		OpenUserId:   openUserId,
	}
	data := toJson(entity)
	return sendPost(url, data)
}
/*打印网页信息*/
func printHtmlContent(uuid string, printUrl string, openUserId string) string {
	var url = getUrl("/home/printhtmlcontent")
	entity := PrintContentEntity{
		Uuid:       uuid,
		PrintUrl:   printUrl,
		OpenUserId: openUserId,
	}
	data := toJson(entity)
	return sendPost(url, data)
}
/*获取任务状态*/
func getPrintTaskState(taskId int64) string {
	var url = getUrl("/home/getprinttaskstate")
	entity := TaskStateEntity{
		TaskId: taskId,
	}
	data := toJson(entity)
	return sendPost(url, data)
}

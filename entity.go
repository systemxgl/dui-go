package main

/*用户绑定设备实体*/
type UserBindEntity struct {
	Uuid   string
	UserId string
}

/*设备状态实体*/
type DeviceStateEntity struct {
	Uuid string
}

/*打印内容实体*/
type PrintContentEntity struct {
	Uuid         string
	PrintContent string
	OpenUserId   string
}

/*任务状态实体*/
type TaskStateEntity struct {
	TaskId int64
}

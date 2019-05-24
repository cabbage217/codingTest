package car

//第一版接口，其实并没有用到，因为并没有多种车辆，所以这层抽象可以先不做，以后有了多种车辆再进行抽象也不迟
type CarInterfaceV1 interface {
	Move(command string)
	GetPositionX()
	GetPositionY()
	GetOrientation()
}

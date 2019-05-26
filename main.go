package main

import (
	"os"
	"fmt"
	"./car"
	"./consts"
	"strconv"
	"./command"
	"./park"
	"bufio"
)

// 写了2个版本
// V1是简单版本
// V2把命令抽出来了
func main() {
	runV1()
	println()
	println()
	runV2()
}

//运行V1版
func runV1() {
	parkX := 4
	parkY := 4

	if len(os.Args) > 1 {
		parkX = parseParkParam(os.Args[1], 4)
	}
	if len(os.Args) > 2 {
		parkY = parseParkParam(os.Args[2], 4)
	}

	carV1 := car.NewCarV1(consts.DirectionNorth)
	carV1.Move("turn")
	checkCarV1State(carV1, parkX, parkY)

	carV1 = car.NewCarV1(consts.DirectionNorth)
	carV1.Move("forward")
	checkCarV1State(carV1, parkX, parkY)

	carV1 = car.NewCarV1(consts.DirectionEast)
	carV1.Move("forward")
	checkCarV1State(carV1, parkX, parkY)

	carV1 = car.NewCarV1(consts.DirectionWest)
	carV1.Move("forward")
	checkCarV1State(carV1, parkX, parkY)
}

//解析V1运行参数：car part参数
func parseParkParam(param string, defaultValue int) int {
	intTmp, e := strconv.Atoi(param)
	if e != nil {
		fmt.Printf("Error(%s) when parse param %s, will use default value(%d)\n", e.Error() , os.Args[1], defaultValue)
		intTmp = defaultValue
	} else if intTmp <= 0 {
		fmt.Printf("Invalid param (%s), will use default value(%d)\n" , os.Args[1], defaultValue)
		intTmp = defaultValue
	}
	return intTmp
}

//检查V1版命令执行结果
func checkCarV1State(car *car.CarV1, parkX, parkY int) {
	if car == nil {
		fmt.Println("Invalid car, can not check")
		return
	}

	fmt.Printf("CarV1 state: orientation: %s, x: %d, y: %d\n", car.GetOrientation(), car.GetPositionX(), car.GetPositionY())

	x := car.GetPositionX()
	y := car.GetPositionY()
	if x <= 0 || x > parkX {
		fmt.Printf("Exception, x(%d) overstep [1, %d]\n", x, parkX)
	}
	if y <= 0 || y > parkY {
		fmt.Printf("Exception, y(%d) overstep [1, %d]\n", y, parkY)
	}
}

//运行V2版
func runV2() {
	fmt.Println("Welcome!")
	command.PrintHelp()

	//保证park.Park实现了park.ParkInterface接口
	var _ park.ParkInterface = (*park.Park)(nil)
	//保证car.Car实现了car.CarInterface接口
	var _ car.CarInterface = (*car.Car)(nil)

	p := &park.Park{}
	c := &car.Car{}
	c.Init(consts.DirectionNorth)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if co, e := command.Parse(input.Text()); e != nil {
			fmt.Println("Error! ", e.Error())
			command.PrintHelp()
		} else if c != nil {

			fmt.Printf("command: %v\n", c)

			switch co.Type {
			case command.TypeQuit:
				os.Exit(0)
			case command.TypeHelp:
				command.PrintHelp()
			case command.TypeInitPark:
				p.Init(co.X, co.Y)
			case command.TypeForward:
				fallthrough
			case command.TypeTurn:
				if !p.IsInitialized() {
					fmt.Println("park has not initialized")
				} else {
					if e := c.Exec(co); e != nil {
						fmt.Printf("Error: %s\n", e.Error())
					} else {
						if !checkCarState(c, p) {
							fmt.Println("Will reset park and car.")
							p.Reset()
							c.Reset()
						}
					}
				}
			}
		}
	}
}

//检查V2版命令执行结果
func checkCarState(c *car.Car, p *park.Park) (pass bool) {
	if c == nil {
		fmt.Println("Invalid car, can not check")
		return false
	}
	if p == nil {
		fmt.Println("invalid park, can not check")
		return false
	} else if !p.IsInitialized() {
		fmt.Println("park has not initialized, can not check")
		return false
	}

	fmt.Printf("car: %v, park: %v\n", c, p)

	currentPosition := c.GetCurrentPosition()
	if currentPosition.X <= 0 {
		fmt.Printf("Exception, x(%d) overstep [1, %d]\n", currentPosition.X, p.GetMaxX())
		return false
	}
	if currentPosition.Y <= 0 {
		fmt.Printf("Exception, x(%d) overstep [1, %d]\n", currentPosition.X, p.GetMaxY())
		return false
	}
	if valid, e := p.IsPositionValid(currentPosition); e == nil {
		if !valid {
			fmt.Printf("Exception, car position(%v) overstep X[1, %d] or &[1, %d]\n", currentPosition, p.GetMaxX(), p.GetMaxY())
		}
		pass = valid
	} else {
		fmt.Printf("Error when check car position, %s\n", e.Error())
	}
	return pass
}

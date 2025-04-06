package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
)

func Write_go_code_using_go(dirName string) {

	err := os.Mkdir(dirName, 0750)
	if err != nil && os.IsExist(err) {
		log.Fatal(err)
	}
	filePath := fmt.Sprintf("%s/hello.go", dirName)

	data := []byte("package main\n\n import \"fmt\" \n\n func main() {\n\tfmt.Println(\"Hello, World!\")\n}\n")
	err = os.WriteFile(filePath, data, 0660)

	if err != nil {
		log.Fatal(err)
	}
}

func Log_ble_dron_battery_level() {

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	battery := ble.NewBatteryDriver(bleAdaptor)

	work := func() {
		gobot.Every(5*time.Second, func() {
			fmt.Println("Battery level:", battery.GetBatteryLevel())
		})
	}

	robot := gobot.NewRobot("bleBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{battery},
		work,
	)

	robot.Start()

}

func main() {

	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	tempFile := os.TempDir()
	fmt.Println(tempFile)

	/* This function will create a test dir and main.go file which
	 *  which will print Hello World
	 */
	Write_go_code_using_go("test")

	/* Try to work with BLE earbuds not not working
	 * BECAUSE this Liberary is for ony BLE device 4.0
	 */
	// Log_ble_dron_battery_level()
}

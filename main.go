package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	adb "github.com/zach-klippenstein/goadb"
)

const (
	ARG_INSTALL   string = "i"
	ARG_UNINSTALL string = "u"
)

var (
	port          = adb.AdbPort
	client        *adb.Adb
	err           error
	arg           string
	getArgSuccess bool
)

var stimulatingApks = [4]string{
	"com.android.vending", // play store
	"com.android.chrome",
	"com.google.android.youtube",
	"com.google.android.apps.youtube.music",
}

func getInstallArg() (string, bool) {
	if len(os.Args) <= 1 {
		// return "", false TODO: Uncomment me
		return ARG_INSTALL, true // TODO: Remove me
	}

	var arg string = os.Args[1]

	if arg != ARG_INSTALL && arg != ARG_UNINSTALL {
		return "", false
	}

	return arg, true
}

func uninstallApk(device *adb.Device, apk string) {
	fmt.Println("UNINSTALLING: " + apk)
	response, err := device.RunCommand("pm uninstall --user 0 " + apk)

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(response, "Success") || strings.Contains(response, "not installed for") {
		fmt.Println("UNINSTALL SUCCESS")
	} else {
		fmt.Println("UNKNOWN UNINSTALL RESULT: " + response)
	}
}

func installApk(device *adb.Device, apk string) {
	fmt.Println("INSTALLING: " + apk)
	response, err := device.RunCommand("cmd package install-existing " + apk)

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(response, "Package "+apk+" installed for user: 0") {
		fmt.Println("INSTALL SUCCESS")
	} else {
		fmt.Println("UNKNOWN INSTALL RESULT: " + response)
	}
}

func main() {
	fmt.Println("START")
	arg, getArgSuccess = getInstallArg()

	if !getArgSuccess {
		log.Fatal("Please supply either the command line arguments 'i' or 'u' to install or uninstall stim apps")
	}

	client, err = adb.NewWithConfig(adb.ServerConfig{
		PathToAdb: "./adb",
		Port:      port,
	})

	defer fmt.Println("END")
	defer client.KillServer()

	if err != nil {
		log.Fatal(err)
	}

	serverVersion, err := client.ServerVersion()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ADB SERVER VERSION:", serverVersion)

	devices, err := client.ListDevices()

	if err != nil {
		log.Fatal(err)
	}

	if len(devices) == 0 {
		log.Fatal("No devices connected")
	} else if len(devices) > 1 {
		log.Fatal("More than one device connected")
	}

	deviceInfo := devices[0]
	deviceDescriptor := adb.DeviceWithSerial(deviceInfo.Serial)
	device := client.Device(deviceDescriptor)

	fmt.Println("DEVICE:", deviceInfo.Model)

	for _, apk := range stimulatingApks {
		if arg == ARG_INSTALL {
			installApk(device, apk)

		} else if arg == ARG_UNINSTALL {
			uninstallApk(device, apk)
		}
	}

}

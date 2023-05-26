package main

import (
	"fmt"
	"log"
	"strings"

	survey "github.com/AlecAivazis/survey/v2"
	adb "github.com/zach-klippenstein/goadb"
)

const (
	OPTION_INSTALL   string = "Install"
	OPTION_UNINSTALL string = "Uninstall"
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

func getDevice(client *adb.Adb) *adb.Device {
	defer func() {
		// var err error = recover().(error)
		var rec = recover()
		var err error
		if rec != nil {
			err = rec.(error)
		}

		if err != nil {
			var errorMessage = err.Error()

			if strings.Contains(errorMessage, "runtime error: index out of range [1] with length 1") {
				log.Fatal("Failed to access device. It may not be in File Transfer mode")
			} else {
				log.Fatal(err)
			}
		}
	}()

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
	return client.Device(deviceDescriptor)
}

var cliPrompt = &survey.Select{
	Message: "Install or uninstall stimulating apps",
	Options: []string{OPTION_INSTALL, OPTION_UNINSTALL},
	Default: OPTION_INSTALL,
}

func main() {
	client, err = adb.NewWithConfig(adb.ServerConfig{
		PathToAdb: "./adb",
		Port:      port,
	})

	if err != nil {
		log.Fatal(err)
	}

	var device *adb.Device = getDevice(client)

	var cliAnswer string
	err = survey.AskOne(cliPrompt, &cliAnswer)
	if err != nil {
		log.Fatal(err)
	}

	for _, apk := range stimulatingApks {
		if cliAnswer == OPTION_INSTALL {
			installApk(device, apk)

		} else if cliAnswer == OPTION_UNINSTALL {
			uninstallApk(device, apk)
		}
	}

}

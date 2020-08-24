package conf

import (
	"io/ioutil"
	"net"
	"strings"
)

func cpuinfo() (string, error) {
	data, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func findPrefix(data, pre string) (str2 string) {
	strList := strings.Split(data, "\n")
	for i := 0; i < len(strList); i++ {
		if strings.HasPrefix(strList[i], pre) {
			return strList[i]
		}
	}
	return str2
}

func kv(str, sep string) (k, v string) {

	temp := strings.Split(str, sep)
	if len(temp) < 2 {
		return
	}
	k = strings.TrimSpace(temp[0])
	v = strings.TrimSpace(temp[1])
	return
}

func Serial() (id string) {

	info, err := cpuinfo()
	if err != nil {
		return ""
	}
	str := findPrefix(info, "Serial")
	_, id = kv(str, ":")
	return id
}

func Model() (mod string) {

	info, err := cpuinfo()
	if err != nil {
		return ""
	}
	str := findPrefix(info, "Model")
	_, mod = kv(str, ":")
	return mod
}

func Ether() string {

	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error:" + err.Error())
	}
	for _, inter := range interfaces {
		if inter.Name == "wlan0" {
			return inter.HardwareAddr.String()
		}
	}

	for _, inter := range interfaces {
		if inter.Name == "eth0" {
			return inter.HardwareAddr.String()
		}
	}

	return ""
}

func ID() string {
	return strings.Replace(Ether(), ":", "", -1)
}

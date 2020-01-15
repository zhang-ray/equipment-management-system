package main;

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/denisbrodbeck/machineid"
	"log"
	"net"
)


type DeviceInfo struct {
	ID string
	IPv4 string
}

func getMachineID() string {
	id, err := machineid.ProtectedID("github.com/zhang-ray/equipment-management-system")
	if err != nil {
		log.Fatal(err)
	}
	return id
}


func getLocalIp() (string, error){
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		log.Println("Get local IP addr failed!!!")
		return "", err
	}

	// TODO: return multiple IPv4
	// TODO: MAC and IPv6 addr
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				IpAddr := ipnet.IP.String()
				return IpAddr, nil
			}
		}
	}
	return "", fmt.Errorf("NoIPFound")
}

func main(){
	url := "http://127.0.0.1:9090/post-info"
	if len(os.Args)>1 {
		url = os.Args[1]
	}


	for ;; {
		IPv4, err := getLocalIp()
		if nil == err {
			var deviceInfo DeviceInfo
			deviceInfo.ID = getMachineID()
			deviceInfo.IPv4 = IPv4

			fmt.Println("URL:>", url)

			jsonStr, _ := json.Marshal(deviceInfo)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			fmt.Println("response Status:", resp.Status)
			fmt.Println("response Headers:", resp.Header)
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("response Body:", string(body))

		}

		time.Sleep(10*time.Second)
	}
}
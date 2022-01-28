package main

import (
	"fmt"
	"time"
	"math/rand"
)

var resource string = "x"
var resources map[string]int = map[string]int {"x": 4} // x is the resource with initial value 4
// Number of devices.
var n_devices int = 3

var devices map[string]Device =  map[string]Device {
	"A": Device{
			processName: "A", 
			processQueue: make([]string, n_devices - 1), 
			notInterestedInResource: true, 				// device A is not interested in the resource
			req_chan: make(chan ReqMessage, n_devices - 1), 
			res_chan: make(chan Reply, n_devices - 1)}, 

	"B": Device{
			processName: "B", 
			processQueue: make([]string, n_devices - 1), 
			req_chan: make(chan ReqMessage, n_devices - 1), 
			res_chan: make(chan Reply, n_devices - 1)}, 
	"C": Device{
			processName: "C", 
			processQueue: make([]string, n_devices - 1), 
			req_chan: make(chan ReqMessage, n_devices - 1), 
			res_chan: make(chan Reply, n_devices - 1)}}

// Resource request message format
type ReqMessage struct{
	resourceName string	
	processName string
	timestamp time.Time
}

// Reply message format
type Reply struct{
	processName string
	message string
}

// Device components
type Device struct{
	processName string
	notInterestedInResource bool
	processQueue[] string
	queueCount int

	req_chan chan ReqMessage
	res_chan chan Reply
}

// When device is not interested in the resource, it is designed to send OK to all incoming requests
func (d Device) notInterested(final_end chan string){
	fmt.Println(d.processName, "not interested in the resource")
	count := 0
	for _, v := range devices{
		if v.notInterestedInResource{
			count += 1
		}
	}
	if count != n_devices{
		for req := range d.req_chan{
			devices[req.processName].res_chan <- Reply{processName: d.processName, message: "OK"}
			count += 1
			if count == n_devices{
				break
			}
		}
	}
	final_end <- d.processName
}

func (d Device) modifyResourceValue(resourceName string, newValue int, final_end chan string){
	timestamp := time.Now()
	fmt.Println(d.processName, timestamp)
	req_message := ReqMessage{resourceName: resourceName, processName: d.processName, timestamp: timestamp}

	count := 0
	for i := range devices{
		if i == d.processName{
			continue
		}

		if devices[i].notInterestedInResource{
			count += 1
		}
		fmt.Println(d.processName, "sending request to device", i)
		devices[i].req_chan <- req_message
	}
	if count != n_devices - 1{
		for req := range d.req_chan{
			// fmt.Println(d.processName, "received request from", req.processName)
			
			if req.timestamp.Before(timestamp) || req.timestamp == timestamp{
				devices[req.processName].res_chan <- Reply{processName: d.processName, message: "OK"}
			} else{
				d.processQueue[d.queueCount] = req.processName
				d.queueCount += 1
			}
	
			count += 1
			if count == n_devices - 1{
				break
			}
		}
	}

	count = 0
	for res := range d.res_chan {
		fmt.Println(d.processName,"<-",res.processName, res.message)
		count += 1
		if count == n_devices - 1{
			break
		}
	}
	// Critical section
	resources[resourceName] = newValue	
	fmt.Println(d.processName, "changed", resourceName, "to value:", newValue, "at time:", timestamp)
	fmt.Println(d.processName, "has queued",d.queueCount, "requests")
	// Sends OK message after altering all the values

	for i := 0; i < d.queueCount; i += 1{
		devices[d.processQueue[i]].res_chan <- Reply{processName: d.processName, message: "OK"}
	}
	final_end <- d.processName
}

func main(){
	deviceC := devices["C"]
	
	deviceC.notInterestedInResource = true


	res := make(chan string)
	for _, device := range devices{
		if device.notInterestedInResource{
			go device.notInterested(res)
			continue
		}
		go device.modifyResourceValue(resource, rand.Intn(20), res)
	}
	count := 0
	for s := range res {
		fmt.Println(s, "completed")
		count += 1
		if count == n_devices{
			break
		}
	}
	fmt.Println("Final value of resource", resource, ":", resources[resource])
	fmt.Println("Done")
}
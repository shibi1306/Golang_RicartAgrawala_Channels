# Golang_RicartAgrawala_Channels
### Introduction
This repository implements Ricart-Agrawala Mutual Exclusion Algorithm using channels and Goroutines in Go language.

Usage of channels and Goroutines in Go language for Ricart-Agrawala algorithm can be very useful for creating Distributed System based simulators with entities like devices and resources representing the objects in the simulation.

### Default program
The default program contains three devices, A, B and C. A is not interested in accessing the resource with the line, `notInterestedInResource = true` in `Device` struct which calls a separate function 
`func notInterested(final_end chan string)` instead of the regular function `func modifyResourceValue(resourceName string, newValue int, final_end chan string)`.

Devices B and C attempt to alter the resource "x" using timestamps and communication using channels. 

The process looks similar to the image below 
![Ricart Agrawala diagram](/screenshots/ricart_agrawala_diagram.png)


### How to run
1. Clone the repository to your desired folder.
2. Change the terminal working directory to git folder.
3. Run the program using `go run ricart_agrawala_channels.go`.
4. you should get the mutual exclusion process output in the console.

(Note: The output may not always look sequential which is natural part of the Goroutines.)

### Screenshot
![Execution screenshot](/screenshots/output_screenshot.png)

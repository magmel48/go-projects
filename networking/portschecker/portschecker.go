/**
Port Scanner - Enter an IP address and a port range where the program will then attempt
to find open ports on the given computer by connecting to each of them.
On any successful connections mark the port as open.
*/
package portschecker

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const ip = "127.0.0.1"
const startPort = 1024
const endPort = 65535

func handleOpenedPorts(openedPorts []string) {
	fmt.Println("opened ports", openedPorts)
}

func CheckPorts() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	fmt.Println("started checking")
	openedPorts := make([]string, 0)

	go func() {
		<-signals
		fmt.Println()
		handleOpenedPorts(openedPorts)

		os.Exit(0)
	}()

	for i := startPort; i < endPort; i++ {
		port := strconv.Itoa(i)
		fmt.Printf("scanning %s\n", port)

		conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, port), time.Second)
		if err != nil {
			fmt.Printf("dial error %s\n", err.Error())
			continue
		}

		if conn != nil {
			openedPorts = append(openedPorts, port)

			err = conn.Close()
			if err != nil {
				fmt.Printf("close error %s\n", err.Error())
			}
		} else {
			fmt.Println("conn is nil by some reason")
		}
	}

	handleOpenedPorts(openedPorts)

	fmt.Println("finished checking")
}

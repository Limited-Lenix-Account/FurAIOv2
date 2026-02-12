package proxy

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetProxies() ([]string, int, error) {
	fmt.Println("Reading Proxies...")
	var proxyList []string

	file, err := os.Open("internal/proxies/proxies.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		proxyURL := ("http://" + parts[2] + ":" + parts[3] + "@" + parts[0] + ":" + parts[1])
		proxyList = append(proxyList, proxyURL)
	}

	proxyCount := len(proxyList)

	// log.Println(proxyList[0])
	return proxyList, proxyCount, err

}

func GetRandProxy(proxies *[]string) string {

	count := len(*proxies)
	rand.Seed(time.Now().UnixNano())
	x := count
	randInt := rand.Intn(x)
	return (*proxies)[randInt]
}

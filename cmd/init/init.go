package init

import (
	"FurAIOIgnited/cmd/taskengine"
	"FurAIOIgnited/internal/profiles"
	proxy "FurAIOIgnited/internal/proxies"
	"FurAIOIgnited/internal/server"
	"FurAIOIgnited/sites/pa"
	"FurAIOIgnited/sites/passkey"
	"FurAIOIgnited/util"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/sirupsen/logrus"
)

// var addr = flag.String("addr", "localhost:8080", "http service address")

var (
	log         *logrus.Logger
	messageChan = make(chan []byte)
)

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
}

func FlintAndSteel() {
	var site string

	MenuPrompt := &survey.Select{
		Message: "Choose a Site: ",
		Options: []string{"Passkey", "PA Liquor", "PA Liquor (sitewide)"},
	}

	survey.AskOne(MenuPrompt, &site)

	var profileNames []string
	var selectedProfiles []int
	var validProfiles []profiles.Profile

	proxyList, proxyCount, err := proxy.GetProxies()
	if err != nil {
		fmt.Println("Error Reading Proxies")
	}
	if len(proxyList) == proxyCount {
		fmt.Println("All proxies read successfully!")
	}

	profileList := profiles.ReadUserProfiles()
	for i := range profileList {
		if profileList[i].Site == site {
			profileNames = append(profileNames, (profileList)[i].ProfileName)
			validProfiles = append(validProfiles, profileList[i])
		}
	}

	prompt := &survey.MultiSelect{
		Message: "Which Profiles Would You like to use?",
		Options: profileNames,
	}

	survey.AskOne(prompt, &selectedProfiles)

	fmt.Println("Selected profile")

	// Menu
	// TODO: Add some kind of interface that can actually let different site tasks run

	if site == "Passkey" {

		var BookingURL string
		var flowType string

		var queues []passkey.PasskeyTask

		fmt.Print("Enter Booking URL: ")
		fmt.Scanln(&BookingURL)

		// fmt.Print("Enter Task Count: ")
		// fmt.Scanln(&TaskCount)

		// The idea is that "Manual" mode is gonna force the user to select the dates during checkout
		// A "Fallback" Mode or something is going to automatically select relivant dates, but I'm not sure how to calculate that
		// and then just a "monitor" mode.

		switch util.URLType(BookingURL) {
		case "book.passkey.com":
			flowType = "Booking"
		case "queue.passkey.com":
			flowType = "Queue"
		default:
			fmt.Println("Invalid URL!")

			fmt.Println(util.URLType(BookingURL))
		}

		for i := range selectedProfiles {

			t := taskengine.Task{
				Profile: &profileList[selectedProfiles[i]],
				Stage:   passkey.Start,
				Mode:    "Auto",
				Log:     log,
			}
			p := passkey.PasskeyTask{
				GivenURL: BookingURL,
				Task:     t,
				Flow:     flowType,
			}

			queues = append(queues, p)

		}

		// t := taskengine.Task{
		// 	Profile: &profileList[selectedProfile],
		// 	Stage:   passkey.Start,
		// 	Mode:    "Auto",
		// }
		// p := passkey.PasskeyTask{
		// 	GivenURL: BookingURL,
		// 	Task:     t,
		// 	Flow:     flowType,
		// }

		// p.Ignite()

		if flowType == "Queue" {
			fmt.Println("Queue Detected!")

			// var wg sync.WaitGroup

			// // count := 3

			// p := passkey.PasskeyTask{
			// 	GivenURL: BookingURL,
			// 	Task:     t,
			// 	Flow:     flowType,
			// }

			// for i := 0; i < 1; i++ {
			// 	queues = append(queues, p)
			// }

			// wg.Add(len(queues))

			// for i := 0; i < len(queues); i++ {
			// 	go func(i int) {
			// 		queues[i].Ignite()
			// 	}(i)
			// }

			// wg.Wait()
			// p.Ignite()

		} else if flowType == "Booking" {
			// fmt.Println("Starting Booking...")
			// p := passkey.PasskeyTask{
			// 	GivenURL: BookingURL,
			// 	Task:     t,
			// 	Flow:     flowType,
			// }

			// p.Ignite()

			var wg sync.WaitGroup
			wg.Add(len(queues))

			for i := 0; i < len(queues); i++ {
				go func(i int) {
					queues[i].Initalize()
					// queues[i].SetProxy(&proxyList)
					queues[i].Ignite()
				}(i)
			}

			wg.Wait()

		}

	} else if site == "PA Liquor" {

		//Connect to websocket stuff
		var mode string
		var serverside []string
		var serverChoice bool
		MenuPrompt := &survey.Select{
			Message: "Choose a Site: ",
			Options: []string{"Normal", "Sku Search", "Preload"},
		}

		ServerPrompt := &survey.MultiSelect{
			Message: "Monitoring Method?",
			Options: []string{"Serverless"},
		}

		survey.AskOne(MenuPrompt, &mode)
		survey.AskOne(ServerPrompt, &serverside)

		fmt.Printf("serverside: %v\n", serverside)

		var queues []pa.PaTask

		if len(serverside) == 0 {
			serverChoice = false
		} else {
			serverChoice = true
		}

		if serverChoice {
			var skuData server.SkuData

			// go server.ListenToSocket(conn, messageChan)
			go server.ListenToSocketWithReconnect("wss://websocket.lenixsavesthe.world/skus", messageChan)

			for {

				message := <-messageChan
				fmt.Println(message)
				err := json.Unmarshal(message, &skuData)
				if err != nil {
					log.Fatal("error unmarshalling socket json")
				}

				if len(skuData) == 0 {
					log.Println("no products in stock!")
					continue
				}

				if len(skuData) > 0 {
					for i := range skuData {
						for j := range selectedProfiles {
							fmt.Println(validProfiles[selectedProfiles[j]].ProfileName, skuData[i].ProdName, skuData[i].Sku, skuData[i].ProtectedStock)
							t := taskengine.Task{
								Log: log,
								// Websocket: conn,
								Profile: &validProfiles[selectedProfiles[j]],
								Stage:   pa.Start,
								Mode:    mode,
							}
							p := pa.PaTask{
								Task:        t,
								Method:      "Shipping",
								UseAccount:  false,
								Serverside:  serverChoice,
								ProductID:   skuData[i].Sku,
								ProductName: skuData[i].ProdName,
								ProductVol:  skuData[i].ProdVolume,
								QuantityStr: skuData[i].MaxQuantity,
							}
							queues = append(queues, p)
						}
					}
					break
				}

			}
		} else {
			for j := range selectedProfiles {
				fmt.Println(validProfiles[selectedProfiles[j]].ProfileName)
				t := taskengine.Task{
					Log:     log,
					Profile: &validProfiles[selectedProfiles[j]],
					Stage:   pa.Start,
					Mode:    mode,
				}
				p := pa.PaTask{
					Task:       t,
					Method:     "Shipping",
					UseAccount: false,
					Serverside: serverChoice,
				}
				queues = append(queues, p)
			}

		}

		pool := rod.NewBrowserPool(len(queues))

		ul := launcher.NewUserMode().
			Leakless(false).
			UserDataDir("tmp/t").
			Set("disable-default-apps").
			Set("no-first-run").
			Headless(false).
			MustLaunch()

		create := func() *rod.Browser {
			browser := rod.New().ControlURL(ul).MustConnect().MustIncognito().NoDefaultDevice()
			return browser
		}

		var wg sync.WaitGroup
		wg.Add(len(queues))

		for i := 0; i < len(queues); i++ {
			go func(i int) {
				defer wg.Done()
				fmt.Println("starting:", queues[i].ProductID)

				bPool := pool.Get(create)
				queues[i].Initalize()
				queues[i].Ignite(bPool)
			}(i)
		}

		wg.Wait()

	} else if site == "PA Liquor (sitewide)" {

		pa.SiteMonitor()

	}
}

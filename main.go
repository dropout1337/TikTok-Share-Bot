package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/admin100/util/console"
	"github.com/corpix/uarand"
	"github.com/valyala/fasthttp"
)

var (
	client fasthttp.Client

	threads int
	item_id string

	shared int
	failed int

	app_names = []string{"tiktok_web", "musically_go"}
	domains   = []string{"api19.tiktokv.com", "api.toutiao50.com", "api19.toutiao50.com", "api19-core-c-alisg.tiktokv.com"}
	channels  = []string{"tiktok_web", "googleplay", "App%20Store"}
	platforms = []string{"android", "windows", "iphone", "web"}
	devices   = []string{"iPhone1,1", "iPhone1,2", "iPhone2,1", "iPhone3,1", "iPhone3,2", "iPhone3,3", "iPhone4,1", "iPhone5,1", "iPhone5,2", "iPhone5,3", "iPhone5,4", "iPhone6,1", "iPhone6,2", "iPhone7,1", "iPhone7,2", "iPhone8,1", "iPhone8,2", "iPhone8,4", "iPhone9,1", "iPhone9,2", "iPhone9,3", "iPhone9,4", "iPhone10,1", "iPhone10,2", "iPhone10,3", "iPhone10,4", "iPhone10,5", "iPhone10,6", "iPhone11,2", "iPhone11,4", "iPhone11,6", "iPhone11,8", "iPhone12,1", "iPhone12,3", "iPhone12,5", "iPhone12,8", "iPhone13,1", "iPhone13,2", "iPhone13,3", "iPhone13,4", "iPhone14,2", "iPhone14,3", "iPhone14,4", "iPhone14,5", "iPod1,1", "iPod2,1", "iPod3,1", "iPod4,1", "iPod5,1", "iPod7,1", "iPod9,1", "iPad1,1", "iPad1,2", "iPad2,1", "iPad2,2", "iPad2,3", "iPad2,4", "iPad3,1", "iPad3,2", "iPad3,3", "iPad2,5", "iPad2,6", "iPad2,7", "iPad3,4", "iPad3,5", "iPad3,6", "iPad4,1", "iPad4,2", "iPad4,3", "iPad4,4", "iPad4,5", "iPad4,6", "iPad4,7", "iPad4,8", "iPad4,9", "iPad5,1", "iPad5,2", "iPad5,3", "iPad5,4", "iPad6,3", "iPad6,4", "iPad6,7", "iPad6,8", "iPad6,11", "iPad6,12", "iPad7,1", "iPad7,2", "iPad7,3", "iPad7,4", "iPad7,5", "iPad7,6", "iPad7,11", "iPad7,12", "iPad8,1", "iPad8,2", "iPad8,3", "iPad8,4", "iPad8,5", "iPad8,6", "iPad8,7", "iPad8,8", "iPad8,9", "iPad8,10", "iPad8,11", "iPad8,12", "iPad11,1", "iPad11,2", "iPad11,3", "iPad11,4", "iPad11,6", "iPad11,7", "iPad12,1", "iPad12,2", "iPad14,1", "iPad14,2", "iPad13,1", "iPad13,2", "iPad13,4", "iPad13,5", "iPad13,6", "iPad13,7", "iPad13,8", "iPad13,9", "iPad13,10", "iPad13,11", "Watch1,1", "Watch1,2", "Watch2,6", "Watch2,7", "Watch2,3", "Watch2,4", "Watch3,1", "Watch3,2", "Watch3,3", "Watch3,4", "Watch4,1", "Watch4,2", "Watch4,3", "Watch4,4", "Watch5,1", "Watch5,2", "Watch5,3", "Watch5,4", "Watch5,9", "Watch5,10", "Watch5,11", "Watch5,12", "Watch6,1", "Watch6,2", "Watch6,3", "Watch6,4", "Watch6,6", "Watch6,7", "Watch6,8", "Watch6,9", "SM-G9900", "sm-g950f", "SM-A136U1", "SM-M225FV", "SM-E426B", "SM-M526BR", "SM-M326B", "SM-A528B", "SM-F711B", "SM-F926B", "SM-A037G", "SM-A225F", "SM-M325FV", "SM-A226B", "SM-M426B", "SM-A525F"}
)

func main() {
	clear()
	rand.Seed(time.Now().Unix())

	console.SetConsoleTitle("[TikTok Shares] youtube.com/dropoutuwu")
	wg := sync.WaitGroup{}

	fmt.Printf("[\x1b[38;5;63m%s\x1b[0m] Item ID\x1b[38;5;63m>\x1b[0m ", time.Now().Format("15:04:05"))
	fmt.Scanln(&item_id)

	fmt.Printf("[\x1b[38;5;63m%s\x1b[0m] Threads\x1b[38;5;63m>\x1b[0m ", time.Now().Format("15:04:05"))
	fmt.Scanln(&threads)

	fmt.Println("\x1b[?25l")

	go background()
	for i := 0; i < threads; i++ {
		wg.Add(1)
		fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] Successfully started \x1b[38;5;63m%.0f\x1b[0m threads.", time.Now().Format("15:04:05"), float64(i+1))
		go share()
	}

	fmt.Printf("\n[\x1b[38;5;63m%s\x1b[0m] Press \x1b[38;5;63mCTRL+C\x1b[0m at anytime to stop.", time.Now().Format("15:04:05"))

	wg.Wait()
}

func background() {
	for {
		if shared%100 == 0 {
			console.SetConsoleTitle(fmt.Sprintf("[TikTok Shares] youtube.com/dropoutuwu - Sent %d/%d", shared, (failed + shared)))
		}
		time.Sleep(500)
	}
}

func random_device_id() string {
	device := ""
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 19; i++ {
		num := rand.Intn(9-1) + 9
		device += strconv.Itoa(num)
	}

	return device
}

func generate_data() string {
	domain := domains[rand.Intn(len(domains))]
	platform := platforms[rand.Intn(len(platforms))]
	channel := channels[rand.Intn(len(channels))]
	device := devices[rand.Intn(len(devices))]
	app_name := app_names[rand.Intn(len(app_names))]
	version := strconv.Itoa(rand.Intn(12-1) + 12)
	device_id := random_device_id()

	url := fmt.Sprintf("https://%s/aweme/v1/aweme/stats/?channel=%s&device_type=%s&device_id=%s&os_version=%s&version_code=220400&app_name=%s&device_platform=%s&aid=1988", domain, channel, device, device_id, version, app_name, platform)
	return url
}

func share() {
	for {
		url := generate_data()

		req := fasthttp.AcquireRequest()
		res := fasthttp.AcquireResponse()

		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(res)

		req.Header.SetMethod("POST")
		req.SetRequestURI(url)
		req.SetBody([]byte(fmt.Sprintf("item_id=%v&share_delta=1", item_id)))

		req.Header.Set("User-Agent", uarand.GetRandom())
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

		if err := client.Do(req, res); err != nil {
			failed++
			return
		} else {
			shared++
		}
	}
}

func clear() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

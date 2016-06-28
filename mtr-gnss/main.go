package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ozym/mtr"
	"github.com/taskcluster/httpbackoff"
)

type MTR struct {
	client  *http.Client
	scheme  string
	service string
	user    string
	key     string
}

type stnMtr struct {
	rtime   string
	mean    int64
	min     int64
	max     int64
	epochs  int64
	dnsName string
}

var m map[string]stnMtr

func (m MTR) Put(path string) error {

	u := url.URL{
		Scheme: func() string {
			if m.scheme != "" {
				return m.scheme
			}
			return "https"
		}(),
		Host: m.service,
		Path: path,
	}

	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(func() string {
		if m.user != "" {
			return m.user
		}
		return "mtrproducer"
	}(), m.key)
	req.Header.Add("Accept", "*/*")

	res, _, err := httpbackoff.ClientDo(m.client, req)
	if res != nil && res.Body != nil {
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}

	return err
}

func gettags(fp string) ([]string, []string) {
	df, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer df.Close()
	scanner := bufio.NewScanner(df)
	var stn []string
	var deviceID []string
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), ",")
		if strings.Contains(l[4], "gnss.rt") == true {
			stn = append(stn, l[3])
			deviceID = append(deviceID, l[0])
		}
	}
	return stn, deviceID
}

func main() {
	var service string
	var domain string
	var user string
	var key string
	var logpath string
	var tagfile string
	var test bool
	flag.StringVar(&service, "mtr", "mtr-api.geonet.org.nz", "mtr api end-point")
	flag.StringVar(&domain, "domain", "wan.geonet.org.nz", "device domain")
	flag.StringVar(&user, "user", "mtrproducer", "mtr user")
	flag.StringVar(&key, "key", "", "mtr user key")
	flag.StringVar(&logpath, "logpath", "/home/davec/bnc/", "Location of the BNC Log")
	flag.StringVar(&tagfile, "tagfile", "/home/davec/git/network/data/devices.csv", "Location of the Device file containing the tags")
	flag.BoolVar(&test, "test", true, "Print the ouput only")
	flag.Parse()

	// manage connections to the mtr service.
	handler := mtr.NewMTR(service, domain, user, key)

	starttime := time.Now().UTC()
	stn, deviceID := gettags(tagfile)

	m = make(map[string]stnMtr)

	t := "bnclog_"
	t = t + time.Now().UTC().Format("060102")
	file, err := os.Open(logpath + t)
	if err != nil {
		log.Println("BNC Log File not found")
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var Loglines []string
	for scanner.Scan() {
		Loglines = append(Loglines, scanner.Text())
	}
	for s, stn := range stn {
		for _, line := range Loglines {
			if strings.Contains(line, stn) == true {
				var stnlogs []string
				stnlogs = append(stnlogs, line)
				last := strings.Split(stnlogs[len(stnlogs)-1], " ")
				if strings.Contains(last[3], "Mean") == true {
					var mean int64 // 5
					var complete int64
					var min int64 // 8
					var max int64 // 10
					if y, err := strconv.ParseFloat(last[5], 32); err == nil {
						y = y * 1000
						mean = int64(y)
					}
					if y, err := strconv.ParseFloat(strings.TrimRight(last[8], ","), 32); err == nil {
						y = y * 1000
						min = int64(y)
					}
					if y, err := strconv.ParseFloat(strings.TrimRight(last[10], ","), 32); err == nil {
						y = y * 1000
						max = int64(y)
					}
					if y, err := strconv.ParseFloat(last[13], 32); err == nil {
						y = y / 300 * 100
						complete = int64(y)
					}
					m[stn] = stnMtr{last[1], mean, min, max, complete, deviceID[s]}
				} else {
					if strings.Contains(last[3], "Data") == true {
						m[stn] = stnMtr{last[1], 0, 0, 0, 0, deviceID[s]}
					}
				}
			}
		}
		if m[stn].dnsName != "" {
			fmt.Println(stn)
			fmt.Println(m[stn])
			dat := []mtr.Latency{
				mtr.Latency{
					Site:   stn,
					Metric: "latency.gnss.1hz",
					Mean:   m[stn].mean,
					Min:    m[stn].min,
					Max:    m[stn].max,
					TS:     time.Now(),
				}, /*
					mtr.Data{
						Site:   stn,
						Metric: "completeness",
						Value:  m[stn].epochs,
						TS:     time.Now(),
					},*/
			}

			for _, d := range dat {
				fmt.Println(d.Encode())
				if test == false {
					if err := handler.PutMetric(d); err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
	log.Println(time.Now().Sub(starttime))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

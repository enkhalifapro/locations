package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var cityZones = map[string]string{
	"NY": "America/New_York",
}

type IPManager struct {
	httpclient  *http.Client
	providerURL string
}

func NewIPManager() *IPManager {
	return &IPManager{
		httpclient: &http.Client{
			Timeout: 5 * time.Second,
		},
		providerURL: "http://api.hostip.info/get_html.php",
	}
}

type Location struct {
	Country string
	Date    string
	Time    string
}

func (i *IPManager) GetLocation(ip string) (*Location, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?ip=%s", i.providerURL, ip), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't init request")
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := i.httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusBadRequest {
		return nil, errors.New("Bad request")
	}

	// Handle error response
	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "unexpected response")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// response example should be like the following:
	// Country: UNITED STATES (US)
	// City: New York, NY
	// IP: 207.46.197.32
	response := string(body)
	if len(response) == 0 || strings.Contains(response, "Unknown") {
		return nil, errors.New("Unknown location")
	}

	// extract country
	country := strings.TrimSpace(strings.Split(strings.Split(response, "\n")[0], ":")[1])
	// extract city symbol
	citySymbol := strings.TrimSpace(strings.Split(strings.Split(response, "\n")[1], ",")[1])
	zone, ok := cityZones[citySymbol]
	if !ok {
		return nil, fmt.Errorf("time zone of %s is not supported", citySymbol)
	}

	// get zone date and time
	t := time.Now()
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return nil, fmt.Errorf("time zone of %s is not supported", citySymbol)
	}
	t = t.In(loc)

	return &Location{
		Country: country,
		Date:    t.Format("02-01-2006"),
		Time:    t.Format("15:04:05"),
	}, nil
}

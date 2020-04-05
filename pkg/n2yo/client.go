package n2yo

import (
	"fmt"

	"github.com/codingsince1985/geo-golang"
	"github.com/go-resty/resty/v2"
)

const baseURI = "https://www.n2yo.com/rest/v1/satellite"

type N2yoClient struct {
	apiC *resty.Client
	key  string
}
type VisualPass struct {
	Info struct {
		Satid       int
		Satname     string
		Passescount int
	}
	Passes []Pass
}
type Pass struct {
	StartAz        float64
	StartAzCompass string
	StartEl        float64
	StartUTC       int64
	MaxAz          float64
	MaxAzCompass   string
	MaxEl          float64
	MaxUTC         int64
	EndAz          float64
	EndAzCompass   string
	EndEl          float64
	EndUTC         int64
	Mag            float64
	Duration       int
}

func New(apiKey string) *N2yoClient {
	client := resty.New()
	client.SetHostURL(baseURI)
	client.SetHeaders(map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"User-Agent":   "Wyborowy bot",
	})
	client.SetQueryParam("apiKey", apiKey)
	return &N2yoClient{apiC: client, key: apiKey}
}

func (c *N2yoClient) GetSatelitePosition(code string) error {
	r, err := c.apiC.R().Get(fmt.Sprintf("/tle/%s", code))
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func (c *N2yoClient) GetISSPass(loc *geo.Location) (*VisualPass, error) {
	r, err := c.apiC.R().SetResult(&VisualPass{}).Get(fmt.Sprintf("/visualpasses/25544/%f/%f/250/7/180", loc.Lat, loc.Lng))
	if err != nil {
		return nil, err
	}
	return r.Result().(*VisualPass), nil

}

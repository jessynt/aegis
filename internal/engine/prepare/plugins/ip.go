package plugins

import (
	"sync"

	"aegis/pkg/ipip"
	"aegis/resources/ipdb"
)

var (
	loadOnce sync.Once
	ipipCity *ipip.City
)

func ip2Location(ip string) (country, province, city string, err error) {
	loadOnce.Do(func() {
		content := ipdb.MustAsset("ipipfree.ipdb")
		var err error
		ipipCity, err = ipip.NewCityInMemory(content)
		if err != nil {
			panic(err)
		}
	})

	rv, err := ipipCity.Find(ip, "CN")
	if err != nil {
		return "", "", "", err
	}
	return rv[0], rv[1], rv[2], nil
}

func IP2Location(ip string) (map[string]string, error) {
	country, province, city, err := ip2Location(ip)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"Country":  country,
		"Province": province,
		"City":     city,
	}, nil
}

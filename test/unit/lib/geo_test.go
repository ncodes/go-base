package lib

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/franela/goblin"
	"github.com/ncodes/go-base/lib"
	. "github.com/onsi/gomega"
	"github.com/troophq/golang-geo"
)

func TestGetAddComponentByType(t *testing.T) {
	// g := Goblin(t)
	// RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	// g.Describe("GetState", func() {
	// 	g.It("Should return the correct state and country", func() {

	// 		dir, _ := filepath.Abs(filepath.Dir("../../fixtures/json/"))
	// 		dat, _ := ioutil.ReadFile(dir + "/sample_reverse_geo_data.json")

	// 		geocodeResp := geo.GoogleReverseGeocodeResponse{}
	// 		json.Unmarshal(dat, &geocodeResp)
	// 		addressComponent, _ := lib.GetAddressComponentByType("country", &geocodeResp)

	// 		Expect(addressComponent.LongName).To(Equal("Nigeria"))
	// 		Expect(addressComponent.ShortName).To(Equal("NG"))

	// 		addressComponent, _ = lib.GetAddressComponentByType("administrative_area_level_1", &geocodeResp)
	// 		Expect(addressComponent.LongName).To(Equal("Lagos"))

	// 	})
	// })
}

func TestGetStateFromReverseGeoResponse(t *testing.T) {
	// g := Goblin(t)
	// RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	// g.Describe("GetState", func() {
	// 	g.It("Should return the correct state", func() {

	// 		dir, _ := filepath.Abs(filepath.Dir("../../fixtures/json/"))
	// 		dat, _ := ioutil.ReadFile(dir + "/sample_reverse_geo_data.json")

	// 		geocodeResp := geo.GoogleReverseGeocodeResponse{}
	// 		json.Unmarshal(dat, &geocodeResp)
	// 		result := lib.GetStateFromReverseGeoResponse(&geocodeResp)
	// 		Expect(result).To(Equal("Lagos"))
	// 	})
	// })
}

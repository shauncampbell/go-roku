package roku_test

import (
	"encoding/xml"
	"github.com/koron/go-ssdp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shauncampbell/go-roku"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("GoRoku", func() {
	var deviceInfo roku.DeviceInfo

	BeforeEach(func() {
		deviceInfo = roku.DeviceInfo{
			UDN:                         "8000000-0000-1000-8000-d81399f9e18b",
			SerialNumber:                "X00000PGAVCY",
			DeviceID:                    "S036D99GAVCY",
			AdvertisingID:               "428b7ed-4988-5218-a5d0-6977857dcddd",
			VendorName:                  "TCL",
			ModelName:                   "50S425-CA",
			ModelNumber:                 "C105X",
			ModelRegion:                 "CA",
			IsTV:                        true,
			IsStick:                     false,
			ScreenSize:                  50,
			PanelId:                     17,
			TunerType:                   "ATSC",
			SupportsEthernet:            true,
			WifiMacAddress:              "wifi-mac",
			WifiDriver:                  "realtek",
			EthernetMacAddress:          "ethernet-mac",
			NetworkType:                 "wifi",
			NetworkName:                 "Wireless",
			FriendlyDeviceName:          "TCL Roku TV",
			FriendlyModelName:           "TCL Roku TV",
			DefaultDeviceName:           "TCL•Roku TV - X00000PGAVCY",
			UserDeviceName:              "50\" TCL Roku TV",
			UserDeviceLocation:          "Living Room",
			BuildNumber:                 "939.20E04502A",
			SoftwareVersion:             "9.2.0",
			SoftwareBuild:               4502,
			SecureDevice:                true,
			Language:                    "en",
			Country:                     "CA",
			Locale:                      "en_US",
			TimeZoneAuto:                true,
			TimeZone:                    "Canada/Eastern",
			TimeZoneName:                "Canada/Eastern",
			TimeZoneTz:                  "America/Toronto",
			TimeZoneOffset:              -240,
			ClockFormat:                 "12-hour",
			Uptime:                      123,
			PowerMode:                   "PowerOn",
			SupportsSuspend:             true,
			SupportsFindRemote:          true,
			FindRemoteIsPossible:        false,
			SupportsAudioGuide:          false,
			SupportsRva:                 true,
			DeveloperEnabled:            true,
			KeyedDeveloperId:            "",
			SearchEnabled:               true,
			SearchChannelsEnabled:       true,
			VoiceSearchEnabled:          true,
			NotificationsEnabled:        true,
			NotificationsFirstUse:       true,
			SupportsPrivateListening:    true,
			SupportsPrivateListeningDtv: true,
			SupportsWarmStandby:         true,
			HeadphonesConnected:         true,
			ExpertPqEnabled:             "0.9",
			SupportsEcsTextedit:         true,
			SupportsEcsMicrophone:       true,
			SupportsWakeOnWlan:          true,
			HasPlayOnRoku:               true,
			HasMobileScreensaver:        false,
			SupportUrl:                  "tclcanada.com/support",
			GrandcentralVersion:         "2.9.57",
			TrcVersion:                  "3.0",
			TrcChannelVersion:           "2.9.42",
			HasWifiExtender:             false,
			HasWifi5GSupport:            true,
			CanUseWifiExtender:          true,
		}
	})
	Context("Device Information", func() {
		It("should succeed when a valid xml response is returned from roku", func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				bytes, _ := xml.Marshal(deviceInfo)
				w.Write(bytes)
			}))

			devices := roku.ProcessDevices([]ssdp.Service{{Location: server.URL}}, server.Client())
			Expect(len(devices)).To(Equal(1))

			for _, device := range devices {
				device.GetDeviceInformation()
			}
		})
		It("should fail when an invalid xml response is returned from roku", func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte{})
			}))

			devices := roku.ProcessDevices([]ssdp.Service{
				{Location: server.URL},
			}, server.Client())
			Expect(len(devices)).To(Equal(1))

			for _, device := range devices {
				device.GetDeviceInformation()
			}
		})
		It("should fail when an invalid server is returned from ssdp scan", func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte{})
			}))

			devices := roku.ProcessDevices([]ssdp.Service{{Location: ""}}, server.Client())
			Expect(len(devices)).To(Equal(1))

			for _, device := range devices {
				device.GetDeviceInformation()
			}
		})

	})
})

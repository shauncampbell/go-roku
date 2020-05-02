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
	var appList roku.InstalledAppsList
	var activeApp roku.ActiveAppInfo

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
		appList = roku.InstalledAppsList{
			Apps: []roku.AppInfo{
				{ Id: "tvinput.hdmi1", Type: "tvin", Version: "1.0.0", Name: "Cable TV" },
				{ Id: "tvinput.hdmi2", Type: "tvin", Version: "1.0.0", Name: "PlayStation" },
				{ Id: "tvinput.hdmi3", Type: "tvin", Version: "1.0.0", Name: "Apple TV" },
				{ Id: "tvinput.cvbs", Type: "tvin", Version: "1.0.0", Name: "AV" },
				{ Id: "tvinput.dtv", Type: "tvin", Version: "1.0.0", Name: "Antenna TV" },
				{ Id: "12", Subtype: "ndka", Type: "appl", Version: "5.1.81188066", Name: "Netflix" },
				{ Id: "50025", Subtype: "rsga", Type: "appl", Version: "2.1.20200123", Name: "Google Play Movies &amp; TV" },
				{ Id: "229863", Subtype: "rsga", Type: "appl", Version: "2.0.1", Name: "CBS All Access - Canada" },
				{ Id: "27181", Subtype: "rsga", Type: "appl", Version: "4.51.216", Name: "Sky News" },
			},
		}
		activeApp = roku.ActiveAppInfo{
			App: roku.AppInfo {
				Id: "13", Subtype: "ndka", Type: "appl", Version: "11.2.2020032710", Name: "Prime Video",
			},
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

	Context("Installed Apps", func() {
		It("should succeed when a valid xml response is returned from roku", func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				bytes, _ := xml.Marshal(appList)
				w.Write(bytes)
			}))

			devices := roku.ProcessDevices([]ssdp.Service{{Location: server.URL}}, server.Client())
			Expect(len(devices)).To(Equal(1))

			for _, device := range devices {
				appsList, err := device.GetInstalledApps()
				Expect(err).To(BeNil())
				Expect(len(appsList.Apps)).To(Equal(9))

				i := 0
				Expect(appsList.Apps[i].Id).To(Equal("tvinput.hdmi1"))
				Expect(appsList.Apps[i].Type).To(Equal("tvin"))
				Expect(appsList.Apps[i].Version).To(Equal("1.0.0"))
				Expect(appsList.Apps[i].Name).To(Equal("Cable TV"))

				i++
				Expect(appsList.Apps[i].Id).To(Equal("tvinput.hdmi2"))
				Expect(appsList.Apps[i].Type).To(Equal("tvin"))
				Expect(appsList.Apps[i].Version).To(Equal("1.0.0"))
				Expect(appsList.Apps[i].Name).To(Equal("PlayStation"))

				i++
				Expect(appsList.Apps[i].Id).To(Equal("tvinput.hdmi3"))
				Expect(appsList.Apps[i].Type).To(Equal("tvin"))
				Expect(appsList.Apps[i].Version).To(Equal("1.0.0"))
				Expect(appsList.Apps[i].Name).To(Equal("Apple TV"))

				i++
				Expect(appsList.Apps[i].Id).To(Equal("tvinput.cvbs"))
				Expect(appsList.Apps[i].Type).To(Equal("tvin"))
				Expect(appsList.Apps[i].Version).To(Equal("1.0.0"))
				Expect(appsList.Apps[i].Name).To(Equal("AV"))

				i++
				Expect(appsList.Apps[i].Id).To(Equal("tvinput.dtv"))
				Expect(appsList.Apps[i].Type).To(Equal("tvin"))
				Expect(appsList.Apps[i].Version).To(Equal("1.0.0"))
				Expect(appsList.Apps[i].Name).To(Equal("Antenna TV"))

				i++
				Expect(appsList.Apps[i].Id).To(Equal("12"))
				Expect(appsList.Apps[i].Subtype).To(Equal("ndka"))
				Expect(appsList.Apps[i].Type).To(Equal("appl"))
				Expect(appsList.Apps[i].Version).To(Equal("5.1.81188066"))
				Expect(appsList.Apps[i].Name).To(Equal("Netflix"))

				i++
				Expect(appsList.Apps[i].Id).To(Equal("50025"))
				Expect(appsList.Apps[i].Subtype).To(Equal("rsga"))
				Expect(appsList.Apps[i].Type).To(Equal("appl"))
				Expect(appsList.Apps[i].Version).To(Equal("2.1.20200123"))
				Expect(appsList.Apps[i].Name).To(Equal("Google Play Movies &amp; TV"))

				i++
				Expect(appsList.Apps[i].Id).To(Equal("229863"))
				Expect(appsList.Apps[i].Subtype).To(Equal("rsga"))
				Expect(appsList.Apps[i].Type).To(Equal("appl"))
				Expect(appsList.Apps[i].Version).To(Equal("2.0.1"))
				Expect(appsList.Apps[i].Name).To(Equal("CBS All Access - Canada"))

				i++
				Expect(appsList.Apps[i].Id).To(Equal("27181"))
				Expect(appsList.Apps[i].Subtype).To(Equal("rsga"))
				Expect(appsList.Apps[i].Type).To(Equal("appl"))
				Expect(appsList.Apps[i].Version).To(Equal("4.51.216"))
				Expect(appsList.Apps[i].Name).To(Equal("Sky News"))
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
				appsList, err := device.GetInstalledApps()
				Expect(err).ToNot(BeNil())
				Expect(appsList).To(BeNil())
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
				appsList, err := device.GetInstalledApps()
				Expect(err).ToNot(BeNil())
				Expect(appsList).To(BeNil())
			}
		})
	})

	Context("Current App Information", func() {
		It("should succeed when a valid xml response is returned from roku", func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				bytes, _ := xml.Marshal(activeApp)
				w.Write(bytes)
			}))

			devices := roku.ProcessDevices([]ssdp.Service{{Location: server.URL}}, server.Client())
			Expect(len(devices)).To(Equal(1))

			for _, device := range devices {
				app, err := device.GetActiveApp()
				Expect(err).To(BeNil())
				Expect(app).ToNot(BeNil())
				Expect(app.App.Id).To(Equal("13"))
				Expect(app.App.Subtype).To(Equal("ndka"))
				Expect(app.App.Type).To(Equal("appl"))
				Expect(app.App.Name).To(Equal("Prime Video"))
				Expect(app.App.Version).To(Equal("11.2.2020032710"))
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
				app, err := device.GetActiveApp()
				Expect(err).ToNot(BeNil())
				Expect(app).To(BeNil())
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
				app, err := device.GetActiveApp()
				Expect(err).ToNot(BeNil())
				Expect(app).To(BeNil())
			}
		})
	})
})

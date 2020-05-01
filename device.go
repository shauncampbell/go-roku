package roku

import (
	"encoding/xml"
	"net/http"
)

//	DeviceInformationProvider contains the controls to query a roku device
type DeviceInformationProvider interface {
	GetDeviceInformation() (*DeviceInfo, error)
}

//	Device is a structure which contains information about a roku device.
type Device struct {
	URL    string
	client *http.Client
	DeviceInformationProvider
}

//	DeviceInfo is a structure for holding information retrieved via roku ecp api.
type DeviceInfo struct {
	UDN                         string `xml:"udn"`
	SerialNumber                string `xml:"serial-number"`
	DeviceID                    string `xml:"device-id"`
	AdvertisingID               string `xml:"advertising-id"`
	VendorName                  string `xml:"vendor-name"`
	ModelName                   string `xml:"model-name"`
	ModelNumber                 string `xml:"model-number"`
	ModelRegion                 string `xml:"model-region"`
	IsTV                        bool   `xml:"is-tv"`
	IsStick                     bool   `xml:"is-stick"`
	ScreenSize                  int    `xml:"screen-size"`
	PanelId                     int    `xml:"panel-id"`
	TunerType                   string `xml:"tuner-type"`
	SupportsEthernet            bool   `xml:"supports-ethernet"`
	WifiMacAddress              string `xml:"wifi-mac"`
	WifiDriver                  string `xml:"wifi-driver"`
	EthernetMacAddress          string `xml:"ethernet-mac"`
	NetworkType                 string `xml:"network-type"`
	NetworkName                 string `xml:"network-name"`
	FriendlyDeviceName          string `xml:"friendly-device-name"`
	FriendlyModelName           string `xml:"friendly-model-name"`
	DefaultDeviceName           string `xml:"default-device-name"`
	UserDeviceName              string `xml:"user-device-name"`
	UserDeviceLocation          string `xml:"user-device-location"`
	BuildNumber                 string `xml:"build-number"`
	SoftwareVersion             string `xml:"software-version"`
	SoftwareBuild               int    `xml:"software-build"`
	SecureDevice                bool   `xml:"secure-device"`
	Language                    string `xml:"language"`
	Country                     string `xml:"country"`
	Locale                      string `xml:"locale"`
	TimeZoneAuto                bool   `xml:"time-zone-auto"`
	TimeZone                    string `xml:"time-zone"`
	TimeZoneName                string `xml:"time-zone-name"`
	TimeZoneTz                  string `xml:"time-zone-tz"`
	TimeZoneOffset              int    `xml:"time-zone-offset"`
	ClockFormat                 string `xml:"clock-format"`
	Uptime                      int    `xml:"uptime"`
	PowerMode                   string `xml:"power-mode"`
	SupportsSuspend             bool   `xml:"supports-suspend"`
	SupportsFindRemote          bool   `xml:"supports-find-remote"`
	FindRemoteIsPossible        bool   `xml:"find-remote-is-possible"`
	SupportsAudioGuide          bool   `xml:"supports-audio-guide"`
	SupportsRva                 bool   `xml:"supports-rva"`
	DeveloperEnabled            bool   `xml:"developer-enabled"`
	KeyedDeveloperId            string `xml:"keyed-developer-id"`
	SearchEnabled               bool   `xml:"search-enabled"`
	SearchChannelsEnabled       bool   `xml:"search-channels-enabled"`
	VoiceSearchEnabled          bool   `xml:"voice-search-enabled"`
	NotificationsEnabled        bool   `xml:"notifications-enabled"`
	NotificationsFirstUse       bool   `xml:"notifications-first-use"`
	SupportsPrivateListening    bool   `xml:"supports-private-listening"`
	SupportsPrivateListeningDtv bool   `xml:"supports-private-listening-dtv"`
	SupportsWarmStandby         bool   `xml:"supports-warm-standby"`
	HeadphonesConnected         bool   `xml:"headphones-connected"`
	ExpertPqEnabled             string `xml:"expert-pq-enabled"`
	SupportsEcsTextedit         bool   `xml:"supports-ecs-textedit"`
	SupportsEcsMicrophone       bool   `xml:"supports-ecs-microphone"`
	SupportsWakeOnWlan          bool   `xml:"supports-wake-on-wlan"`
	HasPlayOnRoku               bool   `xml:"has-play-on-roku"`
	HasMobileScreensaver        bool   `xml:"has-mobile-screensaver"`
	SupportUrl                  string `xml:"support-url"`
	GrandcentralVersion         string `xml:"grandcentral-version"`
	TrcVersion                  string `xml:"trc-version"`
	TrcChannelVersion           string `xml:"trc-channel-version"`
	HasWifiExtender             bool   `xml:"has-wifi-extender"`
	HasWifi5GSupport            bool   `xml:"has-wifi-5G-support"`
	CanUseWifiExtender          bool   `xml:"can-use-wifi-extender"`
}

// GetDeviceInformation
func (d *Device) GetDeviceInformation() (*DeviceInfo, error) {
	resp, err := d.client.Get(d.URL + "/query/device-info")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var info DeviceInfo
	if err := xml.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}

	return &info, nil
}

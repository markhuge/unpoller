package inputunifi

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/unifi-poller/unifi"
	"github.com/unifi-poller/webserver"
)

/* This code reformats our data to be displayed on the built-in web interface. */

func updateWeb(metrics *Metrics) {
	log.Println("here")
	webserver.UpdateInput(&webserver.Input{
		Sites:   formatSites(metrics.Sites),
		Clients: formatClients(metrics.Clients),
		Devices: formatDevices(metrics.Devices),
	})
}

func formatConfig(config *Config) *Config {
	return &Config{
		Default:     *formatControllers([]*Controller{&config.Default})[0],
		Disable:     config.Disable,
		Dynamic:     config.Dynamic,
		Controllers: formatControllers(config.Controllers),
	}
}

func formatControllers(controllers []*Controller) []*Controller {
	fixed := []*Controller{}
	for _, c := range controllers {
		fixed = append(fixed, &Controller{
			VerifySSL:  c.VerifySSL,
			SaveAnomal: c.SaveAnomal,
			SaveAlarms: c.SaveAlarms,
			SaveEvents: c.SaveEvents,
			SaveIDS:    c.SaveIDS,
			SaveDPI:    c.SaveDPI,
			HashPII:    c.HashPII,
			SaveSites:  c.SaveSites,
			User:       c.User,
			Pass:       strconv.FormatBool(c.Pass != ""),
			URL:        c.URL,
			Sites:      c.Sites,
		})
	}

	return fixed
}

func formatSites(sites []*unifi.Site) (s webserver.Sites) {
	for _, site := range sites {
		s = append(s, &webserver.Site{
			ID:     site.ID,
			Name:   site.Name,
			Desc:   site.Desc,
			Source: site.SourceName,
		})
	}

	return s
}

func formatClients(clients []*unifi.Client) (c webserver.Clients) {
	for _, client := range clients {
		clientType, deviceMAC := "unknown", "unknown"
		if client.ApMac != "" {
			clientType = "wireless"
			deviceMAC = client.ApMac
		} else if client.SwMac != "" {
			clientType = "wired"
			deviceMAC = client.SwMac
		}

		if deviceMAC == "" {
			deviceMAC = client.GwMac
		}

		c = append(c, &webserver.Client{
			Name:      client.Name,
			SiteID:    client.SiteID,
			Source:    client.SourceName,
			MAC:       client.Mac,
			IP:        client.IP,
			Type:      clientType,
			DeviceMAC: deviceMAC,
			Since:     time.Unix(client.FirstSeen, 0),
			Last:      time.Unix(client.LastSeen, 0),
		})
	}

	return c
}

func formatDevices(devices *unifi.Devices) (d webserver.Devices) {
	for _, device := range devices.UAPs {
		d = append(d, &webserver.Device{
			Name:    device.Name,
			SiteID:  device.SiteID,
			Source:  device.SourceName,
			MAC:     device.Mac,
			IP:      device.IP,
			Type:    device.Type,
			Model:   device.Model,
			Version: device.Version,
			Config:  nil,
		})
	}

	for _, device := range devices.UDMs {
		d = append(d, &webserver.Device{
			Name:    device.Name,
			SiteID:  device.SiteID,
			Source:  device.SourceName,
			MAC:     device.Mac,
			IP:      device.IP,
			Type:    device.Type,
			Model:   device.Model,
			Version: device.Version,
			Config:  nil,
		})
	}

	for _, device := range devices.USWs {
		d = append(d, &webserver.Device{
			Name:    device.Name,
			SiteID:  device.SiteID,
			Source:  device.SourceName,
			MAC:     device.Mac,
			IP:      device.IP,
			Type:    device.Type,
			Model:   device.Model,
			Version: device.Version,
			Config:  nil,
		})
	}

	for _, device := range devices.USGs {
		d = append(d, &webserver.Device{
			Name:    device.Name,
			SiteID:  device.SiteID,
			Source:  device.SourceName,
			MAC:     device.Mac,
			IP:      device.IP,
			Type:    device.Type,
			Model:   device.Model,
			Version: device.Version,
			Config:  nil,
		})
	}

	return d
}

// Logf logs a message.
func (u *InputUnifi) Logf(msg string, v ...interface{}) {
	webserver.NewInputEvent(PluginName, PluginName, &webserver.Event{
		Ts:   time.Now(),
		Msg:  fmt.Sprintf(msg, v...),
		Tags: map[string]string{"type": "info"},
	})
	u.Logger.Logf(msg, v...)
}

// LogErrorf logs an error message.
func (u *InputUnifi) LogErrorf(msg string, v ...interface{}) {
	webserver.NewInputEvent(PluginName, PluginName, &webserver.Event{
		Ts:   time.Now(),
		Msg:  fmt.Sprintf(msg, v...),
		Tags: map[string]string{"type": "error"},
	})
	u.Logger.LogErrorf(msg, v...)
}

// LogDebugf logs a debug message.
func (u *InputUnifi) LogDebugf(msg string, v ...interface{}) {
	webserver.NewInputEvent(PluginName, PluginName, &webserver.Event{
		Ts:   time.Now(),
		Msg:  fmt.Sprintf(msg, v...),
		Tags: map[string]string{"type": "debug"},
	})
	u.Logger.LogDebugf(msg, v...)
}

package house

import (
	"context"
	"fmt"
	"github.com/beaujr/nmap_prometheus/notifications"
	"gopkg.in/yaml.v2"
	"log"
	"strconv"
	"strings"
)

type home struct {
	Empty   bool `json:"empty",yaml:"empty"`
	Timeout int  `json:"timeout",yaml:"timeout"`
}

func (s *Server) writeHome(id string, item *home) error {
	d1, err := yaml.Marshal(item)
	if err != nil {
		log.Fatalf(err.Error())
	}

	key := fmt.Sprintf("%s%s", homePrefix, id)
	_, err = s.etcdClient.Put(context.Background(), key, string(d1))
	return err
}

func (s *Server) iotStatusManager() error {
	gHouseEmpty, err := s.readHomesConfig()
	if err != nil {
		return err
	}
	for home, empty := range gHouseEmpty {
		if houseEmpty := s.isHouseEmpty(home); houseEmpty != *empty {
			err = s.toggleHouseStatus(home, houseEmpty)
			if err != nil {
				return err
			}
		}
		if !s.isHouseEmpty(home) {
			keys, err := s.getTcKeys()
			if err != nil {
				return err
			}
			for _, key := range keys {
				if strings.Contains(*key, home) {
					err = s.deleteTcByKey(*key)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (s *Server) toggleHouseStatus(home string, houseEmpty bool) error {
	_, err := s.etcdClient.Put(context.Background(), fmt.Sprintf("%s%s", homePrefix, home), strconv.FormatBool(houseEmpty))
	if err != nil {
		log.Println(err)
		return err
	}
	if *debug {
		log.Printf("House (%s) is Empty(%v)", home, houseEmpty)
	} else {
		err := notifications.SendNotification("House Empty", fmt.Sprintf("No Humans in %s", home), home)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	devices, err := s.readNetworkConfig()
	if err != nil {
		log.Println(err)
		return err
	}
	for _, device := range devices {
		if device.PresenceAware && strings.Compare(home, device.Home) == 0 {
			err = s.createTimedCommand(3600, device.Id.Mac, home, fmt.Sprintf("Turn %s off", device.Name), device.Name)
			if err != nil {
				return err
			}
		}
	}
	return err
}

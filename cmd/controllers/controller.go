package controllers

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/LombardiDaniel/ip-checker/services"
	"github.com/LombardiDaniel/ip-checker/utils"
)

var deviceName string = utils.GetEnvVarDefault("DEVICE_NAME", "raspberry")

type Controller struct {
	ipCheckerService services.IpCheckerService
	notifierService  services.NotifierService
}

func NewController(ipCheckerService services.IpCheckerService, notifierService services.NotifierService) Controller {
	return Controller{
		ipCheckerService: ipCheckerService,
		notifierService:  notifierService,
	}
}

func (c *Controller) Loop() error {
	oldIp, err := c.ipCheckerService.ReadOldIp()
	if err != nil {
		err = errors.Join(err, errors.New("could not read ip from ip file on disk"))
		return err
	}

	for {
		newIp, err := c.ipCheckerService.GetCurrIp()
		if err != nil {
			err = errors.Join(err, errors.New("could not read curr ip"))
			return err
		}

		// fmt.Printf("newIp: %v: %v\n", newIp, len(newIp))
		// fmt.Printf("oldIp: %v: %v\n", oldIp, len(oldIp))

		if newIp != oldIp {
			slog.Info(fmt.Sprintf("ip change: new: '%s'", newIp))
			oldIp = newIp

			err = c.notifierService.SendNotification(newIp)
			if err != nil {
				err = errors.Join(err, errors.New("could not send notification, check notifications env vars"))
				return err
			}

			err = c.ipCheckerService.StoreIp(newIp)
			if err != nil {
				err = errors.Join(err, errors.New("could not write ip to ip file on disk"))
				return err
			}
			oldIp = newIp
			slog.Info("TROCOU")
		}

		time.Sleep(1 * time.Minute)
	}
}

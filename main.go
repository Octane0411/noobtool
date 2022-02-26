package main

import (
	"log"
	"noobtool/cmd"
	"noobtool/global"
	settingpkg "noobtool/pkg/setting"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func setupSetting() error {
	setting, err := settingpkg.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Ipgw", &global.IpgwSetting)
	if err != nil {
		return err
	}
	return nil
}

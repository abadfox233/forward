package main

import (
	"os"

	"github.com/kardianos/service"
)

var logger = service.ConsoleLogger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	Start()
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "sshforward",
		DisplayName: "forward ssh to wsl2",
		Description: "forward ssh to wsl2",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		logger.Error(err)
	}

	if err != nil {
		logger.Error(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			s.Install()
			logger.Info("服务安装成功!")
			s.Start()
			logger.Info("服务启动成功!")

		case "start":
			s.Start()
			logger.Info("服务启动成功!")

		case "stop":
			s.Stop()
			logger.Info("服务关闭成功!")

		case "restart":
			s.Stop()
			logger.Info("服务关闭成功!")
			s.Start()
			logger.Info("服务启动成功!")

		case "remove":
			s.Stop()
			logger.Info("服务关闭成功!")
			s.Uninstall()
			logger.Info("服务卸载成功!")
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

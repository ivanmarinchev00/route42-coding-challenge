package main

import (
	"fmt" 
	"coding-challenge/config"
	"coding-challenge/handler"
	"os"
)

func main() {
	err:= run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
// create and load config, execute handler
func run() error {
	cnfg := config.NewConfig()
	err := config.LoadConfig(cnfg)
	
	if err != nil {
		return err
	}

	handler.Handler(cnfg.SERVICE, cnfg.CTX, cnfg.CANCEL)
	cnfg.SERVICE.Wait()

	return nil
}
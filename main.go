package main 
import (
	"github.com/urfave/cli/v2"
	"fmt"
	"os"
	"log"
)


func main(){
	app := &cli.App{
		Name: "Healthchecker",
		Usage: "A tiny tool that checks whether a website is up and running or not "
		Flags: [] cli.Flag{
			&cli.StringFlag{
				Name:"domain",
				Alises::[]string{"d"},
				Usage:"Domain name to check.",
				Required:true,
			},
			&cli.StringFlag{
				Name:"port",
				Aliases:[]string("p")
				Usage:"Port number to check",
				Requred:false,
			},
		Action: func (c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == ""{
				// the default port 
				port = "80"
			}
			status := Check(c.String("domain"),port)// we send domain and port to check function present in check.go file
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err !=nil {
		log.Fatal(err)
	}

}


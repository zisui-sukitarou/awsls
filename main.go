package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AdRoll/goamz/aws"
	"github.com/AdRoll/goamz/ec2"

	"github.com/urfave/cli/v2"
)

func showEC2(c *cli.Context) error {
	auth, err := aws.CredentialFileAuth(os.Getenv("HOME") + "/.aws/credentials", c.String("profile"), 1)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	e := ec2.New(auth, aws.APNortheast)

	response, err := e.DescribeInstances(nil, nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, reservations := range response.Reservations {
		for _, instance := range reservations.Instances {
			fmt.Println(instance.InstanceId, instance.DNSName, instance.IPAddress)
		}
	}

	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "awsls"
	app.Version = "0.0.0"
	app.Commands = []*cli.Command{
		{
			Name: "ec2",
			Usage: "show ec2 instances",
			Action: showEC2,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "profile, p",
					Value: "",
					Usage: "[profile]",
				},
			},
		},
	}
	app.Run(os.Args)
}
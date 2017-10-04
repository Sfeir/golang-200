package main

import (
	"encoding/base64"
	"fmt"
	"github.com/Sfeir/golang-200/dao"
	"github.com/Sfeir/golang-200/utils"
	"github.com/Sfeir/golang-200/web"
	logger "github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
	"time"
)

var (
	// Version is the version of the software
	Version string
	// BuildStmp is the build date
	BuildStmp string
	// GitHash is the git build hash
	GitHash string

	port               = 8020
	logLevel           = "warning"
	db                 = "mongodb://mongo/todos"
	logFormat          = utils.TextFormatter
	statisticsDuration = 20 * time.Second

	header, _ = base64.StdEncoding.DecodeString(
		"DQoNCiAgOzs7Ozs7Ozs7ICAgICAgICAgICAgICAgICAgICAgICA7Ozs7Ozs7OzsNCiAgOzs7Ozs7Ozs7ICAgICAgICAgICAgICAgIC" +
			"AgICAgICA7Ozs7Ozs7OzsNCiAgOzs7OyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDs7OzsNCiAgOzs7OyAgICAg" +
			"ICAgICAgICAgICAgICAgICAgICAgICAgICAgIDs7OzsNCiAgOzs7OyAgICAgICA7Ozs7Ozs7Ozs7Ozs7Ozs7OzsgICAgICAgID" +
			"s7OzsNCiAgOzs7OyAgICAgICA7Ozs7Ozs7Ozs7Ozs7Ozs7OzsgICAgICAgIDs7OzsNCiAgOzs7OyAgICAgICAgICAgICAgICAg" +
			"ICAgICAgICAgICAgICAgIDs7OzsNCiAgOzs7OyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDs7OzsNCiAgOzs7Oy" +
			"AgICAgICA7Ozs7Ozs7Ozs7Ozs7Ozs7OzsgICAgICAgIDs7OzsNCiAgOzs7OyAgICAgICA7Ozs7Ozs7Ozs7Ozs7Ozs7OzsgICAg" +
			"ICAgIDs7OzsNCiAgOzs7OyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDs7OzsNCiAgOzs7OyAgICAgICAgICAgIC" +
			"AgICAgICAgICAgICAgICAgICAgIDs7OzsNCiAgOzs7OyAgICAgICA7Ozs7Ozs7Ozs7Ozs7Ozs7OzsgICAgICAgIDs7OzsNCiAg" +
			"Ozs7OyAgICAgICA7Ozs7Ozs7Ozs7Ozs7Ozs7OzsgICAgICAgIDs7OzsNCiAgOzs7OyAgICAgICAgICAgICAgICAgICAgICAgIC" +
			"AgICAgICAgIDs7OzsNCiAgOzs7OyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDs7OzsNCiAgOzs7Ozs7Ozs7ICAg" +
			"ICAgICAgICAgICAgICAgICAgICA7Ozs7Ozs7OzsNCiAgOzs7Ozs7Ozs7ICAgICAgICAgICAgICAgICAgICAgICA7Ozs7Ozs7OzsNCg0K")
)

func main() {
	// new app
	app := cli.NewApp()
	app.Name = utils.AppName
	app.Usage = "todolist service launcher"

	timeStmp, err := strconv.Atoi(BuildStmp)
	if err != nil {
		timeStmp = 0
	}
	app.Version = Version + ", build on " + time.Unix(int64(timeStmp), 0).String() + ", git hash " + GitHash
	app.Authors = []cli.Author{{Name: "sfr"}}
	app.Copyright = "Sfeir " + strconv.Itoa(time.Now().Year())

	// command line flags
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Value:       port,
			Name:        "port, p",
			Usage:       "Set the listening port of the webserver",
			Destination: &port,
		},
		cli.StringFlag{
			Value:       db,
			Name:        "db, d",
			Usage:       "Set the mongo database connection string",
			Destination: &db,
		},
		cli.StringFlag{
			Value:       logLevel,
			Name:        "logl, l",
			Usage:       "Set the output log level (debug, info, warning, error)",
			Destination: &logLevel,
		},
		cli.StringFlag{
			Value:       logFormat,
			Name:        "logf, f",
			Usage:       "Set the log formatter (logstash or text)",
			Destination: &logFormat,
		},
		cli.DurationFlag{
			Value:       statisticsDuration,
			Name:        "statd, s",
			Usage:       "Set the statistics accumulation duration (ex : 1h, 2h30m, 30s, 300ms)",
			Destination: &statisticsDuration,
		},
	}

	// main action
	// sub action are also possible
	app.Action = func(c *cli.Context) error {
		// print header
		fmt.Println(string(header))

		// set timezone as UTC for bson/json time marshalling
		time.Local = time.UTC

		fmt.Print("* --------------------------------------------------- *\n")
		fmt.Printf("|   port                    : %d\n", port)
		fmt.Printf("|   db                      : %s\n", db)
		fmt.Printf("|   logger level            : %s\n", logLevel)
		fmt.Printf("|   logger format           : %s\n", logFormat)
		fmt.Printf("|   statistic duration(s)   : %0.f\n", statisticsDuration.Seconds())
		fmt.Print("* --------------------------------------------------- *\n")

		// init log options from command line params
		err := utils.InitLog(logLevel, logFormat)
		if err != nil {
			logger.Warn("error setting log level, using debug as default")
		}

		// build the web server
		webServer, err := web.BuildWebServer(db, dao.DAOMongo, statisticsDuration)

		if err != nil {
			return err
		}

		// serve
		webServer.Run(":" + strconv.Itoa(port))

		return nil
	}

	// run the app
	err = app.Run(os.Args)
	if err != nil {
		logger.Fatalf("Run error %q\n", err)
	}
}

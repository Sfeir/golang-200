package main

import (
	"encoding/base64"
	"fmt"
	"github.com/Sfeir/golang-200/dao"
	"github.com/Sfeir/golang-200/utils"
	"github.com/Sfeir/golang-200/web"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
	"time"
)

var (
	// TODO watch these variable initialized from the build args in the makefile
	// Version is the version of the software
	Version string
	// BuildStmp is the build date
	BuildStmp string
	// GitHash is the git build hash
	GitHash string

	port               = 8020
	logLevel           = "warning"
	db                 = ""
	dbType             = dao.DAOMockStr
	migrationPath      = "migration"
	logFormat          = utils.TextFormatter
	statisticsDuration = 20 * time.Second

	header, _ = base64.StdEncoding.DecodeString(
		"ICAgICAgICAgLF8tLS1+fn5+fi0tLS0uXyAgICAgICAgIAogIF8sLF8sKl5fX19fICAgICAgX19fX19gYCpnKlwiKi" +
			"wgCiAvIF9fLyAvJyAgICAgXi4gIC8gICAgICBcIF5AcSAgIGYgClsgIEBmIHwgQCkpICAgIHwgIHwgQCkpICAgbCA" +
			"gMCBfLyAgCiBcYC8gICBcfl9fX18gLyBfXyBcX19fX18vICAgIFwgICAKICB8ICAgICAgICAgICBfbF9fbF8gICAg" +
			"ICAgICAgIEkgICAKICB9ICAgICAgICAgIFtfX19fX19dICAgICAgICAgICBJICAKICBdICAgICAgICAgICAgfCB8I" +
			"HwgICAgICAgICAgICB8ICAKICBdICAgICAgICAgICAgIH4gfiAgICAgICAgICAgICB8ICAKICB8ICAgICAgICAgIC" +
			"AgICAgICAgICAgICAgICAgIHwgICAKICAgfCAgICAgICAgICAgICAgICAgICAgICAgICAgIHwg")
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
		// TODO add an Int flag called "port", for the webserver
		// TODO add a String flag called "db", for the MongoDB connection string
		cli.StringFlag{
			Value:       dbType,
			Name:        "dbt, dt",
			Usage:       "Set the database type to use for the service (mongodb, postgresql or mock)",
			Destination: &dbType,
		},
		cli.StringFlag{
			Value:       migrationPath,
			Name:        "mp, m",
			Usage:       "Set the database migration folder path",
			Destination: &migrationPath,
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
		// TODO add a Duration flag called "statd" for the statistics duration (ex. 1h, 30s)
	}

	// main action
	// sub action are also possible
	app.Action = func(c *cli.Context) error {
		// print header
		fmt.Println(string(header))

		fmt.Print("* --------------------------------------------------- *\n")
		fmt.Printf("|   port                    : %d\n", port)
		fmt.Printf("|   db                      : %s\n", db)
		fmt.Printf("|   dbt                     : %s\n", dbType)
		fmt.Printf("|   mp                      : %s\n", migrationPath)
		fmt.Printf("|   logger level            : %s\n", logLevel)
		fmt.Printf("|   logger format           : %s\n", logFormat)
		fmt.Printf("|   statistic duration(s)   : %0.f\n", statisticsDuration.Seconds())
		fmt.Print("* --------------------------------------------------- *\n")

		// init log options from command line params
		err := utils.InitLog(logLevel, logFormat)
		if err != nil {
			logger.Warn("error setting log level, using debug as default")
		}

		// parse the database type
		dbt, err := dao.ParseDBType(dbType)
		if err != nil {
			return err
		}

		// build the web server
		webServer, err := web.BuildWebServer(db, migrationPath, dbt, statisticsDuration)

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

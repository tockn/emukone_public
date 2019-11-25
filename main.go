package main

import (
	"errors"
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/twitter"
	"github.com/speps/go-hashids"
	"github.com/tockn/emukone/infrastructure/repository/mysql"
	"github.com/tockn/emukone/interface/web/router"
	"github.com/tockn/emukone/registory"
	"log"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func main() {
	var (
		addr   = flag.String("addr", ":8081", "addr to bind")
		dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env    = flag.String("env", "development", "application envirionment (production, development etc.)")
	)
	flag.Parse()

	log.SetFlags(log.Lshortfile)

	rdbConn, err := getRdbConn(*dbconf, *env)
	if err != nil {
		log.Fatalf("alert: cannnot init rdb connection. %s", err)
	}

	store, err := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	if err != nil {
		log.Fatalf("cannot open redis connection. %s", err)
	}

	hash, err := getHashId()
	if err != nil {
		log.Fatal("hashid New error: ", err)
	}

	setGothProvider()

	frontEndpoint, ok := os.LookupEnv("FRONT_ENDPOINT")
	if !ok {
		frontEndpoint = "http://localhost:3000/"
	}

	awsSession, err := getAWSSession()
	if err != nil {
		log.Fatal("cannot create AWS session.", err)
	}

	i := registory.NewInteractor(rdbConn, hash, frontEndpoint, awsSession)
	h := i.NewAppHandler()
	r := gin.Default()
	router.NewRouter(r, h, store)
	if err := r.Run(*addr); err != nil {
		log.Fatal("Run error!!!")
	}
}

func getRdbConn(dbconf, env string) (*gorm.DB, error) {
	rdbConfs, err := mysql.NewDBConfigsFromFile(dbconf)
	if err != nil {
		return nil, err
	}
	rdbConn, err := rdbConfs[env].Open()
	if err != nil {
		return nil, err
	}

	file, _ := os.OpenFile("./mysql.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	rdbConn.LogMode(true)
	rdbConn.SetLogger(log.New(file, "", 0))
	return rdbConn, nil
}

func getHashId() (*hashids.HashID, error) {
	hd := hashids.NewData()
	hd.MinLength = 30
	return hashids.NewWithData(hd)
}

func setGothProvider() {
	// TODO 良い感じにKey取ってくる (credential用のS3 bucket作って、そこから取ってくるとかかなぁ)
	key := os.Getenv("TWITTER_KEY")
	secret := os.Getenv("TWITTER_SECRET")
	callback := os.Getenv("TWITTER_CALLBACK_URL")
	goth.UseProviders(
		twitter.NewAuthenticate(key, secret, callback),
	)
}

func getAWSSession() (*session.Session, error) {
	// TODO　環境変数以外で取ってくることになるかも？
	region, ok := os.LookupEnv("AWS_REGION")
	if !ok {
		return nil, errors.New("not set AWS_REGION")
	}
	key, ok := os.LookupEnv("AWS_KEY")
	if !ok {
		return nil, errors.New("not set AWS_KEY")
	}
	secret, ok := os.LookupEnv("AWS_SECRET")
	if !ok {
		return nil, errors.New("not set AWS_SECRET")
	}

	return session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(key, secret, ""),
			Region:      aws.String(region),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
}

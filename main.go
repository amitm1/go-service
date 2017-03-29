package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/amitm1/go-service/config"
	"github.com/amitm1/go-service/misc"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gorilla/mux"
	"gopkg.in/alexcesaro/statsd.v2"
	"net/http"
	"os"
	"time"
	// "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/ec2"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
)

var stats *statsd.Client
var conf *config.Config
var accesslog = new(log.TextFormatter)

const (
	RFC3339Milli = "2006-01-02T15:04:05.000Z07:00"
	SERVICE      = "svc"
)

func main() {
fmt.Println("I am here")
	/*svc := dynamodb.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tables:")
	for _, table := range result.TableNames {
		fmt.Println(*table)
	}
	*/
	/*
	   sess, err := session.NewSession()
	   	if err != nil {
	   		panic(err)
	   	}

	   	// Create an EC2 service object in the "us-west-2" region
	   	// Note that you can also configure your region globally by
	   	// exporting the AWS_REGION environment variable
	   	svc := ec2.New(sess, &aws.Config{Region: aws.String("us-west-2")})

	   	// Call the DescribeInstances Operation
	   	resp, err := svc.DescribeInstances(nil)
	   	if err != nil {
	   		panic(err)
	   	}

	   	// resp has all of the response data, pull out instance IDs:
	   	fmt.Println("> Number of reservation sets: ", len(resp.Reservations))
	   	for idx, res := range resp.Reservations {
	   		fmt.Println("  > Number of instances: ", len(res.Instances))
	   		for _, inst := range resp.Reservations[idx].Instances {
	   			fmt.Println("    - Instance ID: ", *inst.InstanceId)
	   		}
	   	}
	*/
	/*
	   fmt.Printf("starting")

	   	var endpoint = "http://127.0.0.1:8080"
	       //creds := credentials.NewStaticCredentials("asdasd", "adasdasd", "")
	       awsConfig := &aws.Config{
	           Credentials: creds,
	           Endpoint: &endpoint,
	       }

	       dynamodbconn := dynamodb.New('', awsConfig)

	       req := &dynamodb.DescribeTableInput{
	           TableName: aws.String("my_table"),
	       }

	       result, err := dynamodbconn.DescribeTable(req)

	       if err != nil {
	           fmt.Printf("%s", err)
	       }

	       table := result.Table

	          // some code

	       fmt.Printf("done")
	*/

	conf = config.GetConfig()
	if conf == nil {

	}

	accesslog.DisableColors = true
	accesslog.DisableTimestamp = true
	accesslog.DisableSorting = true

	stats, _ = statsd.New()
	defer stats.Close()

	router := NewRouter()
	log.Info("Starting up...")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	s := router.PathPrefix("/service").Subrouter()

	// Sets up the default service routes
	setupRoutes(s, serviceroutes)

	// Sets up all the routes the application is requesting.
	setupRoutes(router, routes)

	return router
}

func setupRoutes(router *mux.Router, routes Routes) {

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = SetupRequestHandler(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

}

func SetupRequestHandler(next http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We grab the time as early as possible in the request
		startRequestTime := time.Now().UTC()

		ll := log.New()
		ll.Out = os.Stdout

		ll.Formatter = new(log.TextFormatter)
		rh := misc.RequestHelpers{}
		rid := RequestId(r)

		rh.Logging = ll
		rh.Statsd = stats
		rh.Config = conf
		rh.RequestId = rid

		ctx := context.WithValue(r.Context(), "RequestHelper", &rh)

		next.ServeHTTP(w, r.WithContext(ctx))

		elapsed := time.Since(startRequestTime)

		//ll.Formatter = accesslog
		ll.WithFields(log.Fields{
			"t":   startRequestTime.Format(RFC3339Milli),
			"rid": rid,
			"rt":  elapsed.Nanoseconds() / 1e6, // Converting to milliseconds by dividing by a million

		}).Info()

	})
}

func RequestId(r *http.Request) string {
	rid := r.Header.Get("X-Request-ID")

	if rid != "" {
		return rid
	} else {
		return genRequestId(12, SERVICE)
	}
}

func genRequestId(strSize int, prefix string) string {

	dict := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz()*^%$#@!"

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dict[v%byte(len(dict))]
	}

	return prefix + "-" + string(bytes)
}

func GetCache(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)

	m["apple"] = "1"
	m["orange"] = "2"

	p, err := json.Marshal(m)
	if err != nil {
		log.Fatal("WTF")
	}

	ll := r.Context().Value("RequestHelper").(*misc.RequestHelpers).Logging

	ll.Info("Please help me")

	w.Write(p)

}

package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"
	"os"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/valyala/fasthttp/reuseport"

	"github.com/rgumi/whatsapp-mock/controller"
	"github.com/rgumi/whatsapp-mock/model"
)

var (
	apiPrefix      = flag.String("apiprefix", "/v1", "the prefix for the API")
	configFile     = flag.String("configfile", "/home/app/config.json", "the application config")
	addr           = flag.String("addr", "0.0.0.0:9090", "port the webserver listens on")
	signingKey     = []byte(*flag.String("skey", "abcde", "key which is used to sign jwt"))
	webhookURL     = flag.String("webhook", "", "URL of the webhook")
	enableTLS      = flag.Bool("tls", true, "run the API with TLS (HTTPS) enabled")
	soReuseport    = flag.Bool("sharding", false, "(experimental) uses SO_REUSEPORT option to start TCP listener") // see https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/
	staticAPIToken = os.Getenv("WA_API_KEY")
)

func readConfig(path string) *model.InternalConfig {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	config := new(model.InternalConfig)
	err = jsonpb.Unmarshal(f, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	flag.Parse()

	config := readConfig(*configFile)

	if *webhookURL != "" {
		config.WebhookUrl = *webhookURL
	}

	controller.Users = config.Users
	controller.UploadDir = config.UploadDir
	controller.SigningKey = signingKey

	contacts := make([]*model.Contact, len(config.Contacts))
	for i, c := range config.Contacts {
		contacts[i] = &model.Contact{
			WaId: c.Id,
			Profile: &model.Contact_Profile{
				Name: c.Name,
			},
		}
	}

	log.Printf("Creating new webserver with prefix %v", *apiPrefix)
	server := controller.NewServer(*apiPrefix, staticAPIToken)
	generators := model.NewGenerators(config.UploadDir, contacts, config.InboundMedia)
	webhook := controller.NewWebhookConfig(config.WebhookUrl, generators)

	controller.Webhook = webhook
	errors := make(chan error, 5)
	stopWebhook := webhook.Run(errors)

	go func() {
		for {
			err := <-errors
			log.Printf("ERROR " + err.Error())
		}
	}()
	var ln net.Listener
	var err error

	if *soReuseport {
		ln, err = reuseport.Listen("tcp4", *addr)
		if err != nil {
			log.Fatalf("Reuseport listener failed with %v", err)
		}
	} else {
		ln, err = net.Listen("tcp", *addr)
		if err != nil {
			log.Fatalf("Listener failed with %v", err)
		}
	}

	if *enableTLS {
		log.Println("Creating new Server TLS config as TLS is enabled")
		tlsCfg, err := controller.GenerateServerTLS()
		if err != nil {
			log.Fatalf("Unable to generate Server TLS config due to %v", err)
		}
		ln = tls.NewListener(ln, tlsCfg)
	}

	log.Printf("Starting webserver with addr %v", *addr)
	if err := server.Serve(ln); err != nil {
		log.Fatalf("Server listen failed with %v\n", err)
	}

	log.Printf("Shutting down application\n")
	ln.Close()
	stopWebhook <- 1
}

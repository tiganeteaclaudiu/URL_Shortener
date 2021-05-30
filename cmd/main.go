package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/fullstorydev/grpcui/standalone"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"jobtome.com/urlshortener/database"
	"jobtome.com/urlshortener/internal"
	proto "jobtome.com/urlshortener/proto"
)

func main() {
	run(context.Background())
}

// variable used to determine if /ui endpoint should expose GRPC UI (value of "yes" for true)
var enableGRPCUI = os.Getenv("ENABLE_GRPC_UI")

// variable used to determine if /ui endpoint should expose GRPC UI (value of "yes" for true)
var exposeDocs = os.Getenv("EXPOSE_DOCS")

// Ports at which the two instance services run
// If ran using docker-compose, these will not be exposed externally, only accessible via proxy

// GRPCPort : port at which GRPC Service which handles URL shortening runs
var GRPCPort = 4040

// HTTPPort : port at which HTTP Service which handles redirection and optionally GRPC UI runs
var HTTPPort = 8080

// Starts up services and handlers required for service:
//   GRPC service, HTTP call handlers, Redis instance, optionally GRPC UI
func run(ctx context.Context) {
	// seed randomizer
	rand.Seed(time.Now().UnixNano())

	//----------------------------
	// Start up Redis cache
	//----------------------------
	db, err := database.CreateRedisDatabase()
	if err != nil {
		panic(err)
	}

	//----------------------------
	// Initialize and start up gRPC server
	//----------------------------
	runGRPCServer(db)

	// define HTTP request multiplexer used to handle HTTP calls at different port than GRPC ones, but on same instance
	serveMux := http.NewServeMux()

	// conditionally expose service documentation at /docs path
	if exposeDocs == "yes" {
		exposeDocumentation(ctx, serveMux)
	}

	//----------------------------
	// Start up HTTP call handler
	//----------------------------
	fmt.Printf("Starting HTTP server on port %d...\n", HTTPPort)
	// start up HTTP multiplexer which handles redirections to shortened urls
	httpListener := runHTTPService(ctx, serveMux, db)

	//----------------------------
	// Start up GRPC UI
	//----------------------------
	// Conditionally expose GRPC UI instance at /ui endpoint of service
	// Handled by the HTTP multiplexer which detects GET requests at /ui endpoint
	if enableGRPCUI == "yes" {
		runGRPCUI(ctx, serveMux)
	}

	// Once all HTTP handlers have been defined, serve them using multiplexer
	err = http.Serve(httpListener, serveMux)
	if err != nil {
		log.Panicf("failed to serve HTTP: %v", err)
	}
}

// runGRPCServer starts up GRPC service which handles URL shortening CRUD
// service is exposed at separate port than HTTP handlers
func runGRPCServer(db database.Database) {
	fmt.Printf("Starting gRPC server on port %d...\n", GRPCPort)
	svr := grpc.NewServer()
	// register UrlShortener service
	proto.RegisterUrlShortenerServiceServer(svr, service.InitializeService(db))

	// enable reflection (used for GRPC UI, provides protobuf docs)
	reflection.Register(svr)

	// listen to specific port which will handle GRPC calls
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", GRPCPort))
	if err != nil {
		log.Panicf("failed to open listen socket on port %d: %v", GRPCPort, err)
	}

	// serve GRPC service
	go func() {
		err := svr.Serve(l)
		if err != nil {
			log.Panicf("failed to serve gRPC: %v", err)
		}
	}()
}

// runGRPCUI starts up GRPC UI, used as both documentation and playground platform for GRPC service
func runGRPCUI(ctx context.Context, mux *http.ServeMux) {
	// initialize internal client to GRPC service (used exclusively for GRPC UI)
	cc, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", GRPCPort), grpc.WithInsecure())
	if err != nil {
		log.Panicf("failed to create client to GRPC service for GRPC UI: %v", err)
	}
	target := fmt.Sprintf("%s:%d", filepath.Base(os.Args[0]), GRPCPort)

	// Pull out pre-defined handler which exposes GRPC UI
	h, err := standalone.HandlerViaReflection(ctx, cc, target)
	if err != nil {
		log.Panicf("failed to create client to local server: %v", err)
	}

	// register GRPC UI handler at /ui endpoint of HTML calls handler
	// (note use of mux, which defines HTTP calls multiplexer)
	mux.Handle("/ui/", http.StripPrefix("/ui", h))
}

// runHTTPService starts up HTTP calls multiplexer, which handler redirect to shortened URL's
// Uses Redis to fetch url to redirect to, based on key previously stored in DB
func runHTTPService(ctx context.Context, mux *http.ServeMux, db database.Database) net.Listener {
	// define new handler at root path of HTTP calls handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// define default status as SeeOther (default for redirecting to external URL)
		status := http.StatusSeeOther
		// Attempt to fetch shortened URL by key using Redis call
		value, err := db.Get(ctx, r.URL.String()[1:])
		if err != nil {
			fmt.Printf("ERROR getting shortened URL: %s", err.Error())
			status = http.StatusNotFound
		}
		// redirect to either shortened URL or 404 page
		http.Redirect(w, r, value, status)
	})

	// enable listener for HTTP calls at root path
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", HTTPPort))
	if err != nil {
		log.Panicf("failed to open HTTP listen socket on port %d: %v", HTTPPort, err)
	}

	return l
}

// exposeDocumentation exposes documentation generated automatically on "go generate" at /docs/ path
func exposeDocumentation(ctx context.Context, mux *http.ServeMux) {
	fmt.Printf("Exposing documentation at %d/docs...\n", HTTPPort)
	mux.HandleFunc("/docs/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("documentation")
		p := "./docs/index.html"
		http.ServeFile(w, r, p)
	})
}

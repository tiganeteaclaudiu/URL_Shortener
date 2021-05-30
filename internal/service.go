package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"

	"jobtome.com/urlshortener/database"
	proto "jobtome.com/urlshortener/proto"
)

// Defines Redis cache client
var db database.Database

// UrlShortener : handles operations regarding url shortening.
type UrlShortener struct {
	proto.UnimplementedUrlShortenerServiceServer
}

// Load proxy URL from env variable
// Proxy URL is used to redirect the user
var proxyUrl = os.Getenv("PROXY_URL")

// InitializeService initializes GRPC service, defines global Redis client
func InitializeService(serviceDB database.Database) *UrlShortener {
	if proxyUrl == "" {
		fmt.Println("[WARNING] PROXY env variable NOT SET. Defaulting to HTTP calls multiplexer host")
		proxyUrl = "localhost:8080"
	}

	db = serviceDB
	return &UrlShortener{}
}

// GetShortenedUrl fetches already existing shortened URL by key
func (u UrlShortener) GetShortenedUrl(ctx context.Context, i *proto.Key) (*proto.Url, error) {
	value, err := db.Get(ctx, i.GetKey())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error getting entry: %s", err.Error()))
	}
	// return shortened URL
	return &proto.Url{Url: value}, nil
}

// SetShortenedUrl creates new shortened URL, caches it, sets optional expiry, and returns it in response
// If expiry set operation fails, entry is deleted before returning error.
func (u UrlShortener) SetShortenedUrl(ctx context.Context, i *proto.SetShortenedUrlInput) (*proto.Url, error) {
	// Create new random sequence of n length
	rnd := randSeq(14)

	// attempt to associate random sequence to received URL in cache
	_, err := db.Set(ctx, rnd, i.GetUrl())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating shortened URL entry: %s", err.Error()))
	}

	// attempt to set expiry to new entry IF expiry time is defined in input
	if i.GetExpiryMinutes() != 0 {
		_, err = db.Expire(ctx, rnd, time.Duration(i.GetExpiryMinutes())*time.Minute)
		if err != nil {
			// If EXPIRY call fails, delete entry before returning
			// This is done to avoid persisting entries which the user does not want to keep indefinitely
			_, deleteErr := u.DeleteShortenedUrl(ctx, &proto.Key{Key: rnd})
			if deleteErr != nil {
				fmt.Println("[CRITICAL] Failed to delete entry after EXPIRE operation failed.")
				return nil, errors.Wrap(errors.New("Failed to set expiry of entry"), deleteErr.Error())
			}

			return nil, errors.Wrap(err, "Failed to set expiry on entry: Entry has been deleted")
		}
	}

	// return new shortened URL
	return &proto.Url{Url: fmt.Sprintf("%s/%s", proxyUrl, rnd)}, nil
}

// DeleteShortenedUrl deletes a shortened URL by it's value
// example input: "ljkshbdlfybsd"
func (u UrlShortener) DeleteShortenedUrl(ctx context.Context, i *proto.Key) (*proto.Void, error) {
	_, err := db.Delete(ctx, i.GetKey())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error deleting entry: %s", err.Error()))
	}
	return &proto.Void{}, nil
}

package gatekeeper

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type AnonRequestDTO struct {
	IP                string    `db:"ip"`
	ReqCount          int       `db:"req_count"`
	ReqLimit          int       `db:"req_limit"`
	ReqAllocationTime time.Time `db:"req_allocation_time"`
}

type AnonGateKeeper interface {
	CanPerformRequest(ctx context.Context, ip string, usage string) error
	RecordRequest(ctx context.Context, ip string, usage string) error
}

func getClientIP(httpRequest *http.Request) string {
	clientIp := httpRequest.RemoteAddr
	forwardedAddr, ok := httpRequest.Header["X-Forwarded-For"]

	if ok {
		logrus.Infof("Request remoteAddr=%s, ForwardedFor0=%s, len=%v", clientIp, forwardedAddr[0], len(forwardedAddr))
		ips := strings.Split(forwardedAddr[0], ", ")
		clientIp = ips[0]
	} else {
		clientIp = strings.Split(clientIp, ":")[0]
	}
	return clientIp
}

func CanPerhormHTTPRequest(ctx context.Context, httpRequest *http.Request, resource string, anonGateKeeper AnonGateKeeper) error {
	_, ok := httpRequest.Header["X-FEATHER-TOKEN"]
	if ok == false {
		clientIp := getClientIP(httpRequest)
		// Anonymous user - throttle
		err := anonGateKeeper.CanPerformRequest(ctx, clientIp, resource)
		if err != nil {
			return err
		}
	} else {
		// Authenticated user - validate auth, then check if the user has enough credentials for this call
	}
	return nil
}

func RecordHTTPRequest(ctx context.Context, httpRequest *http.Request, resource string, anonGateKeeper AnonGateKeeper) error {
	_, ok := httpRequest.Header["X-FEATHER-TOKEN"]
	if ok == false {
		clientIp := getClientIP(httpRequest)
		err := anonGateKeeper.RecordRequest(ctx, clientIp, resource)
		if err != nil {
			return err
		}
	} else {
		// Authenticated user - validate auth, then check if the user has enough credentials for this call
	}
	return nil

}

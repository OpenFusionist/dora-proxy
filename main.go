package main

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
    log.SetFormatter(&logrus.TextFormatter{
        FullTimestamp:   true,
        TimestampFormat: "2006-01-02 15:04:05",
    })

	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	upstream, err := url.Parse(cfg.UpstreamBaseURL)
	if err != nil {
		log.Fatalf("invalid PROXY_UPSTREAM_BASE_URL: %v", err)
	}

	client := &http.Client{Timeout: 20 * time.Second}

	// Initialize attestation cache and tracker
	cache := NewLastAttestCache()
	tracker := NewAttestationTracker(client, cfg.ConsensusAPIURL, cache, log)
	// Kick off startup backfill (best-effort) and periodic epoch scans
	go func() {
		log.Info("starting attestation backfill (last 3 epochs)")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		if err := tracker.Backfill(ctx); err != nil {
			log.WithError(err).Warn("backfill returned with error")
		}
		cancel()
		log.Info("attestation backfill finished")
	}()
	tracker.Start()

	r := buildRouter(cfg, client, upstream, cache)

	srv := &http.Server{
		Addr:         cfg.ListenAddr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Infof("dora-proxy listening on %s, upstream=%s, consensus_api=%s", cfg.ListenAddr, cfg.UpstreamBaseURL, cfg.ConsensusAPIURL)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("proxy server error: %v", err)
	}
}

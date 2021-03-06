package e2etest

import (
	"context"
	"crypto/rand"
	"io"
	"sync"

	"go.uber.org/zap"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/internal/tgtest"
)

// TestConfig contains some common test server settings.
type TestConfig struct {
	AppID   int
	AppHash string
	DcID    int
	Addr    string
}

// Suite is struct which contains external E2E test parameters.
type Suite struct {
	tgtest.Suite
	TestConfig
	rand io.Reader
	// already used phone numbers
	used    map[string]struct{}
	usedMux sync.Mutex
}

// NewSuite creates new Suite.
func NewSuite(suite tgtest.Suite, config TestConfig, randomSource io.Reader) *Suite {
	return &Suite{
		Suite:      suite,
		TestConfig: config,
		rand:       randomSource,
		used:       map[string]struct{}{},
	}
}

// Client creates new *telegram.Client using this suite.
func (s *Suite) Client(logger *zap.Logger, handler telegram.UpdateHandler) *telegram.Client {
	return telegram.NewClient(s.AppID, s.AppHash, telegram.Options{
		Addr:          s.Addr,
		Logger:        logger,
		UpdateHandler: handler,
	})
}

// Authenticate authenticates client on test server.
func (s *Suite) Authenticate(ctx context.Context, client *telegram.Client) error {
	var auth telegram.UserAuthenticator
	for {
		auth = telegram.TestAuth(rand.Reader, s.DcID)
		phone, err := auth.Phone(ctx)
		if err != nil {
			return err
		}

		s.usedMux.Lock()
		if _, ok := s.used[phone]; !ok {
			s.used[phone] = struct{}{}
			s.usedMux.Unlock()
			break
		}
		s.usedMux.Unlock()
	}

	return telegram.NewAuth(
		auth,
		telegram.SendCodeOptions{},
	).Run(ctx, client)
}

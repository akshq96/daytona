// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package headscale

import (
	"os"
	"path/filepath"

	"github.com/juanfont/headscale/hscontrol"
)

type HeadscaleServer struct {
	ServerId      string
	FrpsDomain    string
	HeadscalePort uint32
}

func (s *HeadscaleServer) Init() error {
	headscaleConfigDir, err := s.getHeadscaleConfigDir()
	if err != nil {
		return err
	}

	return os.MkdirAll(headscaleConfigDir, 0700)
}

func (s *HeadscaleServer) Start() error {
	cfg, err := s.getHeadscaleConfig()
	if err != nil {
		return err
	}

	app, err := hscontrol.NewHeadscale(cfg)
	if err != nil {
		return err
	}

	return app.Serve()
}

func (s *HeadscaleServer) getHeadscaleConfigDir() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userConfigDir, "daytona", "server", "headscale"), nil
}

/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msppvdr

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vtbaas/vbaas-go-sdk/pkg/core/config"
	"github.com/vtbaas/vbaas-go-sdk/pkg/core/cryptosuite"
	"github.com/vtbaas/vbaas-go-sdk/pkg/fab"
	"github.com/vtbaas/vbaas-go-sdk/pkg/fabsdk/factory/defcore"
	"github.com/vtbaas/vbaas-go-sdk/pkg/msp"
	"github.com/vtbaas/vbaas-go-sdk/pkg/msp/test/mockmsp"
	"github.com/vtbaas/vbaas-go-sdk/test/metadata"
)

func TestCreateMSPProvider(t *testing.T) {

	coreFactory := defcore.NewProviderFactory()

	configPath := filepath.Join(metadata.GetProjectPath(), metadata.SDKConfigPath, "config_test.yaml")
	configBackend, err := config.FromFile(configPath)()
	if err != nil {
		t.Fatalf(err.Error())
	}

	cryptoSuiteConfig := cryptosuite.ConfigFromBackend(configBackend...)
	if err != nil {
		t.Fatalf(err.Error())
	}

	endpointConfig, err := fab.ConfigFromBackend(configBackend...)
	if err != nil {
		t.Fatalf(err.Error())
	}

	cryptosuite, err := coreFactory.CreateCryptoSuiteProvider(cryptoSuiteConfig)
	if err != nil {
		t.Fatalf("Unexpected error creating cryptosuite provider %s", err)
	}

	userStore := &mockmsp.MockUserStore{}

	provider, err := New(endpointConfig, cryptosuite, userStore)
	assert.Nil(t, err, "New should not have failed")

	if provider.UserStore() != userStore {
		t.Fatal("Invalid user store returned")
	}

	mgr, ok := provider.IdentityManager("Org1")
	if !ok {
		t.Fatal("Expected to return idnetity manager")
	}

	_, ok = mgr.(*msp.IdentityManager)
	if !ok {
		t.Fatal("Unexpected signing manager created")
	}
}

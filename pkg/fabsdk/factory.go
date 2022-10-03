/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabsdk

import (
	"github.com/vtbaas/vbaas-go-sdk/pkg/core/logging/api"
	sdkApi "github.com/vtbaas/vbaas-go-sdk/pkg/fabsdk/api"
)

// pkgSuite provides the package factories that create clients and providers
type pkgSuite interface {
	Core() (sdkApi.CoreProviderFactory, error)
	MSP() (sdkApi.MSPProviderFactory, error)
	Service() (sdkApi.ServiceProviderFactory, error)
	Logger() (api.LoggerProvider, error)
}

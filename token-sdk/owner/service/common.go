/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package service

import (
	"os"

	viewregistry "github.com/hyperledger-labs/fabric-smart-client/platform/view"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
	"github.com/hyperledger-labs/fabric-token-sdk/token/services/ttx"
	"github.com/pkg/errors"
)

// NewTransaction creates the envelope for the transaction and specify the auditor
func NewTransaction(context view.Context) (*ttx.Transaction, error) {
	if os.Getenv("DISABLE_AUDITOR") == "true" {
		return ttx.NewTransaction(context, nil)
	}

	logger.Debug("getting identity of auditor")
	auditor := viewregistry.GetIdentityProvider(context).Identity("auditor") // TODO: should not be hardcoded
	if auditor == nil {
		return nil, errors.New("auditor identity not found")
	}
	return ttx.NewTransaction(context, nil, ttx.WithAuditor(auditor))
}

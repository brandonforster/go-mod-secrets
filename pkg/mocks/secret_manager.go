/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
// Package mocks contains entities which can be used for general testing.
package mocks

import (
	"github.com/edgexfoundry/go-mod-secrets/pkg/errors"
)

// MockSecretStoreManager a mock implementation of the SecretStoreManager which can be used for testing.
type MockSecretStoreManager struct {
	Secrets map[string]string
}

func (mssm MockSecretStoreManager) GetValue(key string) (string, error) {
	value, _ := mssm.Secrets[key]
	return value, nil
}

func (mssm MockSecretStoreManager) SetKeyValue(key string, value string) error {
	if value == "" {
		return errors.ErrUnsupportedValue{}
	}

	mssm.Secrets[key] = value
	return nil
}

func (mssm MockSecretStoreManager) DeleteKeyValue(key string) error {
	_, ok := mssm.Secrets[key]
	if !ok {
		return errors.ErrSecretNotFound{}
	}

	delete(mssm.Secrets, key)
	return nil
}

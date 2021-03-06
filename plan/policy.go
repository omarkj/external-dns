/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plan

import (
	log "github.com/Sirupsen/logrus"
)

// Policy allows to apply different rules to a set of changes.
type Policy interface {
	Apply(changes *Changes) *Changes
}

// Policies is a registry of available policies.
var Policies = map[string]Policy{
	"sync":        &SyncPolicy{},
	"upsert-only": &UpsertOnlyPolicy{},
}

// SyncPolicy allows for full synchronization of DNS records.
type SyncPolicy struct{}

// Apply applies the sync policy which returns the set of changes as is.
func (p *SyncPolicy) Apply(changes *Changes) *Changes {
	return changes
}

// UpsertOnlyPolicy allows evrything but deleting DNS records.
type UpsertOnlyPolicy struct{}

// Apply applies the upsert-only policy which strips out any deletions.
func (p *UpsertOnlyPolicy) Apply(changes *Changes) *Changes {
	log.Debugf("Records to be excluded from the delete list due to upsert-only policy: %v", changes.Delete)
	return &Changes{
		Create:    changes.Create,
		UpdateOld: changes.UpdateOld,
		UpdateNew: changes.UpdateNew,
	}
}

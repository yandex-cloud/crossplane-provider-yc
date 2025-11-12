/*
Copyright 2022 YANDEX LLC

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

package mdb

import (
	"encoding/json"
	"reflect"
	"testing"
)

const attrsJSON = `{
    "config": [{
        "access": [{
            "data_lens": false,
            "web_sql": false
        }],
        "autofailover": true,
        "backup_retain_period_days": 7,
        "backup_window_start": [{
            "hours": 0,
            "minutes": 0
        }],
        "performance_diagnostics": [{
            "enabled": false,
            "sessions_sampling_interval": 60,
            "statements_sampling_interval": 600
        }],
        "pooler_config": [],
        "postgresql_config": {
            "autovacuum_vacuum_scale_factor": "0.34",
            "default_transaction_isolation": "TRANSACTION_ISOLATION_READ_COMMITTED",
            "enable_parallel_hash": "true",
            "max_connections": "395",
            "shared_preload_libraries": "SHARED_PRELOAD_LIBRARIES_AUTO_EXPLAIN",
            "vacuum_cleanup_index_scale_factor": "0.2"
        },
        "resources": [{
            "disk_size": 16,
            "disk_type_id": "network-hdd",
            "resource_preset_id": "s2.micro"
        }],
        "version": "12"
    }],
    "created_at": "2021-12-24T08:51:11Z",
    "database": [{
        "extension": [],
        "lc_collate": "C",
        "lc_type": "C",
        "name": "db_name",
        "owner": "user1"
    }],
    "deletion_protection": false,
    "description": "",
    "environment": "PRESTABLE",
    "folder_id": "b1gj2tg21doe4mcdr530",
    "health": "ALIVE",
    "host": [{
        "assign_public_ip": false,
        "fqdn": "rc1a-y0e4nzvcplls1gsf.mdb.yandexcloud.net",
        "name": "",
        "priority": 0,
        "replication_source": "",
        "replication_source_name": "",
        "role": "MASTER",
        "subnet_id": "e9bps9emqo5refh5rh9c",
        "zone": "ru-central1-a"
    }, {
        "assign_public_ip": false,
        "fqdn": "rc1b-ob8noo5hcc01rtcf.mdb.yandexcloud.net",
        "name": "",
        "priority": 0,
        "replication_source": "",
        "replication_source_name": "",
        "role": "REPLICA",
        "subnet_id": "e2l2ljo993ihcmeve6tl",
        "zone": "ru-central1-b"
    }, {
        "assign_public_ip": false,
        "fqdn": "rc1c-a9byrizp9deqa49b.mdb.yandexcloud.net",
        "name": "",
        "priority": 0,
        "replication_source": "",
        "replication_source_name": "",
        "role": "REPLICA",
        "subnet_id": "b0c1jhn7mg5ub1l8fokt",
        "zone": "ru-central1-c"
    }],
    "host_master_name": "",
    "id": "c9q3n7toi2jc0j7et6tp",
    "labels": {},
    "maintenance_window": [{
        "day": "SAT",
        "hour": 12,
        "type": "WEEKLY"
    }],
    "name": "example-postgesql",
    "network_id": "enp42t1n32u1r35t1qm3",
    "restore": [],
    "security_group_ids": [],
    "status": "RUNNING",
    "timeouts": null,
    "user": [{
        "conn_limit": 50,
        "grants": [],
        "login": true,
        "name": "user1",
        "password": "12345678",
        "permission": [{
            "database_name": "db_name"
        }],
        "settings": {
            "default_transaction_isolation": "read committed",
            "log_min_duration_statement": "5000"
        }
    }]
}`

func Test_connDetails(t *testing.T) {
	var attrs map[string]interface{}
	_ = json.Unmarshal([]byte(attrsJSON), &attrs)

	tests := []struct {
		attr map[string]interface{}
		want map[string][]byte
	}{{
		attr: attrs,
		want: map[string][]byte{
			"attribute.host.0.fqdn":           []byte("rc1a-y0e4nzvcplls1gsf.mdb.yandexcloud.net"),
			"attribute.host.1.fqdn":           []byte("rc1b-ob8noo5hcc01rtcf.mdb.yandexcloud.net"),
			"attribute.host.2.fqdn":           []byte("rc1c-a9byrizp9deqa49b.mdb.yandexcloud.net"),
			"attribute.database.0.name":       []byte("db_name"),
			"attribute.user.0.name":           []byte("user1"),
			"connection-string.user1.db_name": []byte("host=rc1a-y0e4nzvcplls1gsf.mdb.yandexcloud.net,rc1b-ob8noo5hcc01rtcf.mdb.yandexcloud.net,rc1c-a9byrizp9deqa49b.mdb.yandexcloud.net port=6432 sslmode=verify-full dbname=db_name user=user1 target_session_attrs=read-write password=12345678"),
		},
	}}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := postgresqlConnDetails(tt.attr)
			if !reflect.DeepEqual(got, tt.want) {
				gotJSON, _ := json.MarshalIndent(got, "", "  ")
				wantJSON, _ := json.MarshalIndent(tt.want, "", "  ")
				t.Errorf("postgresqlConnDetails() got = %s, want %s", gotJSON, wantJSON)
			}
		})
	}
}

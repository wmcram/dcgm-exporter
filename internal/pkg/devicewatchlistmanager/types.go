/*
 * Copyright (c) 2024, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//go:generate go run -v go.uber.org/mock/mockgen  -destination=../../mocks/pkg/devicewatchlistmanager/mock_device_watchlist_manager.go -package=devicewatchlistmanager -copyright_file=../../../hack/header.txt . Manager

package devicewatchlistmanager

import (
	"github.com/NVIDIA/go-dcgm/pkg/dcgm"

	"github.com/wmcram/dcgm-exporter/internal/pkg/devicewatcher"
)

type Manager interface {
	CreateEntityWatchList(dcgm.Field_Entity_Group, devicewatcher.Watcher, int64) error
	EntityWatchList(dcgm.Field_Entity_Group) (WatchList, bool)
}

/*
 * Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.
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

package main

import (
	"log/slog"
	"os"

	"github.com/wmcram/dcgm-exporter/pkg/cmd"
)

var BuildVersion = "Filled by the build system"

func main() {
	app := cmd.NewApp(BuildVersion)
	if err := app.Run(os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

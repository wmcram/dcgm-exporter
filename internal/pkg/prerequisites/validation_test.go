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

package prerequisites

import (
	debugelf "debug/elf"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	mockelf "github.com/wmcram/dcgm-exporter/internal/mocks/pkg/elf"
	mockexec "github.com/wmcram/dcgm-exporter/internal/mocks/pkg/exec"
	mockos "github.com/wmcram/dcgm-exporter/internal/mocks/pkg/os"

	realos "os"
)

func TestValidate(t *testing.T) {
	type testCase struct {
		Name               string
		OSMockExpectations func(*gomock.Controller, *mockos.MockOS)
		LDConfigPath       string
	}

	tests := []testCase{
		{
			Name: "Ubuntu-based system with /sbin/ldconfig.real",
			OSMockExpectations: func(ctrl *gomock.Controller, mo *mockos.MockOS) {
				mfi := mockos.NewMockFileInfo(ctrl)
				mo.EXPECT().Stat("/sbin/ldconfig.real").Return(mfi, nil)
			},
			LDConfigPath: "/sbin/ldconfig.real",
		},
		{
			Name: "Linux system without /sbin/ldconfig.real",
			OSMockExpectations: func(ctrl *gomock.Controller, mo *mockos.MockOS) {
				mo.EXPECT().Stat("/sbin/ldconfig.real").Return(nil, &realos.PathError{})
			},
			LDConfigPath: "/sbin/ldconfig",
		},
	}

	for _, tc := range tests {

		ctrl := gomock.NewController(t)

		osinstance := mockos.NewMockOS(ctrl)
		tc.OSMockExpectations(ctrl, osinstance)

		os = osinstance

		executor := mockexec.NewMockExec(ctrl)

		output := `1211 libs found in cache '/etc/ld.so.cache'
					libdcgm.so.4 (libc6,x86-64) => /lib/x86_64-linux-gnu/libdcgm.so.4
				Cache generated by: ldconfig (Ubuntu GLIBC 2.35-0ubuntu3.7) stable release version 2.35`
		cmd := mockexec.NewMockCmd(ctrl)
		cmd.EXPECT().Output().AnyTimes().Return([]byte(output), nil)
		executor.EXPECT().Command(gomock.Eq(tc.LDConfigPath), gomock.Eq(ldconfigParam)).AnyTimes().Return(cmd)

		exec = executor

		elfreader := mockelf.NewMockELF(ctrl)

		self := &debugelf.File{
			FileHeader: debugelf.FileHeader{
				Machine: debugelf.EM_X86_64,
			},
		}
		elfreader.EXPECT().Open(gomock.Eq("/proc/self/exe")).AnyTimes().Return(self, nil)

		libdcgm := &debugelf.File{
			FileHeader: debugelf.FileHeader{
				Machine: debugelf.EM_X86_64,
			},
		}
		elfreader.EXPECT().Open(gomock.Eq("/lib/x86_64-linux-gnu/libdcgm.so.4")).AnyTimes().Return(libdcgm, nil)

		elf = elfreader

		err := Validate()
		require.NoError(t, err)
	}
}

// Copyright 2019 IBM Corp
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"github.com/spf13/cobra"
)

var (
	appLong = `
		Create modified YAML with encrypted configmap and scratch image`

	appExample = `
		# This command gets invoked via 'image' command`
)

func NewCmdApp() *cobra.Command {
	cmds := &cobra.Command{
		Use:                   "app",
		DisableFlagsInUseLine: true,
		Short:                 "SecureContainer App",
		Long:                  appLong,
		Example:               appExample,
		Hidden:                true,
	}
	cmds.AddCommand(NewCmdAppCreate())
	return cmds
}

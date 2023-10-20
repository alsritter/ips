/*
 * Copyright (c) 2023 shenjunzheng@gmail.com
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

package ips

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/sjzar/ips/internal/config"
	"github.com/sjzar/ips/internal/ips"
)

const (
	ConfigName = "ips"
	ConfigType = "json"
	EnvIPSDir  = "IPS_DIR"
)

// Global variables storing command-line or configuration values.
var (

	// logLevel specifies the logging level for the application.
	logLevel string

	// operate
	// field select
	// fields specifies the fields to output.
	fields string

	// useDBFields indicates whether to use database fields. (default is common fields)
	useDBFields bool

	// rewriter
	// rewriteFiles specifies the files for data rewriting.
	rewriteFiles string

	// root command flags
	// database
	// rootFormat defines the format for database.
	rootFormat string

	// rootFile specifies the file path for database.
	rootFile string

	// rootIPv4Format defines the format for IPv4 database.
	rootIPv4Format string

	// rootIPv4File specifies the file path for IPv4 database.
	rootIPv4File string

	// rootIPv6Format defines the format for IPv6 database.
	rootIPv6Format string

	// rootIPv6File specifies the file path for IPv6 database.
	rootIPv6File string

	// output
	// rootTextFormat defines the format for text output.
	rootTextFormat string

	// rootTextValuesSep defines the separator for text output.
	rootTextValuesSep string

	// rootJson defines whether to output in JSON format.
	rootJson bool

	// rootJsonIndent defines whether to output in indented JSON format.
	rootJsonIndent bool

	// dump & pack command flags
	// operate
	// dpFields specifies the fields to output for dump and pack operations.
	dpFields string

	// dpRewriterFiles specifies the files for data rewriting during dump and pack operations.
	dpRewriterFiles string

	// inputFile specifies the input file for dump and pack operations.
	inputFile string

	// inputFormat specifies the input format for dump and pack operations.
	inputFormat string

	// outputFile specifies the output file for dump and pack operations.
	outputFile string

	// outputFormat specifies the output format for pack operations.
	outputFormat string
)

// GetFlagConfig initializes and returns the configuration for the IP service.
// It loads the configuration from a file or environment variables, then overrides
// specific fields based on the global variables, which might be set from command-line arguments.
func GetFlagConfig() *ips.Config {
	conf := GetConfig()

	// operate
	if len(fields) != 0 {
		conf.Fields = fields
	}

	if useDBFields {
		conf.UseDBFields = useDBFields
	}

	if len(rewriteFiles) != 0 {
		conf.RewriteFiles = rewriteFiles
	}

	// root command flags
	if len(rootFile) != 0 {
		conf.IPv4File = rootFile
		conf.IPv6File = rootFile
		if len(rootFormat) != 0 {
			conf.IPv4Format = rootFormat
			conf.IPv6Format = rootFormat
		}
	}

	if len(rootIPv4File) != 0 {
		conf.IPv4File = rootIPv4File
		if len(rootIPv4Format) != 0 {
			conf.IPv4Format = rootIPv4Format
		}
	}

	if len(rootIPv6File) != 0 {
		conf.IPv6File = rootIPv6File
		if len(rootIPv6Format) != 0 {
			conf.IPv6Format = rootIPv6Format
		}
	}

	if len(rootTextFormat) != 0 {
		conf.TextFormat = rootTextFormat
	}

	if len(rootTextValuesSep) != 0 {
		conf.TextValuesSep = rootTextValuesSep
	}

	if rootJson {
		conf.OutputType = ips.OutputTypeJSON
	}

	if rootJsonIndent {
		conf.OutputType = ips.OutputTypeJSON
		conf.JsonIndent = rootJsonIndent
	}

	// dump & pack command flags
	if len(dpFields) != 0 {
		conf.DPFields = dpFields
	}

	if len(dpRewriterFiles) != 0 {
		conf.DPRewriterFiles = dpRewriterFiles
	}

	return conf
}

// GetConfig initializes and returns the configuration for the IP service.
func GetConfig() *ips.Config {
	conf := &ips.Config{}
	if err := config.Init(ConfigName, ConfigType, os.Getenv(EnvIPSDir)); err != nil {
		log.Fatal(err)
	}
	if err := config.Load(conf); err != nil {
		log.Fatal(err)
	}
	conf.IPSDir = config.ConfigPath
	return conf
}
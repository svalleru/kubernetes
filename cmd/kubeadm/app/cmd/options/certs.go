/*
Copyright 2018 The Kubernetes Authors.

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

package options

import "github.com/spf13/pflag"

// AddCertificateDirFlag adds the --certs-dir flag to the given flagset
func AddCertificateDirFlag(fs *pflag.FlagSet, certsDir *string) {
	fs.StringVar(certsDir, CertificatesDir, *certsDir, "The path where to save the certificates")
}

// AddCSRFlag adds the --csr-only flag to the given flagset
func AddCSRFlag(fs *pflag.FlagSet, csr *bool) {
	fs.BoolVar(csr, CSROnly, *csr, "Create CSRs instead of generating certificates")
}

// AddCSRDirFlag adds the --csr-dir flag to the given flagset
func AddCSRDirFlag(fs *pflag.FlagSet, csrDir *string) {
	fs.StringVar(csrDir, CSRDir, *csrDir, "The path to output the CSRs and private keys to")
}

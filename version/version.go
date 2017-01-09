package version

import "strings"

// Version describes the Goad version.
const Version = "1.3.0"

// LambdaVersion returns a version string that can be used as a Lambda function
// alias.
func LambdaVersion() string {
	return "v" + strings.Replace(Version, ".", "-", -1)
}

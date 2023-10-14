package shared

import "regexp"

var ONLY_NUMBERS_REGEX = regexp.MustCompile(`\d.*`)
var NON_NUMERIC_REGEX = regexp.MustCompile(`\D`)

package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	httpErrs "github.com/storybuilder/storybuilder/channels/http/errors"
	"github.com/storybuilder/storybuilder/channels/http/request/unpackers"
)

// Unpack the request in to the given unpacker struct.
func Unpack(r *http.Request, unpacker unpackers.UnpackerInterface) error {
	err := json.NewDecoder(r.Body).Decode(unpacker)
	if err != nil {
		return httpErrs.NewValidationError(formatUnpackerMessage(unpacker.RequiredFormat()))
	}

	return nil
}

// formatUnpackerMessage removes any special chatacters from the message string.
func formatUnpackerMessage(p string) string {
	// catch carrage returns and new lines
	reNewLine := regexp.MustCompile(`[\r\n]+`)

	// catch other special characters
	reSpecialChar := regexp.MustCompile(`[\t\"\']*`)

	m := reSpecialChar.ReplaceAllString(reNewLine.ReplaceAllString(p, " "), "")

	return fmt.Sprintf("Required format: %s", m)
}

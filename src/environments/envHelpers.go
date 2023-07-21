// portainerImporter
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/environments/envHelpers.go
// Original time: 2023/07/17 15:00

package environments

import (
	"encoding/json"
	"io"
)

func decodeJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func environmentInventory(printList bool) string {

}

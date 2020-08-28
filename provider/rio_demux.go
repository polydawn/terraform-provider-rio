// copied from github.com/polydawn/rio/cmd/rio/demuxes.go
package provider

import (
	. "github.com/warpfork/go-errcat"

	"go.polydawn.net/go-timeless-api/rio"
	"go.polydawn.net/rio/transmat/git"
	tartrans "go.polydawn.net/rio/transmat/tar"
	ziptrans "go.polydawn.net/rio/transmat/zip"
)

var packTypes = []string{
	string(tartrans.PackType),
	string(ziptrans.PackType),
	string(git.PackType),
}

func isPackType(packType string) bool {
	for _, p := range packTypes {
		if packType == p {
			return true
		}
	}
	return false
}

func demuxPackTool(packType string) (rio.PackFunc, error) {
	switch packType {
	case string(tartrans.PackType):
		return tartrans.Pack, nil
	case string(ziptrans.PackType):
		return ziptrans.Pack, nil
	default:
		return nil, Errorf(rio.ErrUsage, "unsupported packtype %q", packType)
	}
}

func demuxUnpackTool(packType string) (rio.UnpackFunc, error) {
	switch packType {
	case string(tartrans.PackType):
		return tartrans.Unpack, nil
	case string(git.PackType):
		return git.Unpack, nil
	case string(ziptrans.PackType):
		return ziptrans.Unpack, nil
	default:
		return nil, Errorf(rio.ErrUsage, "unsupported packtype %q", packType)
	}
}

func demuxScanTool(packType string) (rio.ScanFunc, error) {
	switch packType {
	case string(tartrans.PackType):
		return tartrans.Scan, nil
	case string(ziptrans.PackType):
		return ziptrans.Scan, nil
	default:
		return nil, Errorf(rio.ErrUsage, "unsupported packtype %q", packType)
	}
}

func demuxMirrorTool(packType string) (rio.MirrorFunc, error) {
	switch packType {
	case string(tartrans.PackType):
		return tartrans.Mirror, nil
	case string(ziptrans.PackType):
		return ziptrans.Mirror, nil
	default:
		return nil, Errorf(rio.ErrUsage, "unsupported packtype %q", packType)
	}
}

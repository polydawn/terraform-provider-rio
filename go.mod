module terraform-provider-rio

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.1
	github.com/polydawn/refmt v0.0.0-20190807091052-3d65705ee9f1 // indirect
	github.com/warpfork/go-errcat v0.0.0-20180917083543-335044ffc86e
	go.polydawn.net/go-timeless-api v0.0.0-00010101000000-000000000000
	go.polydawn.net/rio v0.0.0-00010101000000-000000000000
	golang.org/dl v0.0.0-20200811212135-d149fc5456ff // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1 // indirect
	xi2.org/x/xz v0.0.0-00010101000000-000000000000 // indirect
)

replace go.polydawn.net/rio => github.com/polydawn/rio v0.0.0-20200325050149-e97d9995e350

replace go.polydawn.net/go-timeless-api => github.com/polydawn/go-timeless-api v0.0.0-20190707220600-0ece408663ed

replace xi2.org/x/xz => github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8

// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "834c211602b62487738df0df3a67311a"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"9053b5acb109aa13a67c1d92df9f4853": "1f8b08000000000000ff84d0414b03311005e07b7ec5bbb5450be2b527410f2282143cf438bb19cbc0ec649d2454ffbdd8b9145d684ee1e52333bced1637931c9d1ae37d4ea3f3efadd1a08c5ed96b5a2700908c38831c2bbb90de9e73a389cf79e3af065869b0ae8acc1fd4b561b50a57bb07bde27822d185ffbac967e73033d57a2a9eff9878f4a2b19058c3c2a0fb600329d9c8c0548cbfffb3bb606ffbe7d787fd012f4f07ac256fd26697d265678fe564297b992f3bc34875a4ccbb9f000000ffff6edefa055e010000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("migrations", "./migrations")
		b.SetResolver("001_initial.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "9053b5acb109aa13a67c1d92df9f4853"})
	}()
	return nil
}()
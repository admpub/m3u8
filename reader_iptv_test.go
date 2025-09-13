package m3u8

import (
	"fmt"
	"os"
	"testing"

	"github.com/admpub/pp/ppnocolor"
)

func TestDecodeMasterPlaylistIPTV(t *testing.T) {
	f, err := os.Open("sample-playlists/master-iptv.m3u8")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	m, listType, err := DecodeFrom(f, false)
	if err != nil {
		t.Fatal(err)
	}
	if listType != MEDIA {
		t.Error("Sample not recognized as media playlist.")
	}
	p := m.(*MediaPlaylist)
	// TODO check other values
	/*
		for i, v := range p.Segments {
			if v == nil {
				break
			}
			t.Logf(`%d: %+v`, i, v)
		}
	*/
	fmt.Println(p.Encode().String())
	ppnocolor.Println(p)
}

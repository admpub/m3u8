package m3u8

import (
	"fmt"
	"strings"
)

type TVGParams map[string]string

func (c TVGParams) String() string {
	var encoded []string
	for k, v := range c {
		encoded = append(encoded, k+`=`+fmt.Sprintf(`%q`, v))
	}
	return strings.Join(encoded, ` `)
}

type INFParams map[string]string

func (c INFParams) String() string {
	var encoded []string
	for k, v := range c {
		encoded = append(encoded, k+`=`+fmt.Sprintf(`%q`, v))
	}
	return strings.Join(encoded, ` `)
}

func infParse(line string) (tvg TVGParams, inf INFParams) {
	out := decodeParamsLine(line)
	if len(out) == 0 {
		return
	}
	tvg = TVGParams{}
	for k, v := range out {
		if len(k) < 5 {
			continue
		}
		if strings.HasPrefix(k, `tvg-`) || strings.HasPrefix(k, `x-tvg-`) {
			tvg[k] = v
			delete(out, k)
		}
	}
	if len(tvg) == 0 {
		tvg = nil
	}
	if len(out) > 0 {
		inf = INFParams(out)
	}
	return
}

func MediaSegmentTVG(tvg TVGParams) func(*MediaSegment) {
	return func(seg *MediaSegment) {
		seg.TVG = tvg
	}
}

func MediaSegmentINF(params INFParams) func(*MediaSegment) {
	return func(seg *MediaSegment) {
		seg.Params = params
	}
}

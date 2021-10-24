package m3u8

import (
	"fmt"
	"regexp"
	"strings"
)

var tvgParamRegexp = regexp.MustCompile(`tvg-([a-z]+)="([^"]+)"`)

type TVGParams map[string]string

func (c TVGParams) String() string {
	var encoded []string
	for k, v := range c {
		encoded = append(encoded, `tvg-`+k+`=`+fmt.Sprintf(`%q`, v))
	}
	return strings.Join(encoded, ` `)
}

func tvgParse(line string) TVGParams {
	//tvg-id="CCTVPlus1.cn" tvg-country="CN" tvg-language="Chinese" tvg-logo="" group-title=""
	matches := tvgParamRegexp.FindAllStringSubmatch(line, -1)
	if len(matches) == 0 {
		return nil
	}
	result := TVGParams{}
	for _, vals := range matches {
		key := vals[1]
		val := vals[2]
		result[key] = val
	}
	return result
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
		if k[0:4] == `tvg-` {
			tvg[k[4:]] = v
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

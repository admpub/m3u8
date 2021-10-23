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
	result := TVGParams{}
	//tvg-id="CCTVPlus1.cn" tvg-country="CN" tvg-language="Chinese" tvg-logo="" group-title=""
	matches := tvgParamRegexp.FindAllStringSubmatch(line, -1)
	for _, vals := range matches {
		key := vals[1]
		val := vals[2]
		result[key] = val
	}
	return result
}

func MediaSegmentTVG(tvg TVGParams) func(*MediaSegment) {
	return func(seg *MediaSegment) {
		seg.TVG = tvg
	}
}

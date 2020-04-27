package pindex

import "strings"

func Encode(txt string) []string {
	var inits [][]uint8
	//依次处理各个字符
	ct := 0 //字符类型：1=中文；-1=非中文；0=未知，用于分割不同种类的字符
	for _, c := range txt {
		cs, ok := hanzi[c]
		if ok { //是中文
			if ct != 1 {
				inits = append(inits, []uint8{' '})
			}
			ct = 1
			inits = append(inits, cs)
			continue
		}
		//不是中文
		if ct != -1 {
			inits = append(inits, []uint8{' '})
		}
		ct = -1
		//不是中文的字符只取英文字母和数字，其他一律转化为空格
		if (c >= '0' && c <= 9) || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			inits = append(inits, []uint8{uint8(c)})
		} else {
			inits = append(inits, []uint8{' '})
		}
	}
	//对多音字进行展开
	codes := []string{""}
	for _, cs := range inits {
		var codes2 []string
		for _, s := range codes {
			for _, c := range cs {
				codes2 = append(codes2, s+string(c))
			}
		}
		codes = codes2
	}
	//滤除编码中的多余空格
	var codes2 []string
	for _, code := range codes {
		var ss []string
		for _, c := range strings.Split(code, " ") {
			if len(c) > 0 {
				ss = append(ss, c)
			}
		}
		codes2 = append(codes2, strings.Join(ss, " "))
	}
	return codes2
}

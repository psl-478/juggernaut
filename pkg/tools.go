package pkg

import "regexp"

func ExtractChineseNameUnicode(input string) string {
	// 使用\p{Han}匹配所有汉字字符
	pattern := `\p{Han}+`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(input, -1)

	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

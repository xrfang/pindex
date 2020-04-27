# pindex

Chinese Pinyin Initialism Generator

我需要一个用拼音首字母输入检索汉字的工具，目前能找到的Go的汉字转拼音工具
对我这个目的都太过强（复）大（杂）了。因此，我写了这个超级简单的汉字转换
为拼音首字母的工具。

转换规则如下：

1. 字符串必须为UTF-8编码。
2. 中文和非中文混杂的字符串在中文与非中文之间用空格隔开。中文字符转换为
   其拼音首字母；非中文字符只取英文字母和数字，其他均转换为空格。注意：
   中文符号和全角数字等均视为非中文，按空格处理。
3. 多音字会按照每种发音展开，因此，返回值类型是`[]string`。

用法示例：

    for _, code := range pindex.Encode("不学无术，白术") {
        fmt.Println(code)
    }

上例返回：

    BXMS BS
    BXMS BZ
    BXMZ BS
    BXMZ BZ
    BXWS BS
    BXWS BZ
    BXWZ BS
    BXWZ BZ

因为“无”和“术”是多音字，各自有两种发音，“术”又出现了两次，因此返回了8种
编码组合。
package common

type SupplierSubInfo struct {
	FromWhere string         `json:"from_where"` // 从哪个网站下载来的
	TopN      int64          `json:"top_n"`      // 是 Top 几？
	Name      string         `json:"name"`       // 字幕的名称，这个比较随意，优先是影片的名称，然后才是从网上下载字幕的对应名称
	Language  Language `json:"language"`   // 字幕的语言
	FileUrl   string         `json:"file-url"`   // 字幕文件下载的路径
	Score     int64          `json:"score"`      // TODO 字幕的评分，需要有一个独立的评价体系
	Offset    int64          `json:"offset"`     // 字幕的偏移
	Ext       string         `json:"ext"`        // 字幕文件的后缀名带点，有可能是直接能用的字幕文件，也可能是压缩包
	Data      []byte         `json:"data"`       // 字幕文件的二进制数据
}

func NewSupplierSubInfo(fromWhere string, topN int64, name string, language Language, fileUrl string, score int64, offset int64, ext string, data []byte) *SupplierSubInfo {
	return &SupplierSubInfo{FromWhere: fromWhere, TopN: topN,Name: name, Language: language, FileUrl: fileUrl, Score: score, Offset: offset, Ext: ext, Data: data}
}

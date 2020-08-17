package models

//Index 是这个块在整个链中的位置
//Timestamp 显而易见就是块生成时的时间戳
//Hash 是这个块通过 SHA256 算法生成的散列值
//PrevHash 代表前一个块的 SHA256 散列值
//BPM 每分钟心跳数，也就是心率

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

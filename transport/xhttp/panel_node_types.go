// 依据原样本 RTTI 逐字段生成的 Go 类型，不是作者文件原文。
package xhttp

import ()

// R0J19auWof 对应 0x3729c50，amd64 大小 296 字节。
type R0J19auWof struct {
	TzkkcVTusJD  string            `json:"host,omitempty"`
	HoBz4V       string            `json:"path,omitempty"`
	MfWFUwd1     string            `json:"mode,omitempty"`
	F0XSwqNYoaP  map[string]string `json:"headers,omitempty"`
	N2C4XMihD    *MUVoWl0          `json:"xPaddingBytes,omitempty"`
	A6NxppvA     bool              `json:"noSSEHeader,omitempty"`
	ULnN_iDWLIXN *MUVoWl0          `json:"scMaxEachPostBytes,omitempty"`
	LEaRo0       int64             `json:"scMaxBufferedPosts,omitempty"`
	Iv950ouQ5AW  *MUVoWl0          `json:"scStreamUpServerSecs,omitempty"`
	PnRF9ifkXKD  bool              `json:"xPaddingObfsMode,omitempty"`
	H6PfvFysST   string            `json:"xPaddingKey,omitempty"`
	LjJbxC6vDL   string            `json:"xPaddingHeader,omitempty"`
	V0qykwP      string            `json:"xPaddingPlacement,omitempty"`
	MLlGWkds     string            `json:"xPaddingMethod,omitempty"`
	KJMkBJ       string            `json:"uplinkHTTPMethod,omitempty"`
	XEKp9SbqKQ   string            `json:"sessionPlacement,omitempty"`
	G9oU9diBqr   string            `json:"sessionKey,omitempty"`
	HqtZJaRZNw   string            `json:"seqPlacement,omitempty"`
	NPLvCM       string            `json:"seqKey,omitempty"`
	Aonc19Hp33E  string            `json:"uplinkDataPlacement,omitempty"`
	HFEkUe4a9Sas string            `json:"uplinkDataKey,omitempty"`
	JvaHmm       *MUVoWl0          `json:"uplinkChunkSize,omitempty"`
	Diwr8ey      int32             `json:"serverMaxHeaderBytes,omitempty"`
}

// MUVoWl0 对应 0x364cca0，amd64 大小 8 字节。
type MUVoWl0 struct {
	AS9zG362 int32 `json:"from,omitempty"`
	HHTPwX0N int32 `json:"to,omitempty"`
}

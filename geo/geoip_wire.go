package geo

import (
	"errors"
	"strings"

	"google.golang.org/protobuf/encoding/protowire"
)

// 原 init（0xe99820）创建的五个独立错误对象。
var (
	readBytesErrorAt5096F40     = errors.New("failed to read bytes")
	readLengthErrorAt5096F50    = errors.New("failed to read expected length of bytes")
	invalidFileErrorAt5096F60   = errors.New("invalid geodata file")
	invalidVarintErrorAt5096F70 = errors.New("invalid geodata varint length")
	codeNotFoundErrorAt5096F80  = errors.New("code not found")
)

// kbUmw5F78（0xe99b80）保留原六状态解析器。它只接受消息列表和
// 消息首字段都以 0x0a 开始的布局，不是通用 protobuf 字段扫描器。
func kbUmw5F78(data []byte, code string) ([]byte, error) {
	var buffer [8]byte
	varint := buffer[:0]
	state, offset, size := 1, 0, 1
	var messageLength, countryLength, countryVarintLength int
	readingCountryLength := false
	for offset != len(data) {
		end := offset + size
		if len(data) < end {
			return nil, readLengthErrorAt5096F50
		}
		// 原实现在读取之前切片；负数和溢出引起的边界 panic 不转为 error。
		chunk := data[offset:end]
		switch state {
		case 1, 3:
			if chunk[0] != 0x0a {
				return nil, invalidFileErrorAt5096F60
			}
			state++
			size = 1
		case 2, 4:
			varint = append(varint, chunk...)
			if chunk[0] > 0x7f {
				offset = end
				size = 1
				continue
			}
			value, n := protowire.ConsumeVarint(varint)
			if n < 0 {
				return nil, invalidVarintErrorAt5096F70
			}
			if readingCountryLength {
				countryLength = int(value)
				countryVarintLength = n
				size = int(value)
			} else {
				messageLength = int(value)
				size = 1
			}
			state++
			varint = varint[:0]
			readingCountryLength = !readingCountryLength
		case 5:
			if strings.EqualFold(string(chunk), code) {
				start := end - (countryVarintLength + countryLength) - 1
				if start >= 0 {
					end = start
				}
				size = messageLength
				state = 6
			} else {
				next := end + messageLength - countryLength - countryVarintLength - 1
				if next >= 0 {
					end = next
				}
				size = 1
				state = 1
			}
		case 6:
			return chunk, nil
		}
		offset = end
	}
	// 连续字节恰好用尽时，无论处于哪个状态，原函数都返回未找到。
	return nil, codeNotFoundErrorAt5096F80
}

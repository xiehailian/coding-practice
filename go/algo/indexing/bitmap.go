package indexing

type BitMap []byte

func NewBitMap(length uint) BitMap {
	return make([]byte, length/8+1)
}

func (b BitMap) Set(value uint) {
	byteIndex := value / 8
	if byteIndex >= uint(len(b)) {
		return
	}
	bitIndex := value % 8
	b[byteIndex] |= 1 << bitIndex
}

func (b BitMap) Get(value uint) bool {
	byteIndex := value / 8
	if byteIndex >= uint(len(b)) {
		return false
	}
	bitIndex := value % 8
	return b[byteIndex]&(1<<bitIndex) != 0
}
package main

import "fmt"

////////////Utils//////////////////
func calcFrequency(text []byte) map[byte]int {
	frequencyArray := make(map[byte]int)
	for _, symbol := range text {
		if _, ok := frequencyArray[symbol]; ok {
			frequencyArray[symbol]++
		} else {
			frequencyArray[symbol] = 1
		}
	}
	return frequencyArray
}

/////////////HuffmanNode//////////////
type HuffmanNode struct {
	frequency  int
	leftChild  *HuffmanNode
	rightChild *HuffmanNode
	symbol     byte
	isLeaf     bool
}

func (h *HuffmanNode) IsLeaf() bool {
	return h.isLeaf
}

func (h *HuffmanNode) GetSymbol() byte {
	return h.symbol
}

func (h *HuffmanNode) GetKey() int {
	return h.frequency
}

func newHuffmanNode(left, right *HuffmanNode) *HuffmanNode {
	return &HuffmanNode{
		frequency:  left.frequency + right.frequency,
		leftChild:  left,
		rightChild: right,
		isLeaf:     false,
	}
}

func newHuffmanNodeLeaf(frequency int, symbol byte) *HuffmanNode {
	return &HuffmanNode{
		frequency: frequency,
		symbol:    symbol,
		isLeaf:    true,
	}
}

////////////////Huffman Tree/////////////
type HuffnamTree struct {
	frequency     int
	root          *HuffmanNode
	huffmanCode   map[byte]string
	encodedString string
}

func (ht *HuffnamTree) GetRoot() *HuffmanNode {
	return ht.root
}

func newHuffnamTree(root *HuffmanNode) *HuffnamTree {
	return &HuffnamTree{
		frequency:     root.frequency,
		root:          root,
		huffmanCode:   make(map[byte]string),
		encodedString: "",
	}
}

func (ht *HuffnamTree) PrintHuffmanCode() {
	fmt.Println("------HuffmanCode------")
	for k, v := range ht.huffmanCode {
		fmt.Printf("%c: %s\n", k, v)
	}
}

func (ht *HuffnamTree) encode(root *HuffmanNode, str string) {
	if root == nil {
		return
	}

	if root.IsLeaf() {
		if str == "" {
			ht.huffmanCode[root.symbol] = "1"
		} else {
			ht.huffmanCode[root.symbol] = str
		}
	}
	ht.encode(root.leftChild, str+"0")
	ht.encode(root.rightChild, str+"1")
}

func (ht *HuffnamTree) GetEncodedString() string {
	return ht.encodedString
}

func (ht *HuffnamTree) dencode(root *HuffmanNode,
	index *int, str []byte, text *[]byte) {
	if root == nil {
		return
	}

	if root.IsLeaf() {
		*text = append(*text, root.symbol)
		return
	}
	*index++

	if string(str[*index]) == "0" {
		ht.dencode(root.leftChild, index, str, text)
	} else {
		ht.dencode(root.rightChild, index, str, text)
	}
}

func (ht *HuffnamTree) GetDecodedString(str string) string {
	text := []byte{}
	index := -1

	if ht.root.IsLeaf() {
		for ; ht.root.frequency > 0; ht.root.frequency-- {
			text = append(text, ht.root.symbol)
		}
	} else {
		for index < len(str)-1 {
			ht.dencode(ht.root, &index, []byte(str), &text)
		}
	}

	return string(text)
}

func CreateHuffnamTree(text []byte) *HuffnamTree {
	data := calcFrequency(text)
	heap := NewHeap[*HuffmanNode](len(data))
	for symbol, frequency := range data {
		heap.Insert(newHuffmanNodeLeaf(frequency, symbol))
	}
	for heap.Size() > 1 {
		firstNode, _ := heap.Remove()
		secondNode, _ := heap.Remove()
		heap.Insert(newHuffmanNode(firstNode, secondNode))
	}
	root, _ := heap.Remove()
	hTree := newHuffnamTree(root)
	hTree.encode(hTree.root, "")
	for _, it := range text {
		hTree.encodedString += hTree.huffmanCode[it]
	}
	return hTree
}

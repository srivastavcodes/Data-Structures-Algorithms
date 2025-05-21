package main

import (
	"bytes"
	"strconv"
	"strings"
)

type Codec struct{}

func Constructor() Codec { return Codec{} }

func (cc *Codec) serialize(root *TreeNode) string {
	var buffer bytes.Buffer
	preorderSer(root, &buffer)
	return buffer.String()
}

func preorderSer(node *TreeNode, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString("null,")
		return
	}
	buffer.WriteString(strconv.Itoa(node.Val))
	buffer.WriteString(",")
	preorderSer(node.Left, buffer)
	preorderSer(node.Right, buffer)
}

func (cc *Codec) deserialize(data string) *TreeNode {
	strs := strings.Split(data, ",")
	var index int
	return preorderDeSer(strs, &index)
}

func preorderDeSer(strs []string, i *int) *TreeNode {
	if strs[*i] == "null" {
		*i++
		return nil
	}
	val, _ := strconv.Atoi(strs[*i])
	*i++
	root := &TreeNode{Val: val}

	root.Left = preorderDeSer(strs, i)
	root.Right = preorderDeSer(strs, i)
	return root
}

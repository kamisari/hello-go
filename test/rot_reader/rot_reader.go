package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// TODO:make rot13Reader.Read()
func (rot *rot13Reader) Read(p []byte) (int, error) {
	n := int(0)
	slice := make([]byte, 128, 256)

	// 注意:前に r 忘れて無限再帰してスタックオーバーフローした
	n, err := rot.r.Read(slice)
	if err != nil && err != io.EOF {
		// stderrに吐く関数調べてない
		// fmt.Println(err)
		fmt.Fprintf(os.Stderr, "不明なエラーが発生しました")
		os.Exit(1)
	}


	// DONE:復号化してpに詰める
	rotmap := newRotmap()
	m := int(0)
	for i := range(p) {
		if i < n {
			p[i] = rotmap[slice[i]]
			m++
		} else {
			return m, fmt.Errorf("slice[%v] length caution", i)
		}
	}
	return m, nil
}

// せっかくなのでmap使ってみる
func newRotmap() map[byte]byte {
	newMap := make(map[byte]byte)
	// ASC2
	// TODO:やばい気持ち悪い、何とかしたい
	for i,I, j,J := byte('a'),byte('A'), byte('n'),byte('N'); i <= byte('m'); i,I, j,J = i+1,I+1, j+1,J+1 {
		// 少文字
		newMap[i] = j
		newMap[j] = i

		// 大文字
		newMap[I] = J
		newMap[J] = I
	}
	newMap[byte(' ')] = byte(' ')
	newMap[byte('\n')] = byte('\n')

	return newMap
}

func main() {

	fmt.Println("test")

	// io.Readerをsで受ける
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// io.Reader s を rot13Reader.r に含んだ rot13Reader r を作る
	r := rot13Reader{s}
	// io.Copy で標準出力へ r の内容を吐き出す
	io.Copy(os.Stdout, &r)
}

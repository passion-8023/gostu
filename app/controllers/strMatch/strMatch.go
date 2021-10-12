package strMatch

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gostu/pkg/response"
)

func BruteForce(ctx *gin.Context)  {
	mainStr := "baddef"
	patternStr := "dde"
	response.SuccessResponse(ctx, "获取成功", bfIndexOf(mainStr, patternStr))
	
}

func bfIndexOf(S, T string) (ret int) {
	idxS, idxT := 0, 0
	lenS, lenT := len(S), len(T)
	for {
		if idxS > (lenS - 1) || idxT >= lenT {
			break
		}
		if S[idxS] == T[idxT] {
			idxS++
			idxT++
		} else {
			idxS = idxS - idxT + 1
			idxT = 0
		}
	}

	if idxT > lenT - 1 {
		return idxS - lenT
	}
	return -1
}

func RabinKarp(ctx *gin.Context)  {
	
}

func BoyerMoore(ctx *gin.Context)  {
	mainStr := "baddef"
	patternStr := "cabcab"
	suffix, prefix := generateGs(patternStr)
	fmt.Println(suffix)
	fmt.Println(prefix)
	response.SuccessResponse(ctx, "获取成功", bmIndexOf(mainStr, patternStr))
}

const SIZE = 256

func generateHashTable(T string) (bc []int) {
	bc = make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		bc[i] = -1
	}
	for i := 0; i < len(T); i++ {
		bc[T[i]] = i
	}
	return
}

func generateGs(T string) (suffix []int, prefix []bool) {
	lenT := len(T)
	suffix = make([]int, lenT)
	prefix = make([]bool, lenT)
	for i := 0; i < lenT; i++ {
		suffix[i] = -1
	}
	for i := 0; i < lenT - 1; i++ {
		j := i
		k := 0
		for j >= 0 && T[j] == T[lenT-1-k] {
			j--
			k++
			suffix[k] = j+1
		}
		if j == -1 {
			prefix[k] = true
		}
	}
	return
}

func moveByGS(j int, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	for i := j + 2; i < m - 1; i++ {
		if prefix[m-i] == true {
			return i
		}
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func bmIndexOf(S, T string) int {
	bc := generateHashTable(T)
	suffix, prefix := generateGs(T)
	idxS := 0
	lenS, lenT := len(S), len(T)
	for idxS <= (lenS - lenT) {
		var j int
		for j = lenT - 1; j >= 0; j-- {
			if S[idxS+j] != T[j] {break}
		}
		if j < 0 {
			return idxS
		}
		x := j - bc[S[idxS+j]]
		y := 0
		if j < lenT - 1 {
			y = moveByGS(j, lenT, suffix, prefix)
		}
		idxS = idxS + max(x, y)
	}
	return -1
}

func KnuthMorrisPratt(ctx *gin.Context)  {
	a := "ababaeababc"
	b := "ababacd"
	index := kmpIndexOf(a, b)
	response.SuccessResponse(ctx, "获取成功", index)
}

func kmpIndexOf(S, T string) int {
	next := getNext(T)
	lenS := len(S)
	lenT := len(T)
	j := 0
	for i := 0; i < lenS; i++ {
		for j > 0 && S[i] != T[j] {
			j = next[j-1] + 1
		}
		if S[i] == T[j] {
			j++
		}
		if j == lenT {
			return i - lenT + 1
		}
	}
	return -1
}

func getNext(T string) []int {
	lenT := len(T)
	next := make([]int, lenT)
	next[0] = -1
	k := -1
	for i := 1; i < lenT; i++ {
		for k != -1 && T[k+1] != T[i] {
			k = next[k]
		}
		if T[k+1] == T[i] {
			k++
		}
		next[i] = k
	}
	return next
}


package kata

// 已知：英文字母序列chars
// 未知：英文字母missing
// 条件：
// （1）missing是chars中缺失的字母
// （2）序列中只有一个字母缺失
// （3）序列长度至少为2
// （4）要么全部大写，要么全部小写

type RuneSet map[rune]bool

func newRuneSet() RuneSet {
	return map[rune]bool{}
}

func newRuneSetWithInitials(intials []rune) RuneSet {
	s := map[rune]bool{}
	for _, i := range intials {
		s[i] = true
	}
	return s
}

func (s RuneSet) Add(r rune) {
	s[r] = true
}

func (s RuneSet) Diff(t RuneSet) []rune {
	for r := range t {
		delete(s, r)
	}
	diff := []rune{}
	for r := range s {
		diff = append(diff, r)
	}
	return diff
}

func (s RuneSet) Len() int { return len(s) }

func FindMissingLetter(chars []rune) rune {
	var seriesSet RuneSet = newRuneSet()
	for r := chars[0]; r <= chars[len(chars)-1]; r++ {
		seriesSet.Add(r)
	}

	var charsSet RuneSet = newRuneSetWithInitials(chars)
	diff := seriesSet.Diff(charsSet)
	return diff[0]
}

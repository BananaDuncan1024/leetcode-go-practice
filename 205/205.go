package main

func main() {
	s, t := "egg", "add"

	isIsomorphic(s, t)
}

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]
		if val, ok := sMap[sChar]; ok {
			if val != tChar {
				return false
			}
		} else {
			if _, ok := tMap[tChar]; ok {
				return false
			}
			sMap[sChar] = tChar
			tMap[tChar] = sChar
		}
	}
	return true
}

func isIsomorphic2(s string, t string) bool {
	sMap := make(map[byte]byte, len(s))
	tMap := make(map[byte]byte, len(s))

	for i, _ := range s {
		_, sOK := sMap[s[i]]
		_, tOK := tMap[t[i]]

		if !sOK && !tOK {
			// if char isn't found in both maps
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]

		} else if sOK {
			//if found in sMap but the char isn't same as existing
			if sMap[s[i]] != t[i] {
				return false
			}

		} else if tOK {
			//if found in tMap but the char isn't same as existing
			if tMap[t[i]] != s[i] {
				return false
			}
		}
	}

	return true
}

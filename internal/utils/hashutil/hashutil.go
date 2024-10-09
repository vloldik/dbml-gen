package hashutil

import "hash/fnv"

func FnvSumm(els ...[]byte) uint32 {
	hash := fnv.New32a()
	for _, el := range els {
		hash.Write(el)
	}
	return hash.Sum32()
}

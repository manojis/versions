package main

import "fmt"

func main() {
	version1 := "1.2.3.7"
	version2 := "1.1.7.9"

	if versionCompare(version1, version2) < 0 {
		fmt.Println(version1, "is smaller")
	} else if versionCompare(version1, version2) > 0 {
		fmt.Println(version2, "is smaller")
	} else {
		fmt.Println("Both versions are same")
	}
}

func versionCompare(version1 string, version2 string) int {

	versionLen1, versionLen2 := len(version1), len(version2)

	i, j := 0, 0
	for i < versionLen1 || j < versionLen2 {
		vs1, vs2 := 0, 0

		for i < versionLen1 && '0' <= version1[i] && version1[i] <= '9' {
			vs1 = vs1*10 + int(version1[i]-'0')
			i++
		}

		for j < versionLen2 && '0' <= version2[j] && version2[j] <= '9' {
			vs2 = vs2*10 + int(version2[j]-'0')
			j++
		}
		if vs1 > vs2 {
			return 1
		}
		if vs1 < vs2 {
			return -1
		}
		i, j = i+1, j+1
	}
	return 0
}
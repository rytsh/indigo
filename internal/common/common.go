package common

// Check the errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Version number
var Version string

const info = `
___,___,_______,____
|  :::|///./||'||    \
|  :::|//.//|| || J)  |
|  :::|/.///|!!!|     |
|   _______________   |
|  |:::::::::::::::|  |
|  |_______________|  |
|  |_______________|  |
|  |_______________|  |
||_|\e[1m%s\e[21m||_|
|__|_______________|__|

\e[1mResources\e[21m
%s

\e[1mHome\e[21m
%s
`

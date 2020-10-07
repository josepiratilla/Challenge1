# Challenge1

This is the exercise of the first go challenge as found at <http://golang-challenge.org/go-challenge1/>.

The objective of the challenge is to build the `decoder.go` file (that was missing) to pass the test in `decoder_test.go`. In order to do it, I needed to reverse engineer the files in fixtures, which are binary. The contents of these files allow to obtain all information needed to get the examples present at decoder_test.go which are like the next:
```
Saved with HW Version: 0.808-alpha
Tempo: 120
(0) kick	|x---|x---|x---|x---|
(1) snare	|----|x---|----|x---|
(2) clap	|----|x-x-|----|----|
(3) hh-open	|--x-|--x-|x-x-|--x-|
(4) hh-close	|x---|x---|----|x--x|
(5) cowbell	|----|----|--x-|----|
```
Don't open the decoder.go if you want also to do the challenge. More details at <http://golang-challenge.org/go-challenge1/>
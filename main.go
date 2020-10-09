package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

const groupPerRow = 1
const answerColShift =5

func main()  {
	tm.Clear() // Clear current screen
	fileName := "absgt_01"


	for {
		badGroup := 0
		badChar := 0
		goodGroup := 0
		goodChar := 0

		tm.Clear() // Clear current screen
		right := readFile("lessons/"+fileName+".right")
		answer := readFile("lessons/"+fileName+".answer")

		rightY := 2
		rightX := 1
		rightGroup := 1
		for i := range right {
			tm.MoveCursor(rightX, rightY)
			rightX += len(right[i]) + 1
			if rightGroup% groupPerRow == 0 && rightGroup > 0 || right[i]=="=" {
				rightY++
				rightX = 1
			}
			if len(right[i])>=4 {
				rightGroup++
			}
			if right[i] != answer[i] {
				for j := range right[i] {
					if right[i][j] != answer[i][j] {
						tm.Print( tm.Color(string(right[i][j]), tm.BLUE) )
					} else {
						tm.Print( string(right[i][j]) )
					}
				}
			} else {
				tm.Printf("%s", right[i])
			}
		}

		answerY := 2
		answerX := answerColShift + groupPerRow * 6
		answerGroup := 1
		for i := range answer {
			tm.MoveCursor(answerX, answerY)
			answerX += len(answer[i]) + 1
			if answerGroup% groupPerRow == 0 && answerGroup > 0 || answer[i]=="=" {
				answerY++
				answerX =  answerColShift + groupPerRow * 6
			}
			if len(answer[i])>=4 {
				answerGroup++
			}
			if answer[i] != right[i] {
				badGroup++
				for j := range answer[i] {
					if answer[i][j] != right[i][j] {
						tm.Print( tm.Color(string(answer[i][j]), tm.RED) )
						badChar++
					} else {
						tm.Print( tm.Color(string(answer[i][j]), tm.WHITE) )
						goodChar++
					}
				}
			} else {
				goodGroup++
				goodChar+=len(answer[i])
				tm.Print(answer[i])
			}
		}

		tm.MoveCursor(0,1)
		percentChar := 100-(float32(badChar) / float32(goodChar) * 100)
		percentGroup := 100-(float32(badGroup) / float32(goodGroup) * 100)
		tm.Printf("GOOD: %d, %d --- BAD: %d, %d == %.2f%%, %.2f%%", goodChar, goodGroup-2, badChar, badGroup, percentChar, percentGroup)


		tm.MoveCursor(0,0)
		tm.Flush() // Call it every time at the end of rendering
		time.Sleep(time.Second)
	}
}


func readFile(fileName string) (rows []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.Trim(scanner.Text(), " \n\r\t")
		groupStr := strings.Split( str, " " )
		rows = append(rows, groupStr...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rows
}

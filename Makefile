nijisanji.ics: nijisanji-ics
	./nijisanji-ics > nijisanji.ics

nijisanji-ics: main.go
	go build

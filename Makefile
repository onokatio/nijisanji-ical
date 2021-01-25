main.ics: nijisanji-ics
	./nijisanji-ics > main.ics

nijisanji-ics: main.go
	go build

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	layoutDateTime := "02/01/2006 15:04:05"
	layoutDate := "02/01/2006"

	var t1, t2 time.Time

	// Funzione per input con tentativi e exit
	getDate := func(label string) (time.Time, bool) {
		for attempts := 0; attempts < 3; attempts++ {
			fmt.Printf("%s (gg/mm/aaaa o gg/mm/aaaa hh:mm:ss): ", label)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if strings.EqualFold(input, "exit") {
				return time.Time{}, true
			}
			if input == "" {
				return time.Now().Truncate(24 * time.Hour), false
			}
			t, err := time.Parse(layoutDateTime, input)
			if err == nil {
				return t, false
			}
			t, err = time.Parse(layoutDate, input)
			if err == nil {
				return t, false
			}
			fmt.Println("Formato data non valido. Riprova o digita 'exit' per uscire.")
		}
		return time.Time{}, false
	}

	t1, quit := getDate("Inserisci la prima data")
	if quit {
		fmt.Println("Uscita dal programma.")
		return
	}
	t2, quit = getDate("Inserisci la seconda data")
	if quit {
		fmt.Println("Uscita dal programma.")
		return
	}
	if t1.IsZero() || t2.IsZero() {
		fmt.Println("Numero massimo di tentativi raggiunto. Uscita dal programma.")
		return
	}

	diff := t2.Sub(t1)
	if diff < 0 {
		diff = -diff
	}

	giorni := int(diff.Hours()) / 24
	ore := int(diff.Hours())

	fmt.Printf("Differenza: %d giorni (%d ore)\n", giorni, ore)
}

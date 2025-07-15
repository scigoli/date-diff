# date-diff


Questo programma in Go riceve in input due date e calcola la differenza in giorni e ore, stampando il risultato a terminale.

## Formati accettati
- `gg/mm/aaaa` (es: 15/07/2025)
- `gg/mm/aaaa hh:mm:ss` (es: 15/07/2025 14:30:00)

Se l'orario non viene specificato, viene usato `00:00:00` di default.
Se il campo data viene lasciato vuoto, viene utilizzata la data odierna a mezzanotte.

In caso di errore di inserimento, è possibile riprovare fino a 3 volte oppure uscire digitando `exit` (non sensibile a maiuscole/minuscole).


## Utilizzo
1. Esegui il programma:
   ```sh
   go run main.go
   ```
2. Inserisci le due date richieste in uno dei formati accettati oppure lascia vuoto per usare la data odierna.
3. In caso di errore, puoi riprovare o digitare `exit` per uscire.
4. Verrà stampata la differenza in giorni e ore.

## Requisiti
- Go 1.18 o superiore


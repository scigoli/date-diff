# date-diff



Questo programma in Go riceve in input due date e calcola la differenza in giorni e ore, stampando il risultato a terminale.

## Sicurezza
L'applicazione è progettata per essere sicura rispetto ad attacchi malevoli perché:
- Non esegue mai comandi di sistema o codice arbitrario.
- Non scrive su file e non accetta input da fonti esterne non controllate.
- Limita la lunghezza dell'input a 100 caratteri per evitare buffer overflow.
- Gestisce errori di I/O e chiude in modo sicuro in caso di problemi.
- Utilizza un blocco di recupero (`recover`) per evitare crash improvvisi dovuti a errori imprevisti.

Queste misure rendono l'applicazione robusta contro input malevoli e garantiscono che non possa essere sfruttata per eseguire codice non previsto o compromettere il sistema.

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

## Test
Per eseguire i test automatici:
```sh
go test
```
Questo comando esegue i test unitari definiti nel file `main_test.go` e verifica la correttezza delle funzioni di calcolo della differenza tra date.

## Requisiti
- Go 1.18 o superiore


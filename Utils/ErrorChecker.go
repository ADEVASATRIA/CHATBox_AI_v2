/* Set package program
 * Untuk memberitaukan bahwa ini program dari package Utils. */
package Utils

// Import file DataLogger untuk memanggil function Logger()
import "./DataLogger"

/* Function CheckError
 * Fungsi ini saya buat untuk menampilkan error. 
 *
 * @Params opt int    
 *	-> 1. Menapilkan Error
 *	   2. Menampilkan Message
 *
 * @Params msg string
 *	-> {pesan} / {message}
 *
 * @Params err 'error obj'
 *	-> {err}
 */
func CheckError(opt int, msg string, err error){
	// Jika error tidak kosong (error berisi obj 'error'), maka
	if err != nil {
		if opt == 1 {
			// Panggil function Logger() dengan isian 2 parameter
			DataLogger.Logger(4, err)
		} else if opt == 2 {
			// Panggil function Logger() dengan isian 2 parameter
			DataLogger.Logger(4, msg)
		}
	}
}

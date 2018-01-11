package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/tidwall/gjson"
)

var btcApiPrefix = "https://blockchain.info/rawtx/"

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("\033[0;31mplease input the correct rawtx string hex-encoded\033[0m")
		fmt.Println()
		fmt.Println("usage:")
		fmt.Println("\t./CheckMultiSig rawtx")
		fmt.Println()
		return
	}
	fmt.Println()

	rawTx, err := hex.DecodeString(args[1])
	if err != nil {
		panic(err)
	}
	tx := wire.NewMsgTx(1)
	if err := tx.DeserializeNoWitness(bytes.NewReader(rawTx)); err != nil {
		panic(err)
	}

	pubkeyScript, _ := hex.DecodeString("a914228f554bbf766d6f9cc828de1126e3d35d15e5fe87")
	engine, err := txscript.NewEngine(pubkeyScript, tx, 0, 65503, nil, nil, 0)
	if err != nil {
		fmt.Println("033[0;31malert: Create engine failed...\033[0m" + err.Error())
		return
	}

	//engine.Show()

	err = engine.Execute()
	if err != nil {
		fmt.Println("033[0;31malert: Engine execute...\033[0m" + err.Error())
		return
	}
	fmt.Println()
}

func GetContent(txidHash string) []gjson.Result {
	rawContent, err := http.Get(btcApiPrefix + txidHash)
	if err != nil {
		panic("btc API is down..." + err.Error())
	}

	defer rawContent.Body.Close()
	content, _ := ioutil.ReadAll(rawContent.Body)

	result := gjson.Get(string(content), "inputs.#.prev_out.script").Array()

	return result
}

//txid:33d5732009799f39e666e6dfce8701dac6eaa0321ce82d99a1ca1fae3b3e1ec1

//pk_script:a914228f554bbf766d6f9cc828de1126e3d35d15e5fe87

//origin:010000000116f70d718db1c032e915dfefb25eafeef1cdc46d6e43ab7320890eb9c033e37d00000000fdfe00004830450221009c6c6af600bdb2b918ecb595a0bcae41881eb771da92e4ba4a05ef17249527f402200d105bbb3c462fbb1ab941697a503fb7a9e48bda5a910959e21f11ff213cde5401483045022100d516c3d638076da7b5fdcc8f2fa1914c5caf6548e3cd2efb16a16605b79321230220062991654aed0966b6b3684141583ef0b188043fce867d633abeac653f487337014c695221028bb6ee1127a620219c4f6fb22067536649d439929e177ebfe76386dff52a70842102f9cd8728b12b6c8a17a15cb4a19de000641f78a449c1b619dc271b84643ce0e92103d33aef1ae9ecfcfa0935a8e34bb4a285cfaad1be800fc38f9fc869043c1cbee253aefeffffff01a09b9a62000000001976a914005ee55b3430bc1a882321efcc5cf898a9aeba5988aca9a70700

package users

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"

	// "crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/disintegration/imaging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/fogleman/gg"

	AnCryptoIdentityNFT "github.com/manish-antiersoultions/dynamic_nft_backend/modules/users/contracts"
)

type IPFSPostResponseData struct {
	Name string `json"Name"`
	Hash string `json"Hash"`
	Size string `json"Size"`
}

/**
 * calling smart contract
 */
func getIPFSData(uri string, network string) (string, error) {
	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return "", err
	}

	// Set headers (if required)
	// req.Header.Set("Authorization", "Bearer YOUR_API_TOKEN")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return "", err
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return "", err
	}

	// Print the response
	if obj, ok := data.(map[string]interface{}); ok {
		// name := obj["name"].(string)
		for _, a := range obj["addresses"].([]interface{}) {
			if v, ok := a.(map[string]interface{}); ok {

				// log.Println( a, v[network])
				if v[network] != nil {
					return v[network].(string), nil
				}
			}
		}
	}
	return "", nil
}

/**
 * create png image for username
 */
func creatingTheImage(text string) (string, error) {
	const (
		width  = 800
		height = 1000
	)

	// Create a new image context
	dc := gg.NewContext(width, height)
	ANCRYPTO_BACKGROUND_IMAGE := os.Getenv("ANCRYPTO_BACKGROUND_IMAGE")
	ANCRYPTO_OUTPUT_IMAGE_PATH := os.Getenv("ANCRYPTO_OUTPUT_IMAGE_PATH")
	ANCRYPTO_FONT_FILE := os.Getenv("ANCRYPTO_FONT_FILE")

	// Open the background image
	background, err := gg.LoadImage(ANCRYPTO_BACKGROUND_IMAGE)
	if err != nil {
		return "", err

	}

	// Resize the background image to the desired dimensions
	background = imaging.Resize(background, width, height, imaging.Lanczos)

	// Draw the background image onto the image context
	dc.DrawImage(background, 0, 0)

	// Set the text color and font
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace(ANCRYPTO_FONT_FILE, 42); err != nil {
		return "", err
	}

	// Draw the text on the image
	dc.DrawStringAnchored(text, float64(width)/2, float64(height)/2, 0.5, 0.5)

	if err := dc.LoadFontFace(ANCRYPTO_FONT_FILE, 28); err != nil {
		return "", err
	}
	text2 := "@ancrypto.com"
	// Draw the text on the image
	dc.DrawStringAnchored(text2, float64(width)/2, (float64(height) - 40), 0.5, 0.5)
	//time for file creation
	currentTime := time.Now()
	seconds := currentTime.Unix()
	//convert to string
	secondsString := strconv.FormatInt(seconds, 10)

	fileName := ANCRYPTO_OUTPUT_IMAGE_PATH + "output_" + secondsString + ".png"
	// Save the image to a file
	if err := dc.SavePNG(fileName); err != nil {
		return "", err
	}
	return fileName, nil
}

/**
 * deleting file from server in background
 */
func deleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Println("Error deleting file:", err)
		return
	}

	log.Println("File deleted successfully: ", filePath)
}

/**
 * creating video NFT
 */
func createVideo(text string) (string, error) {

	// Replace these paths with your input video and output video paths
	ANCRYPTO_BACKGROUND_VIDEO := os.Getenv("ANCRYPTO_BACKGROUND_VIDEO")
	ANCRYPTO_OUTPUT_IMAGE_PATH := os.Getenv("ANCRYPTO_OUTPUT_IMAGE_PATH")
	ANCRYPTO_FONT_FILE := os.Getenv("ANCRYPTO_FONT_FILE")

	inputVideo := ANCRYPTO_BACKGROUND_VIDEO

	//time for file creation
	currentTime := time.Now()
	seconds := currentTime.Unix()
	//convert to string
	secondsString := strconv.FormatInt(seconds, 10)
	outputVideo := ANCRYPTO_OUTPUT_IMAGE_PATH + "output_" + secondsString + ".mp4"

	str := fmt.Sprintf("drawtext=text='"+text+"':x=(w-text_w)/2:y=(h-text_h)/2:fontsize=75:fontcolor=white:fontfile=%s", ANCRYPTO_FONT_FILE)

	ffmpegPath := "ffmpeg"
	// Command to add dynamic text overlay
	cmd := exec.Command(ffmpegPath, "-y", "-i", inputVideo, "-vf", str, "-c:a", "copy", outputVideo)

	// Execute the command and check for errors
	// cmd := exec.Command("ls", "-larth", "./")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("Error adding text overlay:", err)
		return "", err
	}

	log.Println("Text overlay added successfully!")
	return outputVideo, nil
}

/**
 * uploading file to IPFS
 */
func uploadFileToIpfs(filePath string) (string, error) {
	INFURA_ADD_URL := os.Getenv("INFURA_ADD_URL")
	INFURA_API_KEY := os.Getenv("INFURA_API_KEY")
	INFURA_API_KEY_SECRET := os.Getenv("INFURA_API_KEY_SECRET")

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	// Create a new multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form field for the file
	fileField, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		log.Println("Error creating form file:", err)
		return "", err
	}

	// Copy the file data to the form field
	_, err = io.Copy(fileField, file)
	if err != nil {
		log.Println("Error copying file data:", err)
		return "", err
	}

	// Close the multipart writer
	writer.Close()

	// Create the HTTP request
	request, err := http.NewRequest("POST", INFURA_ADD_URL, body)
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		return "", err
	}

	// Set the Content-Type header to the multipart form data
	request.Header.Set("Content-Type", writer.FormDataContentType())
	//setting authentication
	request.SetBasicAuth(INFURA_API_KEY, INFURA_API_KEY_SECRET)

	// Make the HTTP request
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return "", err
	}
	defer response.Body.Close()

	// Handle the response
	if response.StatusCode != http.StatusOK {
		log.Println("Error uploading file. Status:", response.StatusCode)
		return "", errors.New("Some error while uploading hash")
	}
	//decode response from ipfs
	var responseData IPFSPostResponseData
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		log.Println("Error parsing JSON response:", err)
		return "", err
	}

	return responseData.Hash, nil
}

/**
 * upload json data to ipfs
 */
func uploadDatatoIPFS(data interface{}) string {
	ANCRYPTO_OUTPUT_IMAGE_PATH := os.Getenv("ANCRYPTO_OUTPUT_IMAGE_PATH")
	// Convert data to JSON
	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println("Failed to convert data to JSON:", err)
		return ""
	}

	//time for file creation
	currentTime := time.Now()
	seconds := currentTime.Unix()
	//convert to string
	secondsString := strconv.FormatInt(seconds, 10)

	fileName := ANCRYPTO_OUTPUT_IMAGE_PATH + "output_" + secondsString + ".json"

	// Write JSON data to a file
	err = ioutil.WriteFile(fileName, jsonBytes, 0777)
	if err != nil {
		log.Println("Failed to write JSON data to file:", err)
		return ""
	}

	hash, err := uploadFileToIpfs(fileName)
	go deleteFile(fileName)
	return hash
}

/**
 * checking username in the smart contract
 */
func checkingUsername(username string) (string, error) {
	POLYGON_RPC := os.Getenv("POLYGON_RPC")

	client, err := ethclient.Dial(POLYGON_RPC)
	if err != nil {
		log.Fatal(err)
	}
	CONTRACT_ADDRESS := os.Getenv("CONTRACT_ADDRESS")
	address := common.HexToAddress(CONTRACT_ADDRESS)

	instance, err := AnCryptoIdentityNFT.NewAnCryptoIdentityNFT(address, client)
	if err != nil {
		log.Println("-----------1-------------", err)
		return "", errors.New("Network is too busy")
	}

	// GetTokenURIFromUsernameTokenId
	uri, err := instance.GetTokenURIFromUsernameTokenId(&bind.CallOpts{}, username)
	if err != nil {
		log.Println("Error loading contract", err)
		return "", err
	}
	return uri, err
}

/**
 * checking username in the smart contract
 */
func checkingAddressAlreadyUsed(address string) (string, error) {
	POLYGON_RPC := os.Getenv("POLYGON_RPC")

	client, err := ethclient.Dial(POLYGON_RPC)
	if err != nil {
		log.Fatal(err)
	}
	CONTRACT_ADDRESS := os.Getenv("CONTRACT_ADDRESS")
	contract_address := common.HexToAddress(CONTRACT_ADDRESS)

	instance, err := AnCryptoIdentityNFT.NewAnCryptoIdentityNFT(contract_address, client)
	if err != nil {
		log.Println("-----------1-------------", err)
		return "", errors.New("Network is too busy")
	}
	user_address := common.HexToAddress(address)
	// GetTokenURIFromUsernameTokenId
	username, err := instance.AddressUsername(&bind.CallOpts{}, user_address)
	if err != nil {
		log.Println("Error loading contract", err)
		return "", err
	}
	return username, err
}

/**
 * convert gwei to wei
 */
func gweiToWei(i uint64) big.Int {
	gweiValue := new(big.Int).SetUint64(i)

	// Convert gwei to wei
	weiValue := new(big.Int).Mul(gweiValue, big.NewInt(1e9))
	// fmt.Printf("type %t", weiValue)

	return *weiValue
}

/**
 * creating Identity NFT with username in the smart contract
 */
func creatingIdentityNFT(to string, uri string, username string) (string, error) {
	POLYGON_RPC := os.Getenv("POLYGON_RPC")
	MANAGER_ADDRESS := os.Getenv("MANAGER_ADDRESS")
	MANAGER_PRIVATE_KEY := os.Getenv("MANAGER_PRIVATE_KEY")

	client, err := ethclient.Dial(POLYGON_RPC)
	if err != nil {
		log.Println("-----------1-------------", err)
		return "", errors.New("Network is too busy")
	}
	CONTRACT_ADDRESS := os.Getenv("CONTRACT_ADDRESS")

	privateKey, err := crypto.HexToECDSA(MANAGER_PRIVATE_KEY)
	if err != nil {
		log.Println("-----------2-------------", err)
		return "", err
	}
	fromAddress := common.HexToAddress(MANAGER_ADDRESS)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println("-----------3-------------", err)
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("-----------4-------------", err)
		return "", err
	}

	toAddress := common.HexToAddress(to)

	var iargs []interface{}
	iargs = append(iargs, toAddress)
	iargs = append(iargs, uri)
	iargs = append(iargs, username)

	contractAbi, err := abi.JSON(strings.NewReader(AnCryptoIdentityNFT.AnCryptoIdentityNFTABI))
	if err != nil {
		log.Println("-----------5-------------", err)
		return "", err

	}
	input, err := contractAbi.Pack("safeMint", iargs...)
	if err != nil {
		log.Println("-----------6-------------", err)
		return "", err
	}
	// contract address
	sendTo := common.HexToAddress(CONTRACT_ADDRESS)

	callMsg := ethereum.CallMsg{
		From:     fromAddress,
		To:       &sendTo,
		Gas:      0,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     input,
	}
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Println("-----------7-------------", err)
		return "", err

	}
	gasLimitBigInt := gweiToWei(gasLimit)

	log.Println("estimate gas", gasLimit, gasPrice, gasLimitBigInt)

	// Get the balance of the manager address
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("manager balance", balance)

	//comparing the user balance
	result := (balance).Cmp(&gasLimitBigInt)

	switch result {
	case -1:
		return "", errors.New("New username creation is stopped please contact support!")
	case 0:
		return "", errors.New("New username creation is stopped please contact support!")
	case 1:
		log.Println("balance is greater than gasLimit")
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("-----------8-------------", err)
		return "", err
	}

	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	// if err != nil {
	// 	log.Println("-----------9-------------", err)
	// 	return "", err
	// }

	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)     // in wei
	// auth.GasLimit = gasLimit // in units
	// auth.GasPrice = gasPrice
	// auth.From = fromAddress

	//smart contract instance
	// instance, err := AnCryptoIdentityNFT.NewAnCryptoIdentityNFT(sendTo, client)
	// if err != nil {
	// 	log.Println("Error loading contract", err)
	// 	return "", err
	// }
	//transactions
	// tx, err := instance.SafeMint(auth, toAddress, uri, username)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	tx := types.NewTransaction(nonce, sendTo, big.NewInt(0), gasLimit, gasPrice, input)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Println("-----------9-------------", err)
		return "", err
	}

	ts := types.Transactions{signedTx}
	rawTxBytes, _ := rlp.EncodeToBytes(ts[0])
	rawTxHex := hex.EncodeToString(rawTxBytes)
	//bycode of the transation
	log.Println("-----------10----------------", rawTxHex) // f86...772

	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Println("-----------11-------------", err)
		return "", err
	}

	log.Printf("tx sent: %s", tx.Hash().Hex())

	return tx.Hash().Hex(), err
}

/**
 * update token uri in the smart contract
 */
func updateIdentityNFT(uri string, username string) (string, error) {
	POLYGON_RPC := os.Getenv("POLYGON_RPC")
	MANAGER_ADDRESS := os.Getenv("MANAGER_ADDRESS")
	MANAGER_PRIVATE_KEY := os.Getenv("MANAGER_PRIVATE_KEY")

	client, err := ethclient.Dial(POLYGON_RPC)
	if err != nil {
		log.Println("-----------1-------------", err)
		return "", errors.New("Network is too busy")
	}
	CONTRACT_ADDRESS := os.Getenv("CONTRACT_ADDRESS")

	privateKey, err := crypto.HexToECDSA(MANAGER_PRIVATE_KEY)
	if err != nil {
		log.Println("-----------2-------------", err)
		return "", err
	}
	fromAddress := common.HexToAddress(MANAGER_ADDRESS)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println("-----------3-------------", err)
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("-----------4-------------", err)
		return "", err
	}

	// toAddress := common.HexToAddress(to)

	var iargs []interface{}
	iargs = append(iargs, username)
	iargs = append(iargs, uri)

	contractAbi, err := abi.JSON(strings.NewReader(AnCryptoIdentityNFT.AnCryptoIdentityNFTABI))
	if err != nil {
		log.Println("-----------5-------------", err)
		return "", err

	}
	input, err := contractAbi.Pack("updateTokenUri", iargs...)
	if err != nil {
		log.Println("-----------6-------------", err)
		return "", err
	}
	// contract address
	sendTo := common.HexToAddress(CONTRACT_ADDRESS)

	callMsg := ethereum.CallMsg{
		From:     fromAddress,
		To:       &sendTo,
		Gas:      0,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     input,
	}
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Println("-----------7-------------", err)
		return "", err

	}
	gasLimitBigInt := gweiToWei(gasLimit)

	log.Println("estimate gas", gasLimit, gasPrice, gasLimitBigInt)

	// Get the balance of the manager address
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("manager balance", balance)

	//comparing the user balance
	result := (balance).Cmp(&gasLimitBigInt)

	switch result {
	case -1:
		return "", errors.New("New username creation is stopped please contact support!")
	case 0:
		return "", errors.New("New username creation is stopped please contact support!")
	case 1:
		log.Println("balance is greater than gasLimit")
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("-----------8-------------", err)
		return "", err
	}

	tx := types.NewTransaction(nonce, sendTo, big.NewInt(0), gasLimit, gasPrice, input)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Println("-----------9-------------", err)
		return "", err
	}

	ts := types.Transactions{signedTx}
	rawTxBytes, _ := rlp.EncodeToBytes(ts[0])
	rawTxHex := hex.EncodeToString(rawTxBytes)
	//bycode of the transation
	log.Println("-----------10----------------", rawTxHex) // f86...772

	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Println("-----------11-------------", err)
		return "", err
	}

	log.Printf("tx sent: %s", tx.Hash().Hex())

	return tx.Hash().Hex(), err
}

/**
 * evm based blockchain verrifying signature
 */
func verifySignature(_address string, _message string, _signature string) error {
	signerAddress := common.HexToAddress(_address)

	// Replace with the signer's Ethereum address
	message := []byte(_message)
	msg := accounts.TextHash(message)

	//decoding hex signature to bytes
	signature := hexutil.MustDecode(_signature)
	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	// log.Println("before bytes compare", message, signature)
	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(msg, signature)
	if sigPublicKeyECDSA == nil {
		err = errors.New("Could not get a public get from the message signature")
		return err
	}

	recoveredAddr := crypto.PubkeyToAddress(*sigPublicKeyECDSA)
	if err != nil {
		log.Println(err, "ecreconver")
		return err
	}
	// Compare the recovered address with the provided address
	// log.Println(signerAddress.Hex(), recoveredAddr.Hex())
	matches := bytes.Equal(signerAddress.Bytes(), recoveredAddr.Bytes())

	if matches {
		return nil
	} else {
		return errors.New("Invalid Signature!")
	}
}

func verifySignatureTron(_address string, _message string, _signature string) error {

	// Message that was signed
	message := []byte(_message)

	// decoded, err := base58.Decode(_address)
	// if err != nil {
	// 	return  err
	// }

	// Convert the public key and signature from hex to bytes
	// TRON address to verify against
	pubKeyBytes, version, err := base58.CheckDecode(_address)
	if err != nil {
		log.Println("Error decoding public key hex:", err)
		return err
	}
	// Remove the version byte
	// pubKeyBytes = pubKeyBytes[1:]

	// Convert to hexadecimal
	p := hex.EncodeToString(pubKeyBytes)
	log.Println("before signature address", p, " version: ", version)
	signatureBytes, err := hex.DecodeString(_signature)
	if err != nil {
		log.Println("Error decoding signature hex:", err)
		return err
	}

	pubKeyCurve := elliptic.P256() // Adjust the curve based on your key
	pubKey := &ecdsa.PublicKey{Curve: pubKeyCurve}
	pubKey.X, pubKey.Y = big.NewInt(0).SetBytes(pubKeyBytes[:32]), big.NewInt(0).SetBytes(pubKeyBytes[32:])

	// Verify the signature
	isValid := ecdsa.VerifyASN1(pubKey, message, signatureBytes)

	if isValid {
		return nil
	} else {
		return errors.New("Invalid Signature!")
	}

}

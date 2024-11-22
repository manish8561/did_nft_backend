package users

import (
	"errors"
	"log"
	"os"

	"github.com/manish-antiersoultions/dynamic_nft_backend/utils/helpers"
	"github.com/skip2/go-qrcode"
)

/**
 * creating nft image, posting the data to ipfs and create token uri smart contract (or blockchain)
 */
func postUserDecenterlized(user *UserDecentrializeValidator, ip string) (string, error) {
	user.Username = helpers.TrimString(user.Username)
	user.Address = helpers.TrimString(user.Address)
	user.Data = helpers.TrimString(user.Data)
	user.Signature = helpers.TrimString(user.Signature)

	//getting data from smart contract
	uri, err := checkingUsername(user.Username)
	if err != nil {
		//for username not exist
		if err.Error() != "execution reverted: ERC721: address zero is not a valid owner" {
			return "", err
		}
	}
	if uri != "" {
		return "", errors.New("Username already exists!")
	}
	//checking address if already used
	old_username, err := checkingAddressAlreadyUsed(user.Address)
	if err != nil {
		return "", err
	}
	if old_username != "" {
		return "", errors.New("Address already used for username!")
	}
	//creating Identity NFT Image
	fileName, err := creatingTheImage(user.Username)
	// fileName, err := createVideo(user.Username)
	// return "test", nil
	if err != nil {
		return "", err
	}
	IPFS_GET_URL := os.Getenv("IPFS_GET_URL")
	// uploading image to ipfs
	hash, err := uploadFileToIpfs(fileName)
	if hash == "" {
		return "", errors.New("Error while uploading image to IPFS")
	}
	media_url := IPFS_GET_URL + hash
	log.Println("response image: ", media_url)
	// deleting file from server
	go deleteFile(fileName)

	//creating encrypted data json for ancrypto_url
	//JSON Data
	data := map[string]interface{}{
		"data": user.Data,
	}
	hash = uploadDatatoIPFS(&data)
	if hash == "" {
		return "", errors.New("Error while uploading encryption to IPFS")
	}
	ancrypto_url := IPFS_GET_URL + hash
	log.Println("response encrypted: ", ancrypto_url)
	/**
	 * example of the meta data
	 * {
		"description": "Friendly OpenSea Creature that enjoys long swims in the ocean.",
		"external_url": "https://openseacreatures.io/3",
		"image": "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/3.png",
		"name": "Dave Starbelly",
		"attributes": [ ... ]
		}
	*/
	//JSON Metadata
	data = map[string]interface{}{
		"name":          user.Username,
		"description":   "An unique NFT of the user: " + user.Username + " in the decentralize world with Ancrypto",
		"image":         media_url,
		"ancrypto_url":  ancrypto_url,
		"animation_url": media_url,
	}
	hash = uploadDatatoIPFS(&data)
	if hash == "" {
		return "", errors.New("Error while uploading json to IPFS")
	}
	metadata_url := IPFS_GET_URL + hash
	log.Println("response encrypted: ", metadata_url)

	//writing to smart contract
	res, err := creatingIdentityNFT(user.Address, metadata_url, user.Username)
	if err != nil {
		return "", err
	}
	//insert in the table the success full entries
	//create user entry in background
	go createUser(&User{
		Username:        user.Username,
		Address:         user.Address,
		IpAddress:       ip,
		IpfsUrl:         metadata_url,
		Signature:       user.Signature,
		TransactionHash: res,
	})
	return res, nil
}

/**
 * creating nft image, posting the data to ipfs and update token uri smart contract (or blockchain)
 */
func updateUserDecenterlized(user *UserDecentrializeValidator, ip string) (string, error) {
	user.Username = helpers.TrimString(user.Username)
	user.Address = helpers.TrimString(user.Address)
	user.Data = helpers.TrimString(user.Data)
	user.Signature = helpers.TrimString(user.Signature)

	//getting data from smart contract
	uri, err := checkingUsername(user.Username)
	if err != nil {
		//for username not exist
		if err.Error() != "execution reverted: ERC721: address zero is not a valid owner" {
			return "", err
		}
	}
	if uri == "" {
		return "", errors.New("Username does not exists!")
	}

	//creating Identity NFT Image
	fileName, err := creatingTheImage(user.Username)
	// fileName, err := createVideo(user.Username)

	// return "test", nil
	if err != nil {
		return "", err
	}
	IPFS_GET_URL := os.Getenv("IPFS_GET_URL")
	// uploading image to ipfs
	hash, err := uploadFileToIpfs(fileName)
	if hash == "" {
		return "", errors.New("Error while uploading image to IPFS")
	}
	media_url := IPFS_GET_URL + hash
	log.Println("response image: ", media_url)
	// deleting file from server
	go deleteFile(fileName)

	//creating encrypted data json for ancrypto_url
	//JSON Data
	data := map[string]interface{}{
		"data": user.Data,
	}
	hash = uploadDatatoIPFS(&data)
	if hash == "" {
		return "", errors.New("Error while uploading encryption to IPFS")
	}
	ancrypto_url := IPFS_GET_URL + hash
	log.Println("response encrypted: ", ancrypto_url)
	//JSON Metadata
	data = map[string]interface{}{
		"name":         user.Username,
		"description":  "An unique NFT of the user: " + user.Username + " in the decentralize world with Ancrypto",
		"image":        media_url,
		"ancrypto_url": ancrypto_url,
		// "animation_url": media_url,
	}
	hash = uploadDatatoIPFS(&data)
	if hash == "" {
		return "", errors.New("Error while uploading json to IPFS")
	}
	metadata_url := IPFS_GET_URL + hash
	log.Println("response encrypted: ", metadata_url)
	// return metadata_url, nil
	//updating to the smart contract
	res, err := updateIdentityNFT(metadata_url, user.Username)
	if err != nil {
		return "", err
	}
	//updateing in the table the success full entry in background
	go updateUser(user, ip, metadata_url, res)

	return res, nil
}

/**
 * update data of the user in the table
 */
func updateUser(user *UserDecentrializeValidator, ip string, uri string, transactionHash string) {
	var retrievedUser User
	result := helpers.DB.Where(&User{Username: user.Username}).First(&retrievedUser)
	if result.Error != nil {
		log.Println("----DB error in middleware---------", result.Error)

	}
	retrievedUser.IpfsUrl = uri
	retrievedUser.Signature = user.Signature
	retrievedUser.IpAddress = ip
	retrievedUser.TransactionHash = transactionHash
	helpers.DB.Save(&retrievedUser)
}

/**
 * get user signature from database
 */
func getUserSignature(username string) (string, error) {
	var retrievedUser User
	username = helpers.TrimString(username)
	result := helpers.DB.First(&retrievedUser, "Binary username = ?", username)
	if result.Error != nil {
		// Handle the error
		return "", result.Error
	}
	return retrievedUser.Signature, nil
}

/**
 * get user details from contract
 */
func getUserDetailsFromContract(user *UserValidator) (string, error) {
	user.Username = helpers.TrimString(user.Username)
	user.Network = helpers.TrimString(user.Network)
	// GetTokenURIFromUsernameTokenId
	uri, err := checkingUsername(user.Username)
	if err != nil {
		log.Println("Error for username", err)
		return "", err
	}

	v, err := getIPFSData(uri, user.Network)
	if err != nil {
		log.Println("Error from IPFS", err)
		return "", err
	}
	return v, nil
}

/**
 * genrate qr code image from text
 */
func getQRCodeImage(text string) ([]byte, error) {
	text = helpers.TrimString(text)

	png, err := qrcode.Encode(text, qrcode.Medium, 256)
	return png, err
}

/**
 * verifying address data as per each blockchain
 */
func verifyAddressData(addressData *AddressesDataValidator) (string, error) {
	addressData.Username = helpers.TrimString(addressData.Username)

	// name := obj["name"].(string)
	for _, a := range addressData.Addresses {
		// log.Println("before:\n", a)
		if value, ok := a.(map[string]interface{}); ok {
			for k, v := range value {
				if k == "sign" {
					continue
				}
				switch k {
				case "eth", "bsc", "polygon", "nrk", "ftm":
					//verifying the signature in evm
					err := verifySignature(value[k].(string), addressData.Username, value["sign"].(string))
					if err != nil {
						return "Invalid Signature in " + k, err
					}
				case "tron":
					log.Println("Tron")
					err := verifySignatureTron(value[k].(string), addressData.Username, value["sign"].(string))
					if err != nil {
						return "Invalid Signature in " + k, err
					}
				case "bch":
					log.Println("bch")
				case "doge":
					log.Println("doge")
				case "ltc":
					log.Println("ltc")
				case "btc":
					log.Println("btc")
				case "xtz":
					log.Println("xtz")
				default:
					log.Println("Not found blockchian")
				}
				log.Println("in condition", k, " : ", v)
			}
		}
	}

	return "v", nil
}

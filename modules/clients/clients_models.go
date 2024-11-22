package clients

import "github.com/manish-antiersoultions/dynamic_nft_backend/utils/helpers"

/**
 * add client
 */
func add(client *ClientValidator) (uint, error) {
	client.Email = helpers.TrimString(client.Email)
	//generating random string for api key
	api_key, err := helpers.GenerateRandomString(16)
	if err != nil {
		return 0, err
	}
	api_key = "sandbox_" + api_key

	return createClient(&Client{
		Email:             client.Email,
		Is_Email_Verified: true,
		API_KEY:           api_key,
	})
}

/**
 * get user signature from database
 */
func getClientSignature(username string) (string, error) {
	var retrievedClient Client
	result := helpers.DB.First(&retrievedClient, "username = ?", username)
	if result.Error != nil {
		// Handle the error
		return "", result.Error
	}
	return "", nil
}

/**
 * status update by admin for client
 */
func statusUpdate(client *ClientStatusValidator) (string, error) {
	var retrievedClient Client
	result := helpers.DB.First(&retrievedClient, client.ID)
	if result.Error != nil {
		// Handle the error
		return "", result.Error
	}

	retrievedClient.Status = client.Status
	result = helpers.DB.Save(&retrievedClient)
	if result.Error != nil {
		// Handle the error
		return "", result.Error
	}
	return "success", nil
}

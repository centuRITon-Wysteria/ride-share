package services

type NewNodePost struct {
	IpAddress string `json:"ip_address"`
}
type NewNodeResponse struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type AnnounceTravelPost struct {
	FromLat    string `json:"from_lat"`
	FromLon    string `json:"from_lon"`
	ToLat      string `json:"to_lat"`
	ToLon      string `json:"to_lon"`
	Time       string `json:"time"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type SyncPost struct {
	BlockchainData string `json:"blockchain_data"`
}

type PublicInfoPost struct {
	IpAddress  string `json:"ip_address"`
	Name       string `json:"name"`
	Driver     bool   `json:"driver"`
	License    string `json:"license"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Hobbies    string `json:"hobbies"`
	Skills     string `json:"skills"`
	Interests  string `json:"interests"`
	Others     string `json:"others"`
}

type Message struct {
	SendMessage     bool   `json:"send_message"`
	ReceiveMessage  bool   `json:"receive_message"`
	PrivateKey      string `json:"private_key"`
	PublicAddress   string `json:"public_key"`
	ToPublicAddress string `json:"to_public_key"`
	Message         string `json:"message"`
}

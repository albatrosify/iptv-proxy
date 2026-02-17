package serde

// AuthInfo is the internal representation used for marshaling/unmarshaling
type AuthInfo struct {
	UserInfo   UserInfo   `json:"user_info"`
	ServerInfo ServerInfo `json:"server_info"`
}

// UserInfo is the internal representation used for marshaling/unmarshaling
type UserInfo struct {
	ActiveConnections    IntegerAsInteger  `json:"active_cons"`
	AllowedOutputFormats []string          `json:"allowed_output_formats"`
	CreatedAt            UnixTimeAsInteger `json:"created_at"`
	ExpiresAt            UnixTimeAsInteger `json:"exp_date"`
	IsAuthorized         BooleanAsInteger  `json:"auth"`
	IsTrial              BooleanAsInteger  `json:"is_trial"`
	MaxConnections       IntegerAsInteger  `json:"max_connections"`
	Message              string            `json:"message"`
	Password             string            `json:"password"`
	Status               string            `json:"status"`
	Username             string            `json:"username"`
}

// ServerInfo is the internal representation used for marshaling/unmarshaling
type ServerInfo struct {
	HTTPPort       IntegerAsInteger  `json:"port"`
	HTTPSPort      IntegerAsInteger  `json:"https_port"`
	Process        bool              `json:"process"`
	RTMPPort       IntegerAsInteger  `json:"rtmp_port"`
	ServerProtocol string            `json:"server_protocol"`
	TimeNow        DateTimeAsString  `json:"time_now"`
	TimestampNow   UnixTimeAsInteger `json:"timestamp_now"`
	Timezone       string            `json:"timezone"`
	URL            string            `json:"url"`
}

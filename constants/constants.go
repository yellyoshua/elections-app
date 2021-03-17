package constants

const (
	// CollectionUsers users database collection name
	CollectionUsers string = "users"
	// CollectionSessions sessions database collection name
	CollectionSessions string = "sessions"
	// CollectionProfiles profiles database collection name
	CollectionProfiles string = "profiles"
)

const (
	// DefaultPort
	DefaultPort string = "3000"
	// DefaultDatabase in case MONGODB_DATABASE not exist
	DefaultDatabase string = "elections"
)

const (
	// BearerTokenTemplate _
	BearerTokenTemplate string = "Bearer "
)

const (
	// APIPublicFolder path for serve static files
	APIPublicFolder string = "public"
	// APIUploadFolder path for serve static files
	APIUploadFolder string = "public/uploads"
)

const (
	// Unauthorized when make a request that required authentication
	Unauthorized string = "Unauthorized"
)

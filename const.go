package sign

// HTTP Header keys
const (
	HTTPHeaderAccept      = "Accept"
	HTTPHeaderContentMD5  = "Content-MD5"
	HTTPHeaderContentType = "Content-Type"
	HTTPHeaderUserAgent   = "User-Agent"
	HTTPHeaderDate        = "Date"
	HTTPHeaderAuthorization = "Authorization"
)

// HTTP Header keys used for aliyun signature
const (
	HTTPHeaderCAPrefix           = "X-Ca-"
	HTTPHeaderCASignature        = "X-Ca-Signature"
	HTTPHeaderCATimestamp        = "X-Ca-Timestamp"
	HTTPHeaderCANonce            = "X-Ca-Nonce"
	HTTPHeaderCAKey              = "X-Ca-Key"
	HTTPHeaderCASignatureHeaders = "X-Ca-Signature-Headers"
)

const (
	HTTPHeaderCMSPrefix           = "X-Cms-"
	HTTPHeaderCMSSignature        = "X-Cms-Signature"
	HTTPHeaderCMSTimestamp        = "X-Cms-Timestamp"
	HTTPHeaderCMSNonce            = "X-Cms-Nonce"
	HTTPHeaderCMSKey              = "X-Cms-Key"
	HTTPHeaderCMSIP              = "X-Cms-IP"
	HTTPHeaderCMSSignatureHeaders = "X-Cms-Signature-Headers"
	HTTPHeaderCMSApiVsrsion = "X-Cms-Api-Version"
)


const (
	defaultUserAgent = "Go-Aliyun-Client"
)

package validations

const (
	STR         = "Case-sensitive string"
	IP          = "IP"
	EMAIL       = "Email"
	FQDN        = "FQDN"
	INTEGER     = "Integer"
	CIDR        = "CIDR"
	CITY        = "City"
	COUNTRY     = "Country"
	FLOAT       = "Float"
	URL         = "URL"
	MD5         = "MD5"
	HEXADECIMAL = "Hexadecimal"
	BASE64      = "BASE64"
	DATE        = "Date"
	MAC         = "MAC"
	MIME        = "MIME type"
	PHONE       = "Phone"
	SHA1        = "SHA-1"
	SHA224      = "SHA-224"
	SHA256      = "SHA-256"
	SHA384      = "SHA-384"
	SHA512      = "SHA-512"
	SHA3_224    = "SHA3-224"
	SHA3_256    = "SHA3-256"
	SHA3_384    = "SHA3-384"
	SHA3_512    = "SHA3-512"
	SHA512_224  = "SHA512-224"
	SHA512_256  = "SHA512-256"
	DATETIME    = "Datetime"
	UUID        = "UUID"
	BOOLEAN     = "Boolean"
	ISTR        = "String"
	PATH        = "Path"
	OBJECT      = "UUID|MD5|SHA3-256"
	ADVERSARY   = "Adversary"
)

type Definition struct {
	Type         string       `json:"type" example:"object"`
	Description  string       `json:"description" example:"Important description about the type"`
	DataType     string       `json:"dataType" example:"String"`
	Example      *Entity      `json:"example,omitempty"`
	Attributes   []Definition `json:"attributes,omitempty"`
	Associations []Definition `json:"associations,omitempty"`
	Tags         []string     `json:"tags,omitempty"`
	Correlate    []string     `json:"correlate,omitempty"`
}

var file = Definition{
	Type:         "file",
	Description:  "Object identifying a file, the value can be a UUID or a SHA3-256 or MD5 checksum",
	DataType:     OBJECT,
	Attributes:   []Definition{fileData, hashSHA1, hashMD5, hashSHA256, hashSHA3256},
	Associations: []Definition{filename, filenamePattern},
	Tags:         []string{"malware", "common-file", "system-file"},
	Correlate:    []string{"md5", "sha1", "sha256", "sha3-256", "file-data"},
	Example:      &eFile,
}

var payload = Definition{
	Type:        "payload",
	Description: "SHA3-256 of a message sent in a network packet",
	DataType:    SHA3_256,
	Attributes:  []Definition{hashSHA1, hashMD5, hashSHA256, hashSHA3256},
}

var fileData = Definition{
	Type:        "file-data",
	Description: "File or attachment URL",
	DataType:    URL,
	Attributes:  []Definition{},
}

var adversary = Definition{
	Type:        "adversary",
	Description: "Object identifying a threat actor",
	DataType:    ADVERSARY,
}

var aso = Definition{
	Type:        "aso",
	Description: "Autonomous System Organization",
	DataType:    ISTR,
}

var asn = Definition{
	Type:        "asn",
	Description: "Autonomous System Organization Number",
	DataType:    INTEGER,
}

var malware = Definition{
	Type:        "malware",
	Description: "Malware",
	DataType:    ISTR,
	Attributes: []Definition{
		malwareFamily,
		malwareType,
	},
	Correlate: []string{"malware-family", "malware-type"},
	Example:   &eMalware,
}

var object = Definition{
	Type:        "object",
	Description: "Generic entity composed of other entities, the value can be a UUID or a SHA3-256 or MD5 checksum",
	DataType:    OBJECT,
	Attributes: []Definition{
		descriptor,
	},
}

var descriptor = Definition{
	Type:        "descriptor",
	DataType:    ISTR,
	Description: "The object descriptor",
}

var abaRtn = Definition{
	Type:        "aba-rtn",
	Description: "ABA routing transit number",
	DataType:    INTEGER,
}

var latitude = Definition{
	Type:        "latitude",
	Description: "GPS latitude",
	DataType:    FLOAT,
}

var longitude = Definition{
	Type:        "longitude",
	Description: "GPS longitude",
	DataType:    FLOAT,
}

var country = Definition{
	Type:        "country",
	Description: "Country name",
	DataType:    COUNTRY,
}

var cookie = Definition{
	Type:         "cookie",
	Description:  "HTTP cookie as often stored on the user web client. This can include authentication cookie or session cookie",
	DataType:     STR,
	Associations: []Definition{value},
}

var text = Definition{
	Type:        "text",
	Description: "Any case insensitive text value",
	DataType:    ISTR,
}

var value = Definition{
	Type:        "value",
	Description: "Any case sensitive text value",
	DataType:    STR,
}

var password = Definition{
	Type:        "password",
	Description: "Password",
	DataType:    STR,
}

var airport = Definition{
	Type:        "airport-name",
	Description: "The airport name",
	DataType:    ISTR,
	Attributes:  []Definition{country, city},
}

var profilePhoto = Definition{
	Type:        "profile-photo",
	Description: "Profile photo URL",
	DataType:    URL,
}

var authentiHash = Definition{
	Type:        "authentihash",
	Description: "Authenticode executable signature hash",
	DataType:    HEXADECIMAL,
}

var bankAccountNr = Definition{
	Type:        "bank-account-nr",
	Description: "Bank account number without any routing number",
	DataType:    INTEGER,
	Attributes:  []Definition{bic, bin},
}

var bic = Definition{
	Type:        "bic",
	Description: "Bank Identifier Code Number also known as SWIFT-BIC, SWIFT code or ISO 9362 code",
	DataType:    ISTR,
}

var bin = Definition{
	Type:        "bin",
	Description: "Bank Identification Number",
	DataType:    INTEGER,
}

var btc = Definition{
	Type:        "btc",
	Description: "Bitcoin Address",
	DataType:    STR,
}

var ccNumber = Definition{
	Type:        "cc-number",
	Description: "Credit-Card Number",
	DataType:    INTEGER,
	Attributes:  []Definition{issuer},
}

var issuer = Definition{
	Type:        "issuer",
	Description: "Issuer name",
	DataType:    ISTR,
}

var cdHash = Definition{
	Type:        "cdhash",
	Description: "An Apple Code Directory Hash, identifying a code-signed Mach-O executable file",
	DataType:    HEXADECIMAL,
}

var certificateFingerprint = Definition{
	Type:        "certificate-fingerprint",
	Description: "The fingerprint of a SSL/TLS certificate",
	DataType:    HEXADECIMAL,
}

var chromeExtension = Definition{
	Type:        "chrome-extension-id",
	Description: "Chrome extension ID",
	DataType:    STR,
}

var subnet = Definition{
	Type:        "cidr",
	Description: "A public network segment",
	DataType:    CIDR,
	Attributes:  []Definition{country, city, latitude, longitude, asn, aso},
}

var cpe = Definition{
	Type:        "cpe",
	Description: "Common Platform Enumeration. Structured naming scheme for information technology systems, software, and packages",
	DataType:    ISTR,
}

var cve = Definition{
	Type:        "cve",
	Description: "",
	DataType:    ISTR,
}

var dash = Definition{
	Type:        "dash",
	Description: "Dash address",
	DataType:    STR,
}

var dkim = Definition{
	Type:        "dkim",
	Description: "DKIM public key",
	DataType:    STR,
}

var dkimSignature = Definition{
	Type:        "dkim-signature",
	Description: "DKIM signature",
	DataType:    STR,
}

var domain = Definition{
	Type:        "domain",
	Description: "Internet domain",
	DataType:    FQDN,
	Attributes:  []Definition{whoIsRegistrant, whoIsRegistrar},
}

var email = Definition{
	Type:         "email",
	Description:  "Email Message ID",
	DataType:     STR,
	Attributes:   []Definition{emailBody, emailDisplayName, emailHeader, emailAddress, emailSubject},
	Associations: []Definition{file},
}

var city = Definition{
	Type:        "city",
	Description: "City name",
	DataType:    CITY,
}

var issuingCountry = Definition{
	Type:        "issuing-country",
	Description: "Issuing country name",
	DataType:    COUNTRY,
}

var emailAddress = Definition{
	Type:        "email-address",
	Description: "Sender email address",
	DataType:    EMAIL,
}

var emailBody = Definition{
	Type:        "email-body",
	Description: "Email body",
	DataType:    ISTR,
}

var emailDisplayName = Definition{
	Type:        "email-display-name",
	Description: "Sender display name",
	DataType:    ISTR,
}

var emailHeader = Definition{
	Type:        "email-header",
	Description: "Email header (all headers)",
	DataType:    STR,
}

var emailMimeBoundary = Definition{
	Type:        "email-mime-boundary",
	Description: "MIME boundaries are strings of 7-bit US-ASCII text that define the boundaries between message parts in a MIME message. MIME boundaries are declared in a Content-Type message header for any message that encapsulates more than one message part and in part headers for those parts that encapsulate nested parts.",
	DataType:    STR,
}

var emailSubject = Definition{
	Type:        "email-subject",
	Description: "The subject of the email",
	DataType:    ISTR,
}

var emailThreadIndex = Definition{
	Type:        "email-thread-index",
	Description: "The email thread index",
	DataType:    BASE64,
}

var emailXMailer = Definition{
	Type:        "email-x-mailer",
	Description: "Email x-mailer header",
	DataType:    ISTR,
}

var eppn = Definition{
	Type:        "eppn",
	Description: "The NetId of the person for the purposes of inter-institutional authentication. Should be stored in the form of user@univ.edu, where univ.edu is the name of the local security domain",
	DataType:    EMAIL,
}

var facebookProfile = Definition{
	Type:        "facebook-profile",
	Description: "Facebook profile",
	DataType:    URL,
}

var ffn = Definition{
	Type:        "ffn",
	Description: "The frequent flyer number of a passanger",
	DataType:    STR,
}

var filename = Definition{
	Type:        "filename",
	Description: "A filename or email attachment name",
	DataType:    ISTR,
}

var sizeInBytes = Definition{
	Type:        "size-in-bytes",
	Description: "The size in bytes of an element",
	DataType:    FLOAT,
}

var filenamePattern = Definition{
	Type:        "filename-pattern",
	Description: "A pattern in the name of a file",
	DataType:    STR,
}

var flight = Definition{
	Type:        "flight",
	Description: "A flight number",
	DataType:    STR,
}

var gene = Definition{
	Type:        "gene",
	Description: "Go Evtx sigNature Engine",
	DataType:    STR,
}

var gitHubOrganization = Definition{
	Type:        "github-organization",
	Description: "Github organization",
	DataType:    URL,
}

var gitHubRepository = Definition{
	Type:        "github-repository",
	Description: "Github repository",
	DataType:    URL,
}

var gitHubUser = Definition{
	Type:        "github-user",
	Description: "Github user",
	DataType:    URL,
}

var link = Definition{
	Type:        "link",
	Description: "External link for reference",
	DataType:    URL,
}

var datetime = Definition{
	Type:        "datetime",
	Description: "Time with nanoseconds in the format 2006-01-02T15:04:05.999999999Z07:00",
	DataType:    DATETIME,
}

var lastAnalysis = Definition{
	Type:        "last-analysis",
	Description: "Time of last analysis. Format 2006-01-02T15:04:05.999999999Z",
	DataType:    DATETIME,
}

var date = Definition{
	Type:        "date",
	Description: "Date in format 2006-01-02",
	DataType:    DATE,
}

var dateOfIssue = Definition{
	Type:        "date-of-issue",
	Description: "Date in format 2006-01-02",
	DataType:    DATE,
}

var expirationDate = Definition{
	Type:        "expiration-date",
	Description: "Date in format 2006-01-02",
	DataType:    DATE,
}

var malwareSample = Definition{
	Type:        "malware-sample",
	Description: "Malware Sample URL",
	DataType:    URL,
	Attributes:  []Definition{malware, file},
}

var group = Definition{
	Type:        "group",
	Description: "Adversaries group",
	DataType:    ISTR,
}

var haSSHMD5 = Definition{
	Type:        "hassh-md5",
	Description: "Network fingerprinting standard which can be used to identify specific Client SSH implementations. The fingerprints can be easily stored, searched and shared in the form of an MD5 fingerprint",
	DataType:    MD5,
}

var haSSHServerMD5 = Definition{
	Type:        "hasshserver-md5",
	Description: "Network fingerprinting standard which can be used to identify specific Server SSH implementations. The fingerprints can be easily stored, searched and shared in the form of an MD5 fingerprint",
	DataType:    MD5,
}

var hexa = Definition{
	Type:        "hex",
	Description: "A value in hexadecimal",
	DataType:    HEXADECIMAL,
}

var base64d = Definition{
	Type:        "base64",
	Description: "A value in BASE64 format",
	DataType:    BASE64,
}

var hostname = Definition{
	Type:        "hostname",
	Description: "A full host/dnsname of an attacker",
	DataType:    FQDN,
}

var iban = Definition{
	Type:        "iban",
	Description: "International Bank Account Number",
	DataType:    ISTR,
}

var idNumber = Definition{
	Type:        "id-number",
	Description: "It can be an ID card, residence permit, etc.",
	DataType:    STR,
	Attributes:  []Definition{issuer, dateOfIssue, expirationDate},
}

var ipAddr = Definition{
	Type:        "ip",
	Description: "IP Address",
	DataType:    IP,
	Attributes:  []Definition{subnet},
}

var ja3Fingerprint = Definition{
	Type:        "ja3-fingerprint-md5",
	Description: "JA3 is a method for creating SSL/TLS client fingerprints that should be easy to produce on any platform and can be easily shared for threat intelligence",
	DataType:    MD5,
}

var jabberID = Definition{
	Type:        "jabber-id",
	Description: "Jabber ID",
	DataType:    EMAIL,
}

var jarmFingerprint = Definition{
	Type:        "jarm-fingerprint",
	Description: "JARM is a method for creating SSL/TLS server fingerprints",
	DataType:    HEXADECIMAL,
}

var macAddr = Definition{
	Type:        "mac-address",
	Description: "Network interface hardware address",
	DataType:    MAC,
}

var malwareFamily = Definition{
	Type:        "malware-family",
	Description: "Malware family",
	DataType:    ISTR,
}

var malwareType = Definition{
	Type:        "malware-type",
	Description: "Malware type",
	DataType:    ISTR,
}

var hashMD5 = Definition{
	Type:        "md5",
	Description: "Hash MD5",
	DataType:    MD5,
}

var mimeType = Definition{
	Type:        "mime-type",
	Description: "A media type (also MIME type and content type) is a two-part identifier",
	DataType:    MIME,
}

var mobileAppID = Definition{
	Type:        "mobile-app-id",
	Description: "The ID of a mobile application",
	DataType:    STR,
}

var passport = Definition{
	Type:        "passport",
	Description: "Passport number",
	DataType:    STR,
	Attributes:  []Definition{issuingCountry, issuer, dateOfIssue, expirationDate},
}

var pathD = Definition{
	Type:        "path",
	Description: "Path to a file, folder or process, also a HTTP request path",
	DataType:    PATH,
}

var patternInFile = Definition{
	Type:        "pattern-in-file",
	Description: "Pattern inside a file",
	DataType:    STR,
}

var patternInMemory = Definition{
	Type:        "pattern-in-memory",
	Description: "Pattern in memory",
	DataType:    STR,
}

var patternInTraffic = Definition{
	Type:        "pattern-in-traffic",
	Description: "Pattern in traffic",
	DataType:    STR,
}

var pgpPrivateKey = Definition{
	Type:        "pgp-private-key",
	Description: "PGP private key",
	DataType:    STR,
}

var pgpPublicKey = Definition{
	Type:        "pgp-public-key",
	Description: "PGP public key",
	DataType:    STR,
}

var phone = Definition{
	Type:        "phone",
	Description: "Phone number",
	DataType:    PHONE,
}

var pnr = Definition{
	Type:        "pnr",
	Description: "The Passenger Name Record Locator is a key under which the reservation for a trip is stored in the system. The PNR contains, among other data, the name, flight segments and address of the passenger. It is defined by a combination of five or six letters and numbers",
	DataType:    STR,
}

var process = Definition{
	Type:        "process",
	Description: "A running process",
	DataType:    ISTR,
	Attributes:  []Definition{processState},
}

var processState = Definition{
	Type:        "process-state",
	Description: "State of a process",
	DataType:    ISTR,
}

var pRtn = Definition{
	Type:        "prtn",
	Description: "Premium-rate telephone number",
	DataType:    ISTR,
}

var redress = Definition{
	Type:        "redress-number",
	Description: "The Redress Control Number is the record identifier for people who apply for redress through the DHS Travel Redress Inquiry Program (DHS TRIP). DHS TRIP is for travelers who have been repeatedly identified for additional screening and who want to file an inquiry to have erroneous information corrected in DHS systems",
	DataType:    STR,
}

var regKey = Definition{
	Type:        "regkey",
	Description: "Registry key",
	DataType:    ISTR,
}

var hashSHA1 = Definition{
	Type:        "sha1",
	Description: "Hash SHA1",
	DataType:    SHA1,
}

var hashSHA224 = Definition{
	Type:        "sha224",
	Description: "Hash SHA224",
	DataType:    SHA224,
}

var hashSHA256 = Definition{
	Type:        "sha256",
	Description: "Hash SHA256",
	DataType:    SHA256,
}

var hashSHA384 = Definition{
	Type:        "sha384",
	Description: "Hash SHA384",
	DataType:    SHA384,
}

var hashSHA512 = Definition{
	Type:        "sha512",
	Description: "Hash SHA512",
	DataType:    SHA512,
}

var hashSHA3224 = Definition{
	Type:        "sha3-224",
	Description: "Hash SHA3-224",
	DataType:    SHA3_224,
}

var hashSHA3256 = Definition{
	Type:        "sha3-256",
	Description: "Hash SHA3-256",
	DataType:    SHA3_256,
}

var hashSHA3384 = Definition{
	Type:        "sha3-384",
	Description: "Hash SHA3-384",
	DataType:    SHA3_384,
}

var hashSHA3512 = Definition{
	Type:        "sha3-512",
	Description: "Hash SHA3-512",
	DataType:    SHA3_512,
}

var hashSHA512224 = Definition{
	Type:        "sha512-224",
	Description: "Hash SHA512-224",
	DataType:    SHA512_224,
}

var hashSHA512256 = Definition{
	Type:        "sha512-256",
	Description: "Hash SHA512-256",
	DataType:    SHA512_256,
}

var sshFingerprint = Definition{
	Type:        "ssh-fingerprint",
	Description: "A fingerprint of SSH key material",
	DataType:    STR,
}

var ssr = Definition{
	Type:        "ssr",
	Description: "A Special Service Request is a function to an airline to provide a particular facility for A Passenger or passengers",
	DataType:    STR,
}

var category = Definition{
	Type:        "category",
	Description: "A category",
	DataType:    ISTR,
}

var threat = Definition{
	Type:        "threat",
	Description: "A cybersecurity threat",
	DataType:    ISTR,
}

var tikTokProfile = Definition{
	Type:        "tiktok-profile",
	Description: "TikTok user profile",
	DataType:    URL,
}

var twitterProfile = Definition{
	Type:        "twitter-profile",
	Description: "A Twitter user profile",
	DataType:    URL,
}

var uri = Definition{
	Type:        "url",
	Description: "URL",
	DataType:    URL,
}

var username = Definition{
	Type:        "username",
	Description: "Username",
	DataType:    ISTR,
}

var visa = Definition{
	Type:        "visa",
	Description: "Visa number",
	DataType:    STR,
}

var whoIsRegistrant = Definition{
	Type:        "whois-registrant",
	Description: "Who is registrant",
	DataType:    ISTR,
}

var whoIsRegistrar = Definition{
	Type:        "whois-registrar",
	Description: "whois-registrar",
	DataType:    ISTR,
}

var windowsScheduledTask = Definition{
	Type:        "windows-scheduled-task",
	Description: "A Windows scheduled task",
	DataType:    ISTR,
}

var windowsServiceDisplayName = Definition{
	Type:        "windows-service-displayname",
	Description: "A windows service’s displayname, not to be confused with the windows-service-name. This is the name that applications will generally display as the service’s name in applications",
	DataType:    ISTR,
}

var windowsServiceName = Definition{
	Type:        "windows-service-name",
	Description: "A windows service name. This is the name used internally by windows. Not to be confused with the windows-service-displayname",
	DataType:    ISTR,
}

var xmr = Definition{
	Type:        "xmr",
	Description: "Monero address",
	DataType:    STR,
}

var x509MD5 = Definition{
	Type:        "x509-fingerprint-md5",
	Description: "X509 fingerprint in MD5",
	DataType:    MD5,
}

var x509SHA1 = Definition{
	Type:        "x509-fingerprint-sha1",
	Description: "X509 fingerprint in SHA1",
	DataType:    SHA1,
}

var x509SHA256 = Definition{
	Type:        "x509-fingerprint-sha256",
	Description: "X509 fingerprint in SHA256",
	DataType:    SHA256,
}

var breach = Definition{
	Type: "breach",
	Description: "Security breach that resulted in a leak of PII or SPII",
	DataType: UUID,
	Attributes: []Definition{domain, link, breachDate, breachCount, breachDescription},
}

var breachDate = Definition{
	Type: "breach-date",
	Description: "Day the breach occurred",
	DataType: DATE,
}

var breachCount = Definition{
	Type: "breach-count",
	Description: "Number of items leaked in the breach",
	DataType: INTEGER,
}

var breachDescription = Definition{
	Type: "breach-description",
	Description: "Detailed description of the breach",
	DataType: STR,
}

var Definitions = []Definition{
	breach,
	breachDate,
	breachCount,
	breachDescription,
	file,
	fileData,
	lastAnalysis,
	adversary,
	aso,
	asn,
	malware,
	malwareFamily,
	malwareType,
	malwareSample,
	object,
	descriptor,
	abaRtn,
	latitude,
	longitude,
	country,
	cookie,
	text,
	value,
	issuer,
	password,
	airport,
	profilePhoto,
	authentiHash,
	bankAccountNr,
	bic,
	bin,
	btc,
	ccNumber,
	cdHash,
	certificateFingerprint,
	chromeExtension,
	subnet,
	cpe,
	cve,
	dash,
	dkim,
	dkimSignature,
	domain,
	city,
	issuingCountry,
	emailAddress,
	emailBody,
	emailDisplayName,
	emailHeader,
	emailMimeBoundary,
	emailSubject,
	emailThreadIndex,
	emailXMailer,
	email,
	eppn,
	facebookProfile,
	ffn,
	filename,
	sizeInBytes,
	filenamePattern,
	flight,
	gene,
	gitHubOrganization,
	gitHubRepository,
	gitHubUser,
	link,
	datetime,
	date,
	dateOfIssue,
	expirationDate,
	group,
	haSSHMD5,
	haSSHServerMD5,
	hexa,
	base64d,
	hostname,
	iban,
	idNumber,
	ipAddr,
	ja3Fingerprint,
	jabberID,
	jarmFingerprint,
	macAddr,
	hashMD5,
	mimeType,
	mobileAppID,
	passport,
	pathD,
	patternInFile,
	patternInMemory,
	patternInTraffic,
	pgpPrivateKey,
	pgpPublicKey,
	phone,
	pnr,
	process,
	processState,
	pRtn,
	redress,
	regKey,
	hashSHA1,
	hashSHA224,
	hashSHA256,
	hashSHA384,
	hashSHA512,
	hashSHA3224,
	hashSHA3256,
	hashSHA3384,
	hashSHA3512,
	hashSHA512224,
	hashSHA512256,
	sshFingerprint,
	ssr,
	category,
	threat,
	tikTokProfile,
	twitterProfile,
	uri,
	username,
	visa,
	whoIsRegistrant,
	whoIsRegistrar,
	windowsScheduledTask,
	windowsServiceDisplayName,
	windowsServiceName,
	xmr,
	x509MD5,
	x509SHA1,
	x509SHA256,	
	payload,
}

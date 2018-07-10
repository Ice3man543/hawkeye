package core

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Ice3man543/hawkeye/utils"
)

const (
	// TypeSimple denotes a Simple Signature
	TypeSimple = "simple"
	// TypePattern denotes a Pattern Signature
	TypePattern = "pattern"

	// PartExtension means match the extension of the file
	PartExtension = "extension"
	// PartFilename means match filename
	PartFilename = "filename"
	// PartPath means match the path
	PartPath = "path"
	// PartContent means match the file content
	PartContent = "content"
)

// SimpleSignature is a signature for simple string comparison
// matches
type SimpleSignature struct {
	part        string
	match       string
	description string
	comment     string
}

// PatternSignature is a signature for regex string comparison
// type matches
type PatternSignature struct {
	part        string
	match       *regexp.Regexp
	description string
	comment     string
}

// MatchFile contains the contents to match the signatures of a file
type MatchFile struct {
	Path      string
	Filename  string
	Extension string
}

type Signature interface {
	Match(file MatchFile) bool
	Description() string
	Comment() string
}

var DefaultSignatures = []string{"CryptoFiles", "PasswordFiles", "ConfigurationFiles", "DatabaseFiles", "MiscFiles"}
var skippableExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".tif", ".psd", ".xcf"}

func (f *MatchFile) IsSkippable() bool {
	ext := strings.ToLower(f.Extension)
	for _, skippableExt := range skippableExtensions {
		if ext == skippableExt {
			return true
		}
	}
	return false
}

func (s SimpleSignature) Match(file MatchFile) bool {
	var haystack *string
	switch s.part {
	case PartPath:
		haystack = &file.Path
	case PartFilename:
		haystack = &file.Filename
	case PartExtension:
		haystack = &file.Extension
	default:
		return false
	}

	return s.match == *haystack
}

func (s SimpleSignature) Description() string {
	return s.description
}

func (s SimpleSignature) Comment() string {
	return s.comment
}

func (s PatternSignature) Match(file MatchFile) bool {
	var haystack *string
	switch s.part {
	case PartPath:
		haystack = &file.Path
	case PartFilename:
		haystack = &file.Filename
	case PartExtension:
		haystack = &file.Extension
	default:
		return false
	}

	return s.match.MatchString(*haystack)
}

func (s PatternSignature) Description() string {
	return s.description
}

func (s PatternSignature) Comment() string {
	return s.comment
}

func NewMatchFile(path string) MatchFile {
	_, filename := filepath.Split(path)
	extension := filepath.Ext(path)
	return MatchFile{
		Path:      path,
		Filename:  filename,
		Extension: extension,
	}
}

// ParseSignaturesFromCommandLine parses signatures from command line
func ParseSignaturesFromCommandLine(state *utils.State) []string {
	signatures := []string{}

	if state.Signatures != "" {
		signatures = strings.Split(state.Signatures, ",")
	} else {
		signatures = DefaultSignatures
	}

	for _, sig := range signatures {
		if strings.EqualFold(sig, "cryptofiles") {
			state.Signature.CryptoFiles = true
		} else if strings.EqualFold(sig, "passwordfiles") {
			state.Signature.PasswordFiles = true
		} else if strings.EqualFold(sig, "configurationfiles") {
			state.Signature.ConfigurationFiles = true
		} else if strings.EqualFold(sig, "databasefiles") {
			state.Signature.DatabaseFiles = true
		} else if strings.EqualFold(sig, "miscfiles") {
			state.Signature.MiscFiles = true
		}
	}

	if state.ExcludeSignatures != "" {
		excludedSignatures := strings.Split(state.ExcludeSignatures, ",")

		for _, sig := range excludedSignatures {
			if strings.EqualFold(sig, "cryptofiles") {
				state.Signature.CryptoFiles = false
			} else if strings.EqualFold(sig, "passwordfiles") {
				state.Signature.PasswordFiles = false
			} else if strings.EqualFold(sig, "configurationfiles") {
				state.Signature.ConfigurationFiles = false
			} else if strings.EqualFold(sig, "databasefiles") {
				state.Signature.DatabaseFiles = false
			} else if strings.EqualFold(sig, "miscfiles") {
				state.Signature.MiscFiles = false
			}
		}
	}

	return signatures
}

// CryptoFilesSignatures contains signatures for cryptographic
// entities like SSH Keys, Private Keys, etc
var CryptoFilesSignatures = []Signature{
	SimpleSignature{
		part:        PartExtension,
		match:       ".pem",
		description: "Potential cryptographic private key",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".pkcs12",
		description: "Potential cryptographic key bundle",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".p12",
		description: "Potential cryptographic key bundle",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".pfx",
		description: "Potential cryptographic key bundle",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".asc",
		description: "Potential cryptographic key bundle",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".jks",
		description: "Java keystore file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".keychain",
		description: "Apple Keychain database file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^.*_rsa.pub$`),
		description: "Public SSH key",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^.*_dsa.pub$`),
		description: "Public SSH key",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^.*_ed25519.pub$`),
		description: "Public SSH key",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^.*_ecdsa.pub$`),
		description: "Public SSH key",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^.*_rsa$`),
		description: "Private SSH key",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^.*_dsa$`),
		description: "Private SSH key",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^.*_ed25519$`),
		description: "Private SSH key",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^.*_ecdsa$`),
		description: "Private SSH key",
		comment:     "",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`\.?ssh/config$`),
		description: "SSH configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartExtension,
		match:       regexp.MustCompile(`^key(pair)?$`),
		description: "Potential cryptographic private key",
		comment:     "",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "otr.private_key",
		description: "Pidgin OTR private key",
		comment:     "",
	},
}

// PasswordFileSignatures can contain passwords, api keys, encrypted
// storge volumes for credentials, etc
var PasswordFileSignatures = []Signature{
	SimpleSignature{
		part:        PartExtension,
		match:       ".bek",
		description: "Microsoft BitLocker recovery key file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".tpm",
		description: "Microsoft BitLocker Trusted Platform Module password file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".psafe3",
		description: "Password Safe database file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "secret_token.rb",
		description: "Ruby On Rails secret token configuration file",
		comment:     "If the Rails secret token is known, it can allow for remote code execution (http://www.exploit-db.com/exploits/27527/)",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "carrierwave.rb",
		description: "Carrierwave configuration file",
		comment:     "Can contain credentials for cloud storage systems such as Amazon S3 and Google Storage",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "database.yml",
		description: "Potential Ruby On Rails database configuration file",
		comment:     "Can contain database credentials",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "omniauth.rb",
		description: "OmniAuth configuration file",
		comment:     "The OmniAuth configuration file can contain client application secrets",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "settings.py",
		description: "Django configuration file",
		comment:     "Can contain database credentials, cloud storage system credentials, and other secrets",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".agilekeychain",
		description: "1Password password manager database file",
		comment:     "Feed it to Hashcat and see if you're lucky",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "credentials.xml",
		description: "Potential Jenkins credentials file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".kwallet",
		description: "KDE Wallet Manager database file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "proftpdpasswd",
		description: "cPanel backup ProFTPd credentials file",
		comment:     "Contains usernames and password hashes for FTP accounts",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "robomongo.json",
		description: "Robomongo MongoDB manager configuration file",
		comment:     "Can contain credentials for MongoDB databases",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "filezilla.xml",
		description: "FileZilla FTP configuration file",
		comment:     "Can contain credentials for FTP servers",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "recentservers.xml",
		description: "FileZilla FTP recent servers file",
		comment:     "Can contain credentials for FTP servers",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "ventrilo_srv.ini",
		description: "Ventrilo server configuration file",
		comment:     "Can contain passwords",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "terraform.tfvars",
		description: "Terraform variable config file",
		comment:     "Can contain credentials for terraform providers",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       ".exports",
		description: "Shell configuration file",
		comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       ".functions",
		description: "Shell configuration file",
		comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       ".extra",
		description: "Shell configuration file",
		comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?pgpass$`),
		description: "PostgreSQL password file",
		comment:     "",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`\.?recon-ng/keys\.db$`),
		description: "Recon-ng web reconnaissance framework API key database",
		comment:     "",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`etc/shadow$`),
		description: "Potential Linux shadow file",
		comment:     "Contains hashed passwords for system users",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`etc/passwd$`),
		description: "Potential Linux passwd file",
		comment:     "Contains system user information",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?dockercfg$`),
		description: "Docker configuration file",
		comment:     "Can contain credentials for public or private Docker registries",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?npmrc$`),
		description: "NPM configuration file",
		comment:     "Can contain credentials for NPM registries",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`\.?chef/(.*)\.pem$`),
		description: "Chef private key",
		comment:     "Can be used to authenticate against Chef servers",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`doctl/config.yaml$`),
		description: "DigitalOcean doctl command-line client configuration file",
		comment:     "Contains DigitalOcean API key and other information",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?git-credentials$`),
		description: "git-credential-store helper credentials file",
		comment:     "",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`config/hub$`),
		description: "GitHub Hub command-line client configuration file",
		comment:     "Can contain GitHub API access token",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^(\.|_)?netrc$`),
		description: "Configuration file for auto-login process",
		comment:     "Can contain username and password",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`\.?gem/credentials$`),
		description: "Rubygems credentials file",
		comment:     "Can contain API key for a rubygems.org account",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?(bash|zsh|csh)rc$`),
		description: "Shell configuration file",
		comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?(bash_|zsh_)?profile$`),
		description: "Shell profile configuration file",
		comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?(bash_|zsh_)?aliases$`),
		description: "Shell command alias configuration file",
		comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`\.?aws/credentials$`),
		description: "AWS CLI credentials file",
		comment:     "",
	},
}

// ConfigurationFileSignatures contains signatures for configuration files
var ConfigurationFileSignatures = []Signature{
	SimpleSignature{
		part:        PartExtension,
		match:       ".ovpn",
		description: "OpenVPN client configuration file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".cscfg",
		description: "Azure service configuration schema file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".rdp",
		description: "Remote Desktop connection file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "jenkins.plugins.publish_over_ssh.BapSshPublisherPlugin.xml",
		description: "Jenkins publish over SSH plugin file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "LocalSettings.php",
		description: "Potential MediaWiki configuration file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".tblk",
		description: "Tunnelblick VPN configuration file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "configuration.user.xpl",
		description: "Little Snitch firewall configuration file",
		comment:     "Contains traffic rules for applications",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "knife.rb",
		description: "Chef Knife configuration file",
		comment:     "Can contain references to Chef servers",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`\.?purple/accounts\.xml$`),
		description: "Pidgin chat client account configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`\.?xchat2?/servlist_?\.conf$`),
		description: "Hexchat/XChat IRC client server list configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`\.?irssi/config$`),
		description: "Irssi IRC client configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?dbeaver-data-sources.xml$`),
		description: "DBeaver SQL database manager configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?muttrc$`),
		description: "Mutt e-mail client configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?s3cfg$`),
		description: "S3cmd configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^sftp-config(\.json)?$`),
		description: "SFTP connection configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?trc$`),
		description: "T command-line Twitter client configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`config(\.inc)?\.php$`),
		description: "PHP configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?htpasswd$`),
		description: "Apache htpasswd file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?tugboat$`),
		description: "Tugboat DigitalOcean management tool configuration",
		comment:     "",
	},

	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?gitconfig$`),
		description: "Git configuration file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?env$`),
		description: "Environment configuration file",
		comment:     "",
	},
}

// DatabaseFileSignatures contains signatures for database files
var DatabaseFileSignatures = []Signature{
	SimpleSignature{
		part:        PartExtension,
		match:       ".mdf",
		description: "Microsoft SQL database file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".sdf",
		description: "Microsoft SQL server compact database file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".sqlite",
		description: "SQLite database file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".fve",
		description: "Windows BitLocker full volume encrypted data file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".pcap",
		description: "Network traffic capture file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".gnucash",
		description: "GnuCash database file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartExtension,
		match:       ".dayone",
		description: "Day One journal file",
		comment:     "Now it's getting creepy...",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "journal.txt",
		description: "Potential jrnl journal file",
		comment:     "Now it's getting creepy...",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "Favorites.plist",
		description: "Sequel Pro MySQL database manager bookmark file",
		comment:     "",
	},
	PatternSignature{
		part:        PartExtension,
		match:       regexp.MustCompile(`^key(store|ring)$`),
		description: "GNOME Keyring database file",
		comment:     "",
	},
	PatternSignature{
		part:        PartExtension,
		match:       regexp.MustCompile(`^kdbx?$`),
		description: "KeePass password manager database file",
		comment:     "Feed it to Hashcat and see if you're lucky",
	},
	PatternSignature{
		part:        PartExtension,
		match:       regexp.MustCompile(`^sql(dump)?$`),
		description: "SQL dump file",
		comment:     "",
	},
}

// MiscSignatures contains signatures for miscellaneous files which don't come
// easily into any specific category
var MiscSignatures = []Signature{
	SimpleSignature{
		part:        PartExtension,
		match:       ".log",
		description: "Log file",
		comment:     "Log files can contain secret HTTP endpoints, session IDs, API keys and other goodies",
	},

	SimpleSignature{
		part:        PartFilename,
		match:       "user.txt",
		description: "Potential user.txt file",
		comment:     "",
	},
	SimpleSignature{
		part:        PartFilename,
		match:       "root.txt",
		description: "Potential root.txt file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?(bash_|zsh_|sh_|z)?history$`),
		description: "Shell command history file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?mysql_history$`),
		description: "MySQL client command history file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?psql_history$`),
		description: "PostgreSQL client command history file",
		comment:     "",
	},
	PatternSignature{
		part:        PartFilename,
		match:       regexp.MustCompile(`^\.?irb_history$`),
		description: "Ruby IRB console history file",
		comment:     "",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`credential`),
		description: "Contains word: credential",
		comment:     "",
	},
	PatternSignature{
		part:        PartPath,
		match:       regexp.MustCompile(`password`),
		description: "Contains word: password",
		comment:     "",
	},
}

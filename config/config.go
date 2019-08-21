package config

var (
	// Version => The version of the app; dynamically set during build, 0.0.1 is a placeholder
	Version = "v0.1.0"
	// SupportedTools => A list of the Hashicorp tools supported by this manager
	SupportedTools = []string{
		"terraform",
	}
	// HomeDir => Store home dir
	HomeDir = ""
	// InstallDir => Path to where to install the binaries
	InstallDir = ".hashiman/installs"
	// BinDir => Where to add the symlinks
	BinDir = "/usr/local/bin"
	// DownloadDir => Where to download to
	DownloadDir = "/tmp"
)

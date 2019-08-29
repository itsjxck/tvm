package config

var (
	// Version => The version of the app; dynamically set during build, 0.0.1 is a placeholder
	Version = "v0.1.2"
	// HomeDir => Store home dir
	HomeDir = ""
	// InstallDir => Path to where to install the binaries
	InstallDir = ".tvm/installs"
	// BinDir => Where to add the symlinks
	BinDir = "/usr/local/bin"
	// DownloadDir => Where to download to
	DownloadDir = "/tmp"
	// Tool =>
	Tool = "terraform"
)

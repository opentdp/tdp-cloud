package service

func Create() {

	if vInstall == "server" {
		serverInstall()
	}

	if vUninstall == "server" {
		serverUninstall()
	}

}

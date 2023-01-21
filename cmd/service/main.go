package service

func Create() {

	if vInstall == "server" {
		serverInstall()
	}
	if vUninstall == "server" {
		serverUninstall()
	}

	if vInstall == "worker" {
		workerInstall()
	}
	if vUninstall == "worker" {
		workerUninstall()
	}

}

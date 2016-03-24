push:
	@echo "-----BEGIN RSA PRIVATE KEY-----\n${git_ssh}-----END RSA PRIVATE KEY-----\n" > id_rsa.pub
	@git remote add deploy "${git_origin}"
	GIT_SSH_COMMAND="ssh -i id_rsa.pub -o StrictHostKeyChecking=no" git push -u deploy master

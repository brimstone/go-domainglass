push:
	@git remote add deploy "${git_origin}"
	git push -u deploy master

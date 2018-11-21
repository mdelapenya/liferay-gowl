package liferay

// DXP implementation for Liferay DXP official images
type DXP struct {
	Tag string
}

// GetContainerName returns the name of the container generated by this type of image
func (d DXP) GetContainerName() string {
	return "lpn-" + d.GetType()
}

// GetDeployFolder returns the deploy folder under Liferay Home
func (d DXP) GetDeployFolder() string {
	return d.GetLiferayHome() + "/deploy"
}

// GetDockerHubTagsURL returns the URL of the available tags on Docker Hub
func (d DXP) GetDockerHubTagsURL() string {
	return "liferay/dxp"
}

// GetFullyQualifiedName returns the fully qualified name of the image
func (d DXP) GetFullyQualifiedName() string {
	return d.GetRepository() + ":" + d.GetTag()
}

// GetLiferayHome returns the Liferay home for DXP
func (d DXP) GetLiferayHome() string {
	return "/opt/liferay"
}

// GetRepository returns the repository for DXP
func (d DXP) GetRepository() string {
	return DXPRepository
}

// GetTag returns the tag of the image
func (d DXP) GetTag() string {
	return d.Tag
}

// GetType returns the type of the image
func (d DXP) GetType() string {
	return "dxp"
}
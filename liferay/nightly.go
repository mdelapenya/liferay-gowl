// Copyright (c) 2000-present Liferay, Inc. All rights reserved.
//
// This library is free software; you can redistribute it and/or modify it under
// the terms of the GNU Lesser General Public License as published by the Free
// Software Foundation; either version 2.1 of the License, or (at your option)
// any later version.
//
// This library is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
// details.

package liferay

import internal "github.com/mdelapenya/lpn/internal"

// Nightly implementation for Liferay nightly images
type Nightly struct {
	Tag string
}

// GetContainerName returns the name of the container generated by this type of image
func (n Nightly) GetContainerName() string {
	return internal.LpnConfig.GetPortalContainerName("nightly")
}

// GetDeployFolder returns the deploy folder under Liferay Home
func (n Nightly) GetDeployFolder() string {
	return n.GetLiferayHome() + "/deploy"
}

// GetDockerHubTagsURL returns the URL of the available tags on Docker Hub
func (n Nightly) GetDockerHubTagsURL() string {
	return "mdelapenya/portal-snapshot"
}

// GetFullyQualifiedName returns the fully qualified name of the image
func (n Nightly) GetFullyQualifiedName() string {
	return "docker.io/" + n.GetRepository() + ":" + n.GetTag()
}

// GetLiferayHome returns the Liferay home for nightly builds
func (n Nightly) GetLiferayHome() string {
	return "/opt/liferay"
}

// GetRepository returns the repository for nightly builds
func (n Nightly) GetRepository() string {
	return internal.LpnConfig.GetPortalImageName("nightly")
}

// GetTag returns the tag of the image
func (n Nightly) GetTag() string {
	return n.Tag
}

// GetType returns the type of the image
func (n Nightly) GetType() string {
	return "nightly"
}

// GetUser returns the user running the main application
func (n Nightly) GetUser() string {
	return "liferay"
}

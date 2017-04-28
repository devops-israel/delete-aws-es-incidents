#!/bin/bash
set -xe
cd $HOME

# Remove the build cache and get Glide
rm -Rf $GOPATH/src/*
go get github.com/Masterminds/glide
go get -u -v github.com/jstemmer/go-junit-report

mkdir -p $HOME/.go_workspace/src/github.com/$CIRCLE_PROJECT_USERNAME/
cd $HOME/.go_workspace/src/github.com/$CIRCLE_PROJECT_USERNAME
mv $HOME/$CIRCLE_PROJECT_REPONAME .
cd $CIRCLE_PROJECT_REPONAME/

# Switch glide to the appropriate branch
if [[ "$CIRCLE_BRANCH" != "master" ]]; then
	sed -i "s/version: master/version: ${CIRCLE_BRANCH}/g" glide.yaml
fi

# Get the deps
glide install
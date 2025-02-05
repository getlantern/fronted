#!/usr/bin/env bash

set -e
# Check if yq is installed and install it if not
if ! command -v yq &> /dev/null
then
    echo "yq could not be found, installing it now"
    # If we're on MacOS, use brew to install yq
    if [[ "$OSTYPE" == "darwin"* ]]; then
        brew install yq
    # If we're on Linux, use apt-get to install yq
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        sudo apt-get install yq
    else
        echo "Unsupported OS"
        exit 1
    fi
fi

cd ../lantern-binaries
echo `pwd`
echo "trustedcas:" > fronted.yaml
curl https://globalconfig.flashlightproxy.com/global.yaml.gz | gunzip | yq '.trustedcas' >> fronted.yaml
curl https://globalconfig.flashlightproxy.com/global.yaml.gz | gunzip | yq '.client.fronted' >> fronted.yaml

# Compress the generated file
gzip -c fronted.yaml > fronted.yaml.gz

# If the generated file is different from the current one, commit and push the changes
if ! git diff --quiet fronted.yaml.gz; then
    git add fronted.yaml.gz
    git commit -m "Update fronted.yaml.gz"
    git push
else
    echo "No changes detected in fronted.yaml.gz"
fi
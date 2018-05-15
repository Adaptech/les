#!/usr/bin/env bash

version=$1
repo_path=$2
if [[ -z "$version" ]]; then
  echo "usage: $0 <version-name>"
  exit 1
fi
if [[ -z "$repo_path" ]]; then
  echo "usage: $0 <version-name> <repo-path>"
  exit 1
fi
if [[ -z "`which jq`" ]]; then
  echo 'This script requires jq. Run: sudo apt-get install jq'
  exit 1
fi

platforms=("linux/amd64" "darwin/amd64" "windows/amd64")

# Ask for github credentials so we don't have to login for each request
echo 'Enter your Github credentials'
read -p 'username: ' github_username
read -s -p 'token or password: ' github_token
echo ''
echo ''

# Create the release
echo 'Creating the github release...'
rc=`curl -s -u ${github_username}:${github_token} \
 -H 'Content-Type:application/json' \
 -d '{"tag_name":"release-'"${version}"'","target_commitish": "master","name":"v'"${version}"'","body":"Release '"${version}"'","draft":false,"prerelease":false}' https://api.github.com/repos/${repo_path}/releases`
error_code=`echo ${rc} | jq -r '.errors[0].code'`
tmp=`echo ${rc} | jq -r '.upload_url'`
if [ ${error_code} = "already_exists" ]; then
    rc=`curl -s -u ${github_username}:${github_token} https://api.github.com/repos/${repo_path}/releases`
    tmp=`echo ${rc} | jq -r '.[] | select(.tag_name=="release-'"${version}"'") | .upload_url'`
elif [ ${tmp} = "null" ]; then
    echo 'Something went wrong'
    echo ${rc}
    exit 1
fi
upload_url=${tmp%%\{*}

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    platform_executable_name=${platform_split[0]}-${platform_split[1]}

    if [ ${platform_split[0]} = "windows" ]; then
        platform_executable_name+='.exe'
    fi

    echo "Uploading les-${platform_executable_name}..."
    curl -s -u ${github_username}:${github_token} \
     -H 'Content-Type:application/octet-stream' \
     -d @../../releases/les/${version}/les-${platform_executable_name} \
     ${upload_url}?name=les-${platform_executable_name}
    echo ''

    echo "Uploading les-node-${platform_executable_name}..."
    curl -s -u ${github_username}:${github_token} \
     -H 'Content-Type:application/octet-stream' \
     -d @../../releases/les-node/${version}/les-node-${platform_executable_name} \
     ${upload_url}?name=les-node-${platform_executable_name}
    echo ''

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done

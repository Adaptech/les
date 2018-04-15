#!/usr/bin/env bash

version=$1
output_file=$2
if [[ -z "$version" ]]; then
  echo "usage: $0 <version-name>"
  exit 1
fi
if [[ -z "$output_file" ]]; then
  echo "usage: $0 <version-name> <output-file>"
  exit 1
fi

platforms=("windows/amd64" "linux/amd64" "darwin/amd64")

echo -e "# Installation\n" > ${output_file}
for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    platform_executable_name=${platform_split[0]}-${platform_split[1]}
    instructions="## ${platform_split[0]^} ${platform_split[1]^^}\n
    Install the 'les' validation tool:

    \`\`\`sudo curl -L https://github.com/Adaptech/les/blob/master/releases/les/${version}/les-${platform_executable_name}?raw=true -o /usr/local/bin/les && sudo chmod +x /usr/local/bin/les\`\`\`

    Install 'les-node':

    \`\`\`sudo curl -L https://github.com/Adaptech/les/blob/master/releases/les-node/${version}/les-node-${platform_executable_name}?raw=true -o /usr/local/bin/les-node && sudo chmod +x /usr/local/bin/les-node\`\`\`
"

    if [ ${platform_split[0]} = "windows" ]; then
        platform_executable_name+='.exe'
        instructions="## ${platform_split[0]^} ${platform_split[1]^^}

    * [Download](https://github.com/Adaptech/les/blob/master/releases/les/${version}/les-${platform_executable_name}?raw=true) the 'les' validation tool .exe

    * [Download](https://github.com/Adaptech/les/blob/master/releases/les-node/${version}/les-node-${platform_executable_name}?raw=true) 'les-node' .exe
"
    fi  

    echo -e "$instructions" >> ${output_file}

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done

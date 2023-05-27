#! /usr/bin/bash

set -eu
declare -a commands=("curl" "jq" "tar")

trap ctrl_c INT

function ctrl_c(){
	clean
	echo -e "\nBye!"
	exit 0
}

command_exists() {
  for command in "${commands[@]}"
  do
    if ! command -v $command &> /dev/null
    then
        echo "ERROR: $command could not be found. Install it!"
        exit 1
    fi
  done
}

clean(){
  echo -e "Cleaning $tmpdir"
  rm -r $tmpdir
}

# Pre flight checks
command_exists

happyexit(){
  echo ""
  echo "tftools successfully installed! 🎉"
  echo ""
  echo "Now run:"
  echo ""
  echo "  tftools usage"
  echo ""
  exit 0
}

validate_checksum(){
  echo "Not implemented yet"
}

# Check OS
OS=$(uname -s)
arch=$(uname -m)
cli_arch=""
case $OS in
  Darwin)
    case $arch in
      x86_64)
        cli_arch=""
        ;;
      arm64)
        cli_arch=$arch
        ;;
      *)
        echo "There is no tftools $OS support for $arch"
        exit 1
        ;;
    esac
    ;;
  Linux)
    case $arch in
      x86_64)
        cli_arch=amd64
        ;;
      armv8*)
        cli_arch=arm64
        ;;
      aarch64*)
        cli_arch=arm64
        ;;
      amd64|arm64)
        cli_arch=$arch
        ;;
      *)
        echo "There is no tftools $OS support for $arch"
        exit 1
        ;;
    esac
    ;;
  *)
    echo "There is no tftools $OS support for $arch"
    exit 1
    ;;
esac
OS=$(echo $OS | tr '[:upper:]' '[:lower:]')

download_release() {
  TFTOOLS_LATEST_VERSION=$(curl -s https://api.github.com/repos/containerscrew/tftools/releases/latest | jq -r ".name")
  INSTALLATION_PATH="/usr/local/bin/"
  tmpdir=$(mktemp -d)

  cd $tmpdir
  echo -e "Downloading... ${TFTOOLS_LATEST_VERSION}/tftools-${OS}-${cli_arch}.tar.gz \n"
  curl -L --fail --remote-name-all https://github.com/containerscrew/tftools/releases/download/${TFTOOLS_LATEST_VERSION}/tftools-${OS}-${cli_arch}.tar.gz
  tar -xzf tftools-${OS}-${cli_arch}.tar.gz tftools
}

# Start install
download_release

if [ "$EUID" -ne 0 ]
  then command_exists sudo
    sudo mv tftools $INSTALLATION_PATH
  else
    mv tftools $INSTALLATION_PATH
  chmod +x $INSTALLATION_PATH/tftools
fi

clean
happyexit

## Install
* use go
    ```shell
    go get github.com/evgeny-klyopov/compound-interest-app
    ```
* source
    * Mac OS (m1)
      ```shell
      curl -s https://api.github.com/repos/evgeny-klyopov/compound-interest-app/releases/latest \
      | grep browser_download_url \
      | grep comi.macos-arm64.tar.gz \
      | cut -d '"' -f 4 \
      | wget -qi - 
      tar -xvf comi.macos-arm64.tar.gz && mv comi /usr/bin/comi 
      ```
    * Mac OS
      ```shell
      curl -s https://api.github.com/repos/evgeny-klyopov/compound-interest-app/releases/latest \
      | grep browser_download_url \
      | grep comi.macos-amd64.tar.gz \
      | cut -d '"' -f 4 \
      | wget -qi - 
      tar -xvf comi.macos-amd64.tar.gz && mv comi /usr/bin/comi 
      ```
    * Linux
      ```shell
      curl -s https://api.github.com/repos/evgeny-klyopov/compound-interest-app/releases/latest \
      | grep browser_download_url \
      | grep comi.linux-amd64.tar.gz \
      | cut -d '"' -f 4 \
      | wget -qi - 
      tar -xvf comi.linux-amd64.tar.gz && mv comi /usr/bin/comi 
      ```
    * Windows
      ```shell
      curl -s https://api.github.com/repos/evgeny-klyopov/compound-interest-app/releases/latest \
      | grep browser_download_url \
      | grep comi.windows-amd64.tar.gz \
      | cut -d '"' -f 4 \
      | wget -qi -
      tar -xvf comi.windows-amd64.tar.gz && mv comi.exe /usr/bin/comi
      ```
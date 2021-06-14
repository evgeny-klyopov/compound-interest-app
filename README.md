# compound-interest-app

## Calculate compound interest 

### Install
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

### Options
* ```--monthly-payment=30000``` - monthly payment
* ```--date-start=2020-12-01``` - start date of investment (default: now date)
* ```--years=10``` - investment term in years (default: "1")
* ```--rate=15``` - percent rate (default: "10")
* ```--initial-payment=100000``` - initial payment (default: "0")
* ```--percent-increase-monthly-payment=15``` - annual percentage increase in monthly payment (default: "15")
* ```--avg-percent-dividend=4``` - avg percent dividend (default: "4")
* ```--percent-increase-dividend=5``` - annual percentage increase dividend (default: "5")

### Uses
```shell

```



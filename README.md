# hashshare-adapter

hashshare-adapter继承了bitcoin-adapter，主要修改了如下内容：

- 重写了Symbol = "HSS"。
- 重写了addressDecoder，实现了HSS地址编码。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建HSS.ini文件，编辑如下内容：

```ini

# RPC Server Type，0: CoreWallet RPC; 1: Explorer API
rpcServerType = 0
# node api url, if RPC Server Type = 0, use bitcoin core full node
;serverAPI = "http://127.0.0.1:8333/"
# RPC Authentication Username
rpcUser = "user"
# RPC Authentication Password
rpcPassword = "password"
# Is network test?
isTestNet = false
# support segWit
supportSegWit = false
# minimum transaction fees
minFees = "0.001"
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```

## 资料介绍

### 区块浏览器

http://128.199.202.220

### github

https://github.com/hss5747/HashShare-Core

### 适配资料

#### 地址编码的相关代码或代码链接

https://github.com/hss5747/HashShare-Core/blob/master/src/chainparams.cpp.cpp的CMainParams()

```C++
base58Prefixes[PUBKEY_ADDRESS] = std::vector<unsigned char>(1, 40);
base58Prefixes[SCRIPT_ADDRESS] = std::vector<unsigned char>(1, 13);
base58Prefixes[SECRET_KEY] = std::vector<unsigned char>(1, 212);

```
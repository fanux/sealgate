# For sealos on clouds
## Create VMs
```
sealgate vm create --provider ali --accessKey xxx --accessSecret xxx \
    --region cn-shanghai --flavor 4C4G --namePrefix kube-master \
    --num 3 --disk data=50G --passwd xxx --FIP true

sealgate vm list
NAME            FIP         IP                   
kube-master-0   44.2.12.122 192.168.0.2
kube-master-1   44.2.12.123 192.168.0.3
kube-master-2   44.2.12.124 192.168.0.4
```

参数名 | 参数值示例 | 作用
---|---|---
provider | ali | 云供应商，如阿里云腾讯云等
para | HostName=master0 | 额外参数，防止用户使用诸如限制带宽等一些高级参数

## check support clouds provider
```
sealgate cloud list

ali
tencent
```

## query support regions
```
sealgate region list -o json

{
"ali":["cn-shanghai","cn-shanghai"],
"tencent":["cn-shanghai","cn-shanghai"]
}
```

## query vm flavor
```
sealgate flavor list

2C2G
2C4G
4C4G
```

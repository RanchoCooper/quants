# Quants

Quants is a quant-trade system based on hexagonal architecture.

## 介绍
这是一款**单方向(多)现货网格交易策略**的量化项目。
支持防踏空，行情上涨，网格价也自动提升。

## 优势：🎉
1. 简单易上手
2. 安全(不用将api_secret告诉他人)

## 为什么选择币安交易所

> 火币手续费 Maker 0.2% Taker 0.2%

> 币安手续费 Maker 0.1% Taker 0.1% （加上BNB家持手续费低至0.075%）

## 如何配置

1. 修改config目录下的config.yaml文件

```
binance:
  key: ""
  secret: ""
dingding:
  access_token: ""
```

申请api_key地址: 

[币安API管理页面](https://www.binance.com/cn/usercenter/settings/api-management)

如果你还没有币安账号：

[注册页面](https://www.binancezh.top/zh-CN/register?ref=OW7U53AB)

[免翻墙地址](https://www.binancezh.cc/zh-CN/register?ref=OW7U53AB)

交易返佣20% 注册立刻返现10元，充值交易再返现10元。

或者可以注册火币账号：

[注册页面](https://www.huobi.ms/zh-cn/topic/double-reward/?invite_code=w2732223)

交易返佣15% 注册立刻返现5元，充值并且交易再返现10元

交易返佣计算公式：交易金额1W元 * 手续费比率0.1% * 0.02 = 2元（交易1w节约2元）

## TODO List

-[x] 支持模拟买入卖出，统计盈利情况

## Reference
[Binance API Doc](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)

[钉钉机器人接入 Doc](https://developers.dingtalk.com/document/robots/custom-robot-access)
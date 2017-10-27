### unLockMusic
Unlock "only available in mainland China". 解锁“仅限中国大陆”限制

### PAC 地址

> **http://pac.unlockmusic.us**

或者

> **http://pac.onenorthtech.com**


之前听网易音乐经常遇到“仅限中国大陆”的烦恼，即便买了会员也不能解决这个问题，每次都要开Firefox设置代理然后通过网页播放器听歌，然而觉得麻烦无比，而且好好的客户端不能用甚是烦恼，所以干脆做个自动代理吧，一劳永逸，用了一段时间觉得很不错，分享出来吧。

> 自动代理PAC可以不用安装任何第三方软件，仅配置一个URL就可以实现“翻墙”

因为我已经架设了一组服务器，所以下面说的配置是基于我的服务器写的，服务器的自动代理文件就是上面的pac.js里的内容，可以访问http://pac.unlockmusic.us查看。当然你可以自由的架设你自己的代理服务器，并不很难。

> 简单的讲就是将 **http://pac.unlockmusic.us** 配置到系统的自动代理中。详细的看下面


#### Mac下设置自动代理
**系统设置** -> **网络** -> 从左边选择你连接的网络 (例如：Wi-Fi) -> **高级设置** -> **代理** -> 勾选**自动代理设置** -> URL中输入 **http://pac.unlockmusic.us** -> 点击**确定** -> 点击 **应用** -> 完成

![mac](https://cloud.githubusercontent.com/assets/7279067/22134364/756f30a4-df01-11e6-838b-d9ce05428aff.gif)

#### iPhone下设置自动代理
**设置** -> **Wi-Fi** -> 点击已连接到的WiFi最右边的**蓝色图标** -> 点选最下面 **网页代理** 中的**自动**标签 -> URL中输入 **http://pac.unlockmusic.us** -> 返回即完成

![iphone](https://cloud.githubusercontent.com/assets/7279067/22134363/743484aa-df01-11e6-9510-562a2326164c.gif)


#### Window下设置自动代理 (我只有window10，所以只能说下window10怎么设置)
**全部设置** -> **网络和互联网** -> **代理** -> 打开 **自动检测设置** 和 **使用设置脚本** -> 脚本地址输入 **http://pac.unlockmusic.us** -> 完成

#### 目前支持的音乐服务商

| # |  Supplier  |
|:-:|:----------:|
| 1 | 网易云音乐   |

很遗憾没太多时间去添加其他服务器，以后会慢慢添加。

![alt tag](https://cloud.githubusercontent.com/assets/7279067/22133926/68bb017e-defe-11e6-9cbc-1735511726d4.jpg)

https://travis-ci.org/neesonqk/unLockMusic.svg?branch=master

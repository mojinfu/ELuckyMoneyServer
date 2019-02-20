

# 功能
    帮助广大贫困群众，薅饿了么的羊毛。（饿了么拼手气红包领取）

# 使用方法
    服务开启后，会部署一个http服务在配置文件描述的相应端口。
    使用者发送http请求到该服务，服务自动开始领取饿了么红包

## 服务开启所需参数
    配置文件在/bin/config.json 
    HttpHost字段配置服务部署的IP和端口
    /bin/users.json  中配置抢饿了么红包的人手，其中cookie字段为访问饿了么红包所需cookie，可通过注册QQ号，用QQ号访问红包H5抓包获得，人数自行扩展。


## 触发绑定手机号验证
    请求Method：POST

    请求URL：  http://[HttpHost]:[port]/Vip/BindPhoneSendSMS

    请求Header："Content-Type:application/json; charset=UTF-8"

    请求Body：
    {
	    "Phone":"15967180000"
    }
## 绑定手机号
    请求Method：POST

    请求URL：  http://[HttpHost]:[port]/Vip/BindPhoneCheckSMS

    请求Header："Content-Type:application/json; charset=UTF-8"

    请求Body：
    {
        "Phone": "15967180000",
        "ValidateToken": "49c23de71825bea798935a7c7f31fe0988493b4d0fc083ef27d30089d61b0435",
        "ValidateCode": "983926",
        "PhoneOwner": "mojinfu"
    }
    --  PhoneOwner  选填
    --  Phone 绑定的手机号 和 上一步请求中的一致
    --  ValidateCode 上一步中的返回结果
    --  ValidateToken 上一步中的返回结果
## 打开一个红包 

    请求Method：POST

    请求URL：  http://[HttpHost]:[port]/OpenIt

    请求Header："Content-Type:application/json; charset=UTF-8"

    请求Body：
    {
        "Phone": "15967180000",
        "EUrl": "https://h5.ele.me/hongbao/?from=singlemessage&isappinstalled=0#hardware_id=&is_lucky_group=True&lucky_number=0&track_id=&platform=0&sn=2a0e6e7eb3acf408&theme_id=3193&device_id=&refer_user_id=4339802"
    }


    -- EUrl：一个饿了么红包的链接。 （在浏览器中打开时的URL）



### 结果实例
    接口返回：lucky success!!
    红包状态：
<div align=center>
<img src="case.jpeg" width="40%" height="40%" />
</div>



                                                                                ------Auther  mojinfu
                                                                                E: mynameless@foxmail.com
## 技术选型  
  
* [MySQL](https://www.mysql.com/): 主要用于储存一些 Metadata 
* [ClickHouse](https://clickhouse.yandex): 用于储存海量的事件数据
* [Expr](https://github.com/antonmedv/expr): 用于规则引擎  

## 名词解释    

### Model：模型
  
用户行为事件，例如注册、登录、交易、充值、提现等    
  
#### 模型五要素  
  
即 Event 五要素，简单来说，一个 Event 就是描述了：谁在某个时间点、某个地方，以某种方式完成了某个具体的事情。  
    
### Abstraction：特征    
特征，例如用户一小时提现金额、IP 30 分钟登录次数、设备一小时注册次数等    
    
### Activation  
可以参考机器学习
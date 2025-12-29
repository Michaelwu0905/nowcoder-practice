# 初始化
arr=[1,2,3]

# 增
arr.append(4)

# 删
last=arr.pop() # 直接弹出最后一个元素
first=arr.pop(0) # 弹出第一个元素，慢，会导致所有元素前移

# 查
x=arr[2]

# 切片，很常用
sub=arr[1:3]
rev=arr[::-1]

# 高阶技巧，利用列表推导式快速生成数组
nums=[x*x for x in range(5)] # [0,1,4,9,16]

# 输入一行数字，转为列表
nums=list(map(int,input().split()))
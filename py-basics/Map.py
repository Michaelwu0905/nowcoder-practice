data=["1","2","3","4","5"]

# 使用map批量加工，把int()函数作用在data的每一个元素上
# py3中map很懒，它返回的是一个迭代器，必须用list包裹起来
result=list(map(int,data))


print(result)

# map妙用
def square(x):
    return x*x

nums=[1,2,3,4,5]
squared_nums=list(map(square,nums)) # 把square的函数作用在nums的每一个元素上

print(squared_nums)

# 一行代码获取多个输入的list，不管输入多少数字，统统转成list
arr=list(map(int,input().split()))

# Conclusion：map()就是一个加工流水线，常用操作就是被list包裹起来，变成一个列表
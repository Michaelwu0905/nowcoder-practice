# python的标准输入输出工作流

# 一行内输入多个变量，如为n，m赋值10 20
n,m=map(int,input().split()) # map(int,list)将列表中每个元素转为int
print(n,m)

# 输入一行数组
nums=list(map(int,input().split()))
print(nums)
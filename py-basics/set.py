s={1,2,3}

# 增
s.add(4)
s.add(2) # 无效，2已经有了

# 查询
if 2 in s:
    print("2 found")

# 数学运算
a={1,2,3}
b={3,4,5}
print(a&b) # 交集
print(a|b) # 并集
print(a-b) # 差集
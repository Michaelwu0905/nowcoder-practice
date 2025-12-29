s=input()

seen=set()  # python去重集合
res=[]  # 结果列表

# 倒序遍历
for char in s[::-1]:
    # 判断去重
    if char not in seen:
        seen.add(char)  # 在seen set中记录
        res.append(char)    # 在结果列表中存入字符

# 拼接输出
print("".join(res)) # "分割符".join(要拼接的列表)
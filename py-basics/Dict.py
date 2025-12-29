# Dict字典，即map，k-v对，key必须唯一不可变
# 常用于需要统计次数，建立映射关系，记忆化搜索的场景
d={"apple":5,"banana":3}

print(d["apple"])

# 技巧使用.get()避免不存在时的报错
count=d.get("pear",0) # 若pear不存在，返回0

# 遍历
for k,v in d.items():
    print(k,v)
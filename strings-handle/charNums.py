# 牛客网：计算某字符出现次数
s=input().lower()
c=input().lower()
count=0
for i in range(len(s)):
    if s[i]==c:
        count=count+1

print(count)

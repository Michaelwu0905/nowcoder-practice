# 牛客网： 合并表记录
n=int(input())

table={}
for _ in range(n):
    k,v=map(int,input().split())
    if k in table:
        table[k]=table[k]+v
    else:
        table[k]=v

for index in sorted(table):
    print(index,table[index])
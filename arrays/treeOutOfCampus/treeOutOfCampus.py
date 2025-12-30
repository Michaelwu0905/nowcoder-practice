# 洛谷：校门外的树
l,m=map(int,input().split())

cutted=set()

for _ in range(m):
    u,v=map(int,input().split())
    for i in range(u,v+1):
        if i not in cutted:
            cutted.add(i)

print(l+1-len(cutted))
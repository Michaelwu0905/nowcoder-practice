# 洛谷 校门外的树 数组模拟法
l,m=map(int,input().split())

# 1. 创建一个全是树的列表，true就代表有树，下标从0到L
trees=[True]*(l+1)

for _ in range(m):
    u,v=map(int,input().split())
    # 砍掉对应区间
    for i in range(u,v+1):
        trees[i]=False

# 统计剩下的True的个数
print(sum(trees))
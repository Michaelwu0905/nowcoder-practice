# 循环读取n行数据
n=int(input())

for _ in range(n):
    # 假设每一行都是两个数字
    a,b=map(int,input().split())
    print(a+b)

# 处理未知行的输入
import sys
for line in sys.stdin:
    # line包含当前行的字符串，自带换行符
    if not line.strip(): # 防止空行干扰
        break
    parts=list(map(int,line.split()))
    print(sum(parts))
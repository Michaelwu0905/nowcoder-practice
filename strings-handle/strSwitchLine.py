# 牛客网：字符串分隔
s=input()

i=0
cnt=0
while 1:
    if cnt==len(s):
        zeroNums=8-i
        print(zeroNums*"0")
        break
    if i>7:
        print("")
        i=0
    else:
        print(s[cnt],end="")
        cnt+=1
        i+=1
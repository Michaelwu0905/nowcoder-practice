# 牛客网：数组中重复的数字
class Solution:
    def duplicate(self,numbers:list[int])->int:
        seen=set()
        for i in numbers:
            if i in seen:
                return i
            else:
                seen.add(i)
        return -1
            
s=Solution()
res=s.duplicate([2,3,1,0,2,5,3])
print(res)
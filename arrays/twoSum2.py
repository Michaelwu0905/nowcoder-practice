# 版本2 字典
class Solution:
    def twoSum(self,nums:list[int],target:int)->list[int]:
        records=dict()

        for index,value in enumerate(nums):
            if target-value in records:
                return[records[target-value],index]
            records[value]=index
        return []
    
s=Solution()
res=s.twoSum([2,7,11,15],9)
print(res)
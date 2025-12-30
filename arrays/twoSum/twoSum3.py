# 使用字典set的方法
class Solution:
    def twoSum(self,nums:list[int],target:int)->list[int]:
        seen=set()
        for i,v in enumerate(nums):
            complement=target-v
            if complement in seen:
                return [nums.index(complement),i]
            seen.add(v)
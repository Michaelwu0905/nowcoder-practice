# 牛客网：回文字符串2
class Solution:
    def palindrome(self,str:str)->bool:
        left=0
        right=len(str)-1

        while(left<right):
            if str[left]!=str[right]:
                if str[left+1:right+1]==str[left+1:right+1:-1]:
                    return True
                if str[left:right]==str[left:right:-1]:
                    return True
                else:
                    return False
            else:
                left+=1
                right-=1
        return True
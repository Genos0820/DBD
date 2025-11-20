
def twoSum(nums,target):
    my_dict={}
    for i,val in enumerate(nums):
        if (target-val) in my_dict:
            return [i,my_dict[target-val]]
        my_dict[val]=1
    return []

print(twoSum([1,2,3,4],5))
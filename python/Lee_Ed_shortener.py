
# -*- coding: utf-8 -*-
"""
Created on Thurs Jun 20 17:21:32 2019

@author: dkedw
"""

######## Shortener

# Given two strings s, t, determine if string s can be changed into string t using the following rule:

# Remove any number of lowercase letters from s (or none at all), then capitalize all remaining letters of s.

# Example of the rule:
# s = AAbcD
# Possible outputs using the rule:
# Remove none, capitalize -> AABCD
# Remove c, capitalize -> AABD
# Remove b, capitalize -> AACD
# Remove b and c, capitalize -> AAD

# If it is possible to create the string t by processing s using the rule, then the function 
# should return True, otherwise return False.

def shortener(s, t):
    
    
    x = len(s)
    y = len(t)
    
    #conditional set to False for range 
    cond = ([[False for i in range(y+1)] for i in range(x+1)])
    #first set to True
    cond[0][0] = True
    
    #for loop length of s
    for i in range(x):
        #second for loop len of t+1
        
        for j in range(y+1):
            if(cond[i][j]):
                #if met set true
                if((j < y and (s[i].upper() == t[j]))):
                    cond[i+1][j+1] = True
                #if met set true
                if(s[i].isupper() == False):
                    cond[i+1][j]  = True
                    

    return (cond[x][y]) #return False unless cond is met
    

# Test Cases
test_cases = [
    ("daBccd", "ABC", True),
    ("sYOCa", "YOCN", False),
    ("aaaaaa", "AAAAAAA", False),
    ("SVAHHHMVIIDYIcOSHMDUAVJRIBxBZQSUBIVEBHfVTZVSHATUYDJGDRRUBQFHEEEUZLQGXTNKFWUYBAeFKUHSFLZEUINBZYRIXOPYYXAEZZWELUPIEIWGZHEIYIROLQLAVHhMKRDSOQTJYYLTCTSIXIDAnPIHNXENWFFZFJASRZRDAPVYPAViVBLVGRHObnwlcyprcfhdpfjkyvgyzpovsgvlqbhtwrucvszaqinbgeafuswkjrcexvyzq","SVAHHHMVIIDYIOSHMDUAVJRIBBZQSUBIVEBHVTZVSHATUYDJGDRRUBQFHEEEUZLQGXTNKFWUYBAFKUHSFLZEUINBZYRIXOPYYXAEZZWELUPIEIWGZHEIYIROLQLAVHMKRDSOQTJYYLTCTSIXIDAPIHNXENWFFZFJASRZRDAPVYPAVVBLVGRHO", True),
    ("a", "AA", False),("UZJMUCYHpfeoqrqeodznwkxfqvzktyomkrVyzgtorqefcmffauqhufkpptaupcpxguscmsbvolhorxnjrheqhxlgukjmgncwyastmtgnwhrvvfgbhybeicaudklkyrwvghpxbtpyqioouttqqrdhbinvbywkjwjkdiynvultxxxmwxztglbqitxmcgiusfewmsvxchkryzxipbmgrnqhfmlghomfbsKjglimxuobomfwutwfcmklzcphbbfohnaxgbaqbgocghaaizyhlctupndmlhwwlxxvighhjjrctcjBvxtagxbhrbrWwsyiiyebdgyfrlztoycxpjcvmzdvfeYqaxitkfkkxwybydcwsbdiovrqwkwzbgammwslwmdesygopzndedsbdixvi","UZJMUCYH", False)
]

for case in test_cases:
    s, t, output = case
    print(shortener(s, t) == output)
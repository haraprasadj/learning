import sys
if len(sys.argv) > 1:
    test_cases = open(sys.argv[1], 'r')
else:
    test_cases = open('test_cases.txt', 'r')
test_case_list = test_cases.read().strip().split('\n')
for test in test_case_list:
    if test:
        zerolst = test.split(' ')
        reslst = []
        i = 0
        while i < len(zerolst):
            item = zerolst[i]
            if item == '0':
                i += 1
                reslst.append(zerolst[i])
            elif item == '00':
                i += 1
                reslst.append(zerolst[i].replace('0', '1'))
            i += 1
        res = int(''.join(reslst), 2)
        print res
test_cases.close()